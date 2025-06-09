package types

// request url
const (
	RevenueDistributionInfo   = "/v2/revenue/distribution/info"
	RevenueDistributionAdd    = "/v2/revenue/distribution/add"
	RevenueDistributionDelete = "/v2/revenue/distribution/delete"
)

// public struct
type RevenueDistribution struct {
	ID          int64   `json:"id,string"`
	Currency    string  `json:"currency"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	Description string  `json:"description"`
	Proportion  float64 `json:"proportion"`
}

// request and response
type RevenueDistributionInfoReq struct {
	Currency    string `json:"currency"`
	Distributor string `json:"distributor"`
	Recipient   string `json:"recipient"`
}

type RevenueDistributionInfoResp struct {
	Data []RevenueDistribution `json:"data"`
}

type RevenueDistributionAddReq struct {
	Currency    string  `json:"currency"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	Description string  `json:"description"`
	Proportion  float64 `json:"proportion"`
}

type RevenueDistributionAddResp struct {
	ID          string  `json:"id"`
	Currency    string  `json:"currency"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	Description string  `json:"description"`
	Proportion  float64 `json:"proportion"`
}

type RevenueDistributionDeleteReq struct {
	ID          int64  `json:"id"`
	Distributor string `json:"distributor"`
}

type RevenueDistributionDeleteResp struct {
	Reason string `json:"reason"`
}
