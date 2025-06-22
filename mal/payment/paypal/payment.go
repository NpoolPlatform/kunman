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

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Item struct {
	Name       string `json:"name"`
	Quantity   string `json:"quantity"`
	UnitAmount Amount `json:"unit_amount"`
}

type PurchaseUnit struct {
	Amount      Amount `json:"amount"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
}

type ApplicationContext struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type CreatePaymentRequest struct {
	Intent             string             `json:"intent"`
	PurchaseUnits      []PurchaseUnit     `json:"purchase_units"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type CreatePaymentResponse struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	CreateTime    string `json:"create_time"`
	PurchaseUnits []struct {
		ReferenceID string `json:"reference_id"`
		Amount      struct {
			CurrencyCode string `json:"currency_code"`
			Value        string `json:"value"`
		} `json:"amount"`
	} `json:"purchase_units"`
	Links []Link `json:"links"`
	Payer struct {
		EmailAddress string `json:"email_address"`
	} `json:"payer"`
}

type ErrorResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Details []struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	} `json:"details"`
}

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

func (cli *PaymentClient) CaptureOrder(ctx context.Context) error {
	return nil
}
