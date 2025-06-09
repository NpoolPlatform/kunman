//nolint:dupl
package client

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/NpoolPlatform/kunman/framework/redis"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool/types"
	"github.com/go-resty/resty/v2"
)

const (
	F2PoolRps           = 1
	defaultLockTimeout  = time.Second / F2PoolRps
	postTimeout         = time.Second * 3
	tlsHandshakeTimeout = time.Second * 1
	maxRetreies         = 3
	successStatusCode   = 200
)

type Client struct {
	BaseURL         string
	AccessToken     string
	TimeTokenBucket chan struct{}
}

func NewClient(baseURL, accessToken string) *Client {
	return &Client{BaseURL: baseURL, AccessToken: accessToken}
}

func (cli *Client) post(path string, req, resp interface{}) error {
	lockKey := fmt.Sprintf("f2pool_client_%v", cli.AccessToken)
	var err error
	for i := 0; i < maxRetreies; i++ {
		err = redis.TryLock(lockKey, defaultLockTimeout)
		if err == nil {
			break
		}
		time.Sleep(defaultLockTimeout)
		err = fmt.Errorf("f2pool api busy")
	}
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["F2P-API-SECRET"] = cli.AccessToken

	url := fmt.Sprintf("%v%v", cli.BaseURL, path)
	errResp := types.ErrorResponse{}

	socksProxy := os.Getenv("ENV_F2POOL_REQUEST_PROXY")

	// set tls handshake timeout
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   tlsHandshakeTimeout,
				DualStack: true,
			}).DialContext,
			TLSHandshakeTimeout: tlsHandshakeTimeout,
		},
	}
	restycli := resty.NewWithClient(client)
	if socksProxy != "" {
		restycli = restycli.SetProxy(socksProxy)
	}

	restyreq := restycli.
		SetTimeout(postTimeout).
		SetHeaders(headers).
		R().
		SetBody(reqBody).
		SetResult(resp).
		SetError(errResp)

	ret, err := restyreq.Post(url)
	if err != nil {
		return err
	}

	if ret.StatusCode() != successStatusCode {
		return fmt.Errorf("wrong response,status code: %v,response: %v", ret.StatusCode(), string(ret.Body()))
	}

	err = json.Unmarshal(ret.Body(), &errResp)
	if err == nil && errResp.Code != 0 {
		return fmt.Errorf("request api %v error,status_code: %v,err: %v", url, errResp.Code, errResp.Msg)
	}

	return nil
}
