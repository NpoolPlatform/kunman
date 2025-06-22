package paypal

import (
	"context"

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
	trialPrice := cli.appGoodHandler.TrialPrice()
	price := cli.appGoodHandler.Price()

	if trialUnits > 0 {
		billingCycles = append(billingCycles, BillingCycle{
			Frequency: CycleFrequency{
				IntervalUnit:  durationUnit,
				IntervalCount: durationUnits,
			},
			TenureType:  "TRIAL",
			Sequence:    1,
			TotalCycles: trialUnits,
			PricingScheme: PricingScheme{
				FixedPrice: Amount{
					CurrencyCode: priceUnit,
					Value:        trialPrice,
				},
			},
		})
	}
	billingCycles = append(billingCycles, BillingCycle{
		Frequency: CycleFrequency{
			IntervalUnit:  durationUnit,
			IntervalCount: durationUnits,
		},
		TenureType:  "REGULAR",
		Sequence:    1,
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
