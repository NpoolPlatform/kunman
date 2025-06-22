package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"resty.dev/v3"
)

func (cli *PaymentClient) CreatePayment(ctx context.Context) (*CreatePaymentResponse, error) {
	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	amount, err := cli.orderHandler.FiatPaymentAmount()
	if err != nil {
		return nil, err
	}

	requestBody := CreatePaymentRequest{
		Intent: "CAPTURE",
		PurchaseUnits: []PurchaseUnit{
			{
				Amount: Amount{
					CurrencyCode: cli.orderHandler.FiatPaymentCurrency(),
					Value:        amount,
				},
				Description: fmt.Sprintf("Payment of %v", cli.OrderID),
			},
		},
		ApplicationContext: ApplicationContext{
			ReturnURL: cli.ReturnURL,
			CancelURL: cli.CancelURL,
		},
	}

	client := resty.New()
	defer client.Close()

	var paymentResponse CreatePaymentResponse
	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		SetHeader("Prefer", "return=representation").
		SetBody(requestBody).
		SetResult(&paymentResponse).
		Post("/v2/checkout/orders")
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() != http.StatusCreated {
		var e ErrorResponse

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, wlog.Errorf("failed read body")
		}

		if err := json.Unmarshal(bytes, &e); err == nil {
			return nil, wlog.Errorf("paypal error: %v - %v, %v", e.Name, e.Message, e.Details)
		}
		return nil, wlog.Errorf("%v: %v", resp.StatusCode(), resp.String())
	}

	return &paymentResponse, nil
}

func (cli *PaymentClient) GetPayment(ctx context.Context) (*PaypalPayment, error) {
	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	client := resty.New()
	defer client.Close()

	var paymentResponse PaypalPayment
	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		SetResult(&paymentResponse).
		Get("/v2/checkout/orders/" + cli.PaypalPaymentID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() != 200 {
		return nil, wlog.Errorf("%v: %v", resp.StatusCode(), resp.String())
	}

	return &paymentResponse, nil
}

func (cli *PaymentClient) CapturePayment(ctx context.Context) error {
	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	client := resty.New()
	defer client.Close()

	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		Post("/v2/checkout/orders/" + cli.PaypalPaymentID + "/capture")
	if err != nil {
		return wlog.WrapError(err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return wlog.Errorf("%v: %v", resp.StatusCode(), resp.String())
	}

	return nil
}
