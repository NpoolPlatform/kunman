package tron

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/logger"

	tronclient "github.com/Geapefurit/gotron-sdk/pkg/client"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/endpoints"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/utils"
	"google.golang.org/grpc"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 20
	retriesSleepTime = 200 * time.Millisecond
	dialTimeout      = 3 * time.Second
)

type TClientI interface {
	GetGRPCClient(timeout time.Duration, endpointmgr *endpoints.Manager) (*tronclient.GrpcClient, string, error)
	WithClient(fn func(*tronclient.GrpcClient) (bool, error)) error
}

type tClients struct{}

func (tClients *tClients) GetGRPCClient(timeout time.Duration, endpointmgr *endpoints.Manager) (*tronclient.GrpcClient, string, error) {
	endpoint, err := endpointmgr.Peek()
	if err != nil {
		return nil, "", err
	}

	ntc := tronclient.NewGrpcClientWithTimeout(endpoint, timeout)
	err = ntc.Start(grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, "", err
	}

	return ntc, endpoint, nil
}

type SyncBlock struct {
	StartingBlock string
	CurrentBlock  string
	HighestBlock  string
}

type SyncingResponse struct {
	ID     int
	Result SyncBlock
}

var client = &http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (tClients *tClients) SyncProgress(ctx context.Context, ip, port string) (*SyncingResponse, error) {
	addr := fmt.Sprintf("http://%v:%v/jsonrpc", ip, port)
	req, err := http.NewRequest(http.MethodPost, addr, strings.NewReader(`{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":64}`))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	syncRes := SyncingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&syncRes); err != nil {
		return nil, err
	}

	if syncRes.Result.CurrentBlock >= syncRes.Result.HighestBlock {
		return nil, nil
	}

	return &syncRes, nil
}

func (tClients *tClients) WithClient(fn func(*tronclient.GrpcClient) (bool, error)) error {
	var (
		err, apiErr error
		retry       bool
		client      *tronclient.GrpcClient
		ep          string
	)

	endpointmgr, err := endpoints.NewManager()
	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, endpointmgr.Len()); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}
		client, ep, err = tClients.GetGRPCClient(dialTimeout, endpointmgr)
		if err != nil {
			logger.Sugar().Errorw("WithClient", "Endpoint", ep, "Error", err)
			continue
		}

		retry, apiErr = fn(client)
		client.Stop()
		if !retry {
			logger.Sugar().Errorw("WithClient", "Endpoint", ep, "Error", apiErr)
			return apiErr
		}
	}

	if apiErr != nil {
		return apiErr
	}
	return err
}

func Client() TClientI {
	return &tClients{}
}
