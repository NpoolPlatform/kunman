package paypal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

func (a *Amount) String() string {
	return fmt.Sprintf("%v %v", a.Value, a.CurrencyCode)
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

type Name struct {
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type PhoneNumber struct {
	CountryCode    string `json:"country_code"`
	NationalNumber string `json:"national_number"`
}

type Phone struct {
	PhoneType   string      `json:"phone_type"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

type ShippingAddress struct {
	AddressLine1 string `json:"address_line_1"`
	AdminArea2   string `json:"admin_area_2"`
	AdminArea1   string `json:"admin_area_1"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}

type Subscriber struct {
	Name            Name             `json:"name"`
	EmailAddress    string           `json:"email_address"`
	Phone           Phone            `json:"phone,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	PayerID         *string          `json:"payer_id,omitempty"`
	Tenant          *string          `json:"tenant,omitempty"`
}

type CreateSubscriptionRequest struct {
	PlanID             string             `json:"plan_id"`
	StartTime          string             `json:"start_time,omitempty"`
	Subscriber         Subscriber         `json:"subscriber"`
	CustomID           string             `json:"custom_id,omitempty"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

type CreateSubscriptionResponse struct {
	ID         string     `json:"id"`
	Status     string     `json:"status"`
	PlanID     string     `json:"plan_id"`
	StartTime  string     `json:"start_time"`
	CreateTime string     `json:"create_time"`
	Links      []Link     `json:"links"`
	Subscriber Subscriber `json:"subscriber"`
}

func (r *CreateSubscriptionResponse) ApproveLink() string {
	for _, link := range r.Links {
		if link.Rel == "approve" {
			return link.Href
		}
	}
	return ""
}

type BillingInfo struct {
	LastPayment           LastPayment      `json:"last_payment"`
	NextBillingTime       string           `json:"next_billing_time"`
	FinalBillingTime      string           `json:"final_billing_time,omitempty"`
	OutstandingBalance    Amount           `json:"outstanding_balance,omitempty"`
	CycleExecutions       []CycleExecution `json:"cycle_executions"`
	FailedPaymentAttempts int              `json:"failed_payment_attempts"`
}

type LastPayment struct {
	Time   string `json:"time"`
	Amount Amount `json:"amount"`
}

type CycleExecution struct {
	CycleType            string `json:"cycle_type"`
	TenureType           string `json:"tenure_type"`
	Sequence             int    `json:"sequence"`
	TotalCyclesCompleted int    `json:"total_cycles_completed"`
	TotalCyclesRemaining int    `json:"total_cycles_remaining"`
	CurrentPeriodStart   string `json:"current_period_start"`
	CurrentPeriodEnd     string `json:"current_period_end"`
}

type PaypalSubscription struct {
	ID               string      `json:"id"`
	Status           string      `json:"status"`
	StatusUpdateTime string      `json:"status_update_time"`
	PlanID           string      `json:"plan_id"`
	StartTime        string      `json:"start_time"`
	Quantity         string      `json:"quantity"`
	ShippingAmount   Amount      `json:"shipping_amount,omitempty"`
	Subscriber       Subscriber  `json:"subscriber"`
	BillingInfo      BillingInfo `json:"billing_info"`
	CreateTime       string      `json:"create_time"`
	UpdateTime       string      `json:"update_time"`
	PlanOverridden   bool        `json:"plan_overridden"`
	Links            []Link      `json:"links"`
}

func (p *PaypalSubscription) Active() bool {
	return p.Status == "ACTIVE"
}

type SubscriptionActivatedResource struct {
	ID              string      `json:"id"`
	PlanID          string      `json:"plan_id"`
	Status          string      `json:"status"`
	StartTime       string      `json:"start_time"`
	NextBillingTime string      `json:"next_billing_time"`
	Subscriber      Subscriber  `json:"subscriber"`
	CustomID        string      `json:"custom_id"`
	BillingInfo     BillingInfo `json:"billing_info"`
}

type SubscriptionUpdatedResource struct {
	ID          string      `json:"id"`
	PlanID      string      `json:"plan_id"`
	Status      string      `json:"status"`
	UpdateTime  string      `json:"update_time"`
	Subscriber  Subscriber  `json:"subscriber"`
	BillingInfo BillingInfo `json:"billing_info"`
	CustomID    string      `json:"custom_id"`
	AutoRenewal bool        `json:"auto_renewal"`
}

type SubscriptionCanceledResource struct {
	ID          string      `json:"id"`
	PlanID      string      `json:"plan_id"`
	Status      string      `json:"status"`
	CancelTime  string      `json:"cancel_time"`
	Subscriber  Subscriber  `json:"subscriber"`
	BillingInfo BillingInfo `json:"billing_info"`
}

type WebhookEvent struct {
	ID           string          `json:"id"`
	EventType    string          `json:"event_type"`
	EventVersion string          `json:"event_version"`
	CreateTime   time.Time       `json:"create_time"`
	ResourceType string          `json:"resource_type"`
	Summary      string          `json:"summary"`
	Resource     json.RawMessage `json:"resource"`
	Links        []Link          `json:"links"`
}

func CustomID2AppUserOrderID(customID string) (string, string, string, error) {
	ids := strings.Split(customID, "@")
	if len(ids) != 3 {
		return "", "", "", wlog.Errorf("invalid customid")
	}
	return ids[0], ids[1], ids[2], nil
}
