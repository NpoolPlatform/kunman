package paypal

import (
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

var config *Config

func init() {
	config, err := LoadConfig()
	if err != nil {
		panic("Invalid paypal config")
	}
}

func cacheAccessToken(ctx context.Coontext, token *Token) error {
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
	formData := url.Values{
		"grant_type": {"client_credentials"},
	}

	resp, err := resty.R().
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

func getAccessToken(ctx context.Context) (*Token, error) {
	cachedToken, err := cachedAccessToken(ctx)
	if err == nil && !cachedToken.Expiry.IsZero() && cachedToken.Expiry.After(time.Now().Add(5*time.Minute)) {
		return cachedToken, nil
	}

	newToken, err := refreshAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return newToken, nil
}

// OneShot payment
func CreatePayment(orderID uuid.UUID, redirectURL string, cancelURL string) (map[string]interface{}, error) {
	accessToken, err := getAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	payload := map[string]interface{}{
		"intent": "CAPTURE",
		"purchase_units": []map[string]interface{}{
			{
				"amount": map[string]interface{}{
					"currency_code": currency,
					"value":         amount,
				},
				"description": description,
			},
		},
		"application_context": map[string]interface{}{
			"return_url": redirectURL,
			"cancel_url": calcelURL,
		},
	}

	resp, err := resty.R().
		SetAuthToken(accessToken.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/v2/checkout/orders")
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, wlog.Errorf("%v", resp.String())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, wlog.WrapError(err)
	}

	return result, nil
}

func CaptureOrder(ctx context.Context, orderID uuid.UUID) error {
	return nil
}

func HandleWebhook(event map[string]interface{}) error {

}
