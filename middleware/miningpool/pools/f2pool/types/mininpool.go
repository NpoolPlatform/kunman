package types

// request url
const (
	BlocksPaging    = "/v2/blocks/paging"
	BlocksDateRange = "/v2/blocks/date_range"
)

// public struct

type LuckyEntry struct {
	Key   string  `json:"key"`
	Value float32 `json:"value"`
}

type Block struct {
	Height          int64   `json:"height"`
	BlockHash       string  `json:"block_hash"`
	BaseReward      float64 `json:"base_reward"`
	TotalReward     float64 `json:"total_reward"`
	Timestamp       int64   `json:"timestamp"`
	SettleTimestamp int64   `json:"settle_timestamp"`
}

// request and response

type BlocksPagingReq struct {
	Page     int32  `json:"page"`
	Pagesize int32  `json:"pagesize"`
	Currency string `json:"currency"`
}
type BlocksPagingResp struct {
	Lucky     []LuckyEntry `json:"lucky"`
	Blocklist []Block      `json:"blocklist"`
	Sum       int32        `json:"sum"`
}

type BlocksDateRangeReq struct {
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Currency  string `json:"currency"`
}
type BlocksDateRangeResp struct {
	Lucky     []LuckyEntry `json:"lucky"`
	Blocklist []Block      `json:"blocklist"`
	Sum       int32        `json:"sum"`
}
