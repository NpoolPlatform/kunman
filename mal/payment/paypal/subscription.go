package paypal

import (
	"context"
	"encoding/json"
	"io"
	"fmt"
	"net/http"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"resty.dev/v3"
)

func (cli *PaymentClient) CreatePlan(ctx context.Context) (*CreatePlanResponse, error) {
	billingCycles := []BillingCycle{}

	trialUnits := cli.appGoodHandler.TrialUnits()

	durationUnit, err := cli.appGoodHandler.DurationUnit()
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	durationUnits := cli.appGoodHandler.DurationUnits()
	priceUnit := cli.appGoodHandler.PriceUnit()

	trialPrice, err := cli.appGoodHandler.TrialPrice()
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	price, err := cli.appGoodHandler.Price()
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	fmt.Println("===================", trialPrice, price)

	sequence := uint32(1)

	if trialUnits > 0 {
		billingCycles = append(billingCycles, BillingCycle{
			Frequency: CycleFrequency{
				IntervalUnit:  durationUnit,
				IntervalCount: durationUnits,
			},
			TenureType:  "TRIAL",
			Sequence:    sequence,
			TotalCycles: trialUnits,
			PricingScheme: PricingScheme{
				FixedPrice: Amount{
					CurrencyCode: priceUnit,
					Value:        trialPrice,
				},
			},
		})
		sequence += 1
	}
	billingCycles = append(billingCycles, BillingCycle{
		Frequency: CycleFrequency{
			IntervalUnit:  durationUnit,
			IntervalCount: durationUnits,
		},
		TenureType:  "REGULAR",
		Sequence:    sequence,
		TotalCycles: 0,
		PricingScheme: PricingScheme{
			FixedPrice: Amount{
				CurrencyCode: priceUnit,
				Value:        price,
			},
		},
	})

	payload := CreatePlanRequest{
		ProductID:     cli.appGoodHandler.ProductID(),
		Name:          cli.appGoodHandler.Name(),
		Description:   cli.appGoodHandler.Description(),
		Status:        "Active",
		BillingCycles: billingCycles,
		PaymentPreferences: PaymentPref{
			AutoBillOutstanding: true,
			SetupFee: &Amount{
				CurrencyCode: priceUnit,
				Value:        "0.00",
			},
			SetupFeeFailureAction:   "CONTINUE",
			PaymentFailureThreshold: 3,
		},
		Taxes: TaxInfo{
			Percentage: "0",
			Inclusive:  false,
		},
	}

	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	client := resty.New()
	defer client.Close()

	var planResponse CreatePlanResponse
	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&planResponse).
		Post("/v1/billing/plans")
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, wlog.Errorf("%v: %v", resp.StatusCode(), resp.String())
	}

	return &planResponse, nil
}

func (cli *PaymentClient) CreateSubscription(ctx context.Context) (*CreateSubscriptionResponse, error) {
	payload := CreateSubscriptionRequest{
		PlanID: cli.PaypalPlanID,
		Subscriber: Subscriber{
			Name: Name{
				GivenName: cli.orderHandler.GivenName(),
				Surname:   cli.orderHandler.Surname(),
			},
			EmailAddress: cli.orderHandler.EmailAddress(),
			// Phone: &Phone{
			// 	PhoneType: "MOBILE",
			// 	PhoneNumber: PhoneNumber{
			// 		CountryCode:    cli.orderHandler.CountryCode(),
			// 		NationalNumber: cli.orderHandler.NationalNumber(),
			// 	},
			// },
		},
		ApplicationContext: ApplicationContext{
			ReturnURL: cli.ReturnURL,
			CancelURL: cli.CancelURL,
		},
		CustomID: cli.orderHandler.CustomID(),
	}

	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	client := resty.New()
	defer client.Close()

	var subscriptionResponse CreateSubscriptionResponse
	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&subscriptionResponse).
		Post("/v1/billing/subscriptions")
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

	return &subscriptionResponse, nil
}

func (cli *PaymentClient) GetSubscription(ctx context.Context) (*PaypalSubscription, error) {
	accessToken, err := cli.GetAccessToken(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	client := resty.New()
	defer client.Close()

	var subscriptionResponse PaypalSubscription
	resp, err := client.
		SetBaseURL(cli.config.BaseURL()).
		R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		SetResult(&subscriptionResponse).
		Get("/v1/billing/subscriptions/" + cli.PaypalSubscriptionID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if resp.StatusCode() != http.StatusOK {
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

	return &subscriptionResponse, nil
}
