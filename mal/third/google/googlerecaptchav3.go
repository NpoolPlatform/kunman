package google

import (
	"fmt"
	"os"
	"time"

	"github.com/NpoolPlatform/kunman/framework/config"
	"github.com/go-resty/resty/v2"
)

const (
	GoogleRecaptchaV3URL    = "recaptcha_url"
	GoogleRecaptchaV3Secret = "recaptcha_secret"
)

type verifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func VerifyGoogleRecaptchaV3(recaptchaToken string) error {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	recaptchaURL := config.GetStringValueWithNameSpace(hostname, GoogleRecaptchaV3URL)
	recaptchaSecret := config.GetStringValueWithNameSpace(hostname, GoogleRecaptchaV3Secret)

	if recaptchaURL == "" || recaptchaSecret == "" {
		return fmt.Errorf("invalid recaptcha parameter")
	}

	socksProxy := os.Getenv("ENV_RECAPTCHA_REQUEST_PROXY")
	url := fmt.Sprintf("%v?secret=%v&response=%v", recaptchaURL, recaptchaSecret, recaptchaToken)

	const timeout = 5

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.
		R().
		SetResult(&verifyResponse{}).
		Post(url)
	if err != nil {
		return fmt.Errorf("fail verify google recaptcha v3: %v", err)
	}

	vResp, ok := resp.Result().(*verifyResponse)
	if !ok {
		return fmt.Errorf("fail get response")
	}

	if !vResp.Success {
		return fmt.Errorf("fail verify google recaptcha v3: %v", vResp.ErrorCodes)
	}

	return nil
}
