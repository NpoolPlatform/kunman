package paypal

import (
	"fmt"
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
		Amount      Amount `json:"amount"`
	} `json:"purchase_units"`
	Links []Link `json:"links"`
	Payer struct {
		EmailAddress string `json:"email_address"`
	} `json:"payer"`
}

func (r *CreatePaymentResponse) ApproveLink() string {
	for _, link := range r.Links {
		if link.Rel == "approve" {
			return link.Href
		}
	}
	return ""
}

type ErrorDetail struct {
	Field       string `json:"field,omitempty"`
	Issue       string `json:"issue"`
	Description string `json:"description,omitempty"`
}

type ErrorResponse struct {
	Name            string        `json:"name"`
	Message         string        `json:"message"`
	DebugID         string        `json:"debug_id"`
	Details         []ErrorDetail `json:"details"`
	InformationLink string        `json:"information_link,omitempty"`
}

type PaypalPayment struct {
	Id            string `json:"id"`
	Status        string `json:"status"`
	PurchaseUnits []struct {
		Amount struct {
			Total    string `json:"total"`
			Currency string `json:"currency_code"`
		} `json:"amount"`
	} `json:"purchase_units"`
}

func (p *PaypalPayment) Approved() bool {
	return p.Status == "APPROVED"
}

func (p *PaypalPayment) String() string {
	return fmt.Sprintf("%v: %v", p.Id, p.Status)
}

type PlanMetaData struct {
	CreateTime     string  `json:"create_time,omitempty"`
	UpdateTime     string  `json:"update_time,omitempty"`
	Links          []Link  `json:"links,omitempty"`
	Category       string  `json:"category,omitempty"`
	ImageURL       string  `json:"image_url,omitempty"`
	HomeURL        string  `json:"home_url,omitempty"`
	ShippingAmount *Amount `json:"shipping_amount,omitempty"`
}

type TaxInfo struct {
	Percentage string `json:"percentage"`
	Inclusive  bool   `json:"inclusive"`
}

type PaymentPref struct {
	AutoBillOutstanding     bool    `json:"auto_bill_outstanding"`
	SetupFee                *Amount `json:"setup_fee,omitempty"`
	SetupFeeFailureAction   string  `json:"setup_fee_failure_action"`
	PaymentFailureThreshold uint32  `json:"payment_failure_threshold"`
}

type PricingScheme struct {
	FixedPrice Amount `json:"fixed_price"`
}

type CycleFrequency struct {
	IntervalUnit  string `json:"interval_unit"`
	IntervalCount uint32 `json:"interval_count"`
}

type BillingCycle struct {
	Frequency     CycleFrequency `json:"frequency"`
	TenureType    string         `json:"tenure_type"`
	Sequence      uint32         `json:"sequence"`
	TotalCycles   uint32         `json:"total_cycles"`
	PricingScheme PricingScheme  `json:"pricing_scheme"`
}

type CreatePlanRequest struct {
	ProductID          string         `json:"product_id"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	Status             string         `json:"status"`
	BillingCycles      []BillingCycle `json:"billing_cycles"`
	PaymentPreferences PaymentPref    `json:"payment_preferences"`
	Taxes              TaxInfo        `json:"taxes"`
	QuantitySupported  bool           `json:"quantity_supported,omitempty"`
	CustomID           string         `json:"custom_id,omitempty"`
	MetaData           *PlanMetaData  `json:"metadata,omitempty"`
}

type CreatePlanResponse struct {
	ID                 string         `json:"id"`
	ProductID          string         `json:"product_id"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	Status             string         `json:"status"`
	BillingCycles      []BillingCycle `json:"billing_cycles"`
	PaymentPreferences PaymentPref    `json:"payment_preferences"`
	Taxes              TaxInfo        `json:"taxes"`
	CreateTime         string         `json:"create_time"`
	UpdateTime         string         `json:"update_time,omitempty"`
	Links              []Link         `json:"links"`
	QuantitySupported  bool           `json:"quantity_supported,omitempty"`
	CustomID           string         `json:"custom_id,omitempty"`
	MetaData           *PlanMetaData  `json:"metadata,omitempty"`
}
