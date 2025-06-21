package paypal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"resty.dev/v3"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	Expiry      time.Time
}

var config *Config

func init() {
	var err error

	config, err = LoadConfig()
	if err != nil {
		panic("Invalid paypal config")
	}
}

func cacheAccessToken(ctx context.Context, token *Token) error {
	key := config.AccessTokenKey()

	data, err := json.Marshal(token)
	if err != nil {
		return wlog.WrapError(err)
	}

	expiry := token.Expiry.Sub(time.Now()) - 5*time.Minute
	if expiry < 0 {
		expiry = 1 * time.Minute
	}

	err = redis2.Set(key, data, expiry)
	if err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func refreshAccessToken(ctx context.Context) (*Token, error) {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")

	cli := resty.New()
	defer cli.Close()

	resp, err := cli.R().
		SetBasicAuth(config.ClientID, config.ClientSecret).
		SetFormDataFromValues(formData).
		SetResult(&Token{}).
		Post("/v1/oauth2/token")
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, wlog.Errorf("%v", resp.String())
	}

	token := resp.Result().(*Token)
	token.Expiry = time.Now().Add(time.Second * time.Duration(token.ExpiresIn))

	if err := cacheAccessToken(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

func cachedAccessToken(ctx context.Context) (*Token, error) {
	key := config.AccessTokenKey()

	token, err := redis2.Get(key)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	_token, ok := token.(*Token)
	if !ok {
		return nil, wlog.Errorf("invalid token")
	}

	return _token, nil
}

func GetAccessToken(ctx context.Context) (string, error) {
	cachedToken, err := cachedAccessToken(ctx)
	if err == nil && !cachedToken.Expiry.IsZero() && cachedToken.Expiry.After(time.Now().Add(5*time.Minute)) {
		return cachedToken.AccessToken, nil
	}

	newToken, err := refreshAccessToken(ctx)
	if err != nil {
		return "", wlog.WrapError(err)
	}

	return newToken.AccessToken, nil
}
