package types

// request url
const (
	HashRateInfo          = "/v2/hash_rate/info"
	HashRateInfoList      = "/v2/hash_rate/info_list"
	HashRateHistory       = "/v2/hash_rate/history"
	HashRateWorkerList    = "/v2/hash_rate/worker/list"
	HashRateWorkerHistory = "/v2/hash_rate/worker/history"
)

// public struct

type HashRateInfoStruct struct {
	Name             string  `json:"name"`
	HashRate         float64 `json:"hash_rate"`
	H1HashRate       float64 `json:"h1_hash_rate"`
	H24HashRate      float64 `json:"h24_hash_rate"`
	H1StaleHashRate  float64 `json:"h1_stale_hash_rate"`
	H24StaleHashRate float64 `json:"h24_stale_hash_rate"`
	H24DelayHashRate float64 `json:"h24_delay_hash_rate"`
	LocalHashRate    float64 `json:"local_hash_rate"`
	H24LocalHashRate float64 `json:"h24_local_hash_rate"`
}

type MiningHistoryItem struct {
	Timestamp     int64   `json:"timestamp"`
	HashRate      float64 `json:"hash_rate"`
	StaleHashRate float64 `json:"stale_hash_rate"`
	DelayHashRate float64 `json:"delay_hash_rate"`
	LocalHashRate float64 `json:"local_hash_rate"`
	NormalReward  float64 `json:"normal_reward"`
	DelayReward   float64 `json:"delay_reward"`
	OnlineMiners  uint32  `json:"online_miners"`
}
type UserMiningReq struct {
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	Currency       string `json:"currency"`
}

// TODO: consider define in proto
type WorkerStatus int

const (
	WorkerStatusOnline  = 0
	WorkerStatusOffline = 1
	WorkerStatusInvalid = 2
)

type WorkerMiningInfo struct {
	HashRateInfo HashRateInfoStruct `json:"hash_rate_info"`
	LastShareAt  int64              `json:"last_share_at"`
	Status       WorkerStatus       `json:"status"`
	Host         string             `json:"host"`
}
type HashRateInfoReq struct {
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	Currency       string `json:"currency"`
}
type HashRateInfoResp struct {
	Info     HashRateInfoStruct `json:"info"`
	Currency string             `json:"currency"`
}
type HashRateInfoListReq struct {
	Reqs []UserMiningReq `json:"reqs"`
}
type HashRateInfoListResp struct {
	Info []HashRateInfoStruct `json:"info"`
}
type HashRateHistoryReq struct {
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	Currency       string `json:"currency"`
	Interval       int64  `json:"interval"`
	Duration       int64  `json:"duration"`
}
type HashRateHistoryResp struct {
	History []MiningHistoryItem `json:"history"`
}
type HashRateWorkerListReq struct {
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	Currency       string `json:"currency"`
}
type HashRateWorkerListResp struct {
	Workers []WorkerMiningInfo `json:"workers"`
}
type HashRateWorkerHistoryReq struct {
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	Currency       string `json:"currency"`
	WorkerName     string `json:"worker_name"`
	Interval       int64  `json:"interval"`
	Duration       int64  `json:"duration"`
}
type HashRateWorkerHistoryResp struct {
	History []MiningHistoryItem `json:"history"`
}

// request url
const (
	HashRateDistributionInfo        = "/v2/hash_rate/distribution/info"
	HashRateDistributionOrders      = "/v2/hash_rate/distribution/orders"
	HashRateDistributionSettlements = "/v2/hash_rate/distribution/settlements"
	HashRateDistribute              = "/v2/hash_rate/distribute"
	HashRateDistributionTerminate   = "/v2/hash_rate/distribution/terminate"
)

// public struct
type HashRateDistributionOrder struct {
	ID          int64   `json:"id"`
	StartDate   int64   `json:"start_date"`
	EndDate     int64   `json:"end_date"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	HashRate    float64 `json:"hash_rate"`
	TerminateAt int64   `json:"terminate_at"`
}

type HashRateDistributionSettlement struct {
	ID        int64                     `json:"id"`
	Order     HashRateDistributionOrder `json:"order"`
	Hashes    float64                   `json:"hashes"`
	Income    float64                   `json:"income"`
	Timestamp int64                     `json:"timestamp"`
}

// request and response

type HashRateDistributionInfoReq struct {
	Currency    string  `json:"currency"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	StartDate   int64   `json:"start_date"`
	EndDate     int64   `json:"end_date"`
	HashRate    float64 `json:"hash_rate"`
}
type HashRateDistributionInfoResp struct {
	CurrentHashRate   float64 `json:"current_hash_rate"`
	SoldHashRate      float64 `json:"sold_hash_rate"`
	RemainingHashRate float64 `json:"remaining_hash_rate"`
}

type HashRateDistributionOrdersReq struct {
	Currency    string `json:"currency"`
	Distributor string `json:"distributor"`
}
type HashRateDistributionOrdersResp struct {
	Orders HashRateDistributionOrder `json:"orders"`
}

type HashRateDistributionSettlementsReq struct {
	Currency    string `json:"currency"`
	Distributor string `json:"distributor"`
}
type HashRateDistributionSettlementsResp struct {
	Settlements HashRateDistributionSettlement `json:"settlements"`
}

type HashRateDistributeReq struct {
	Currency    string  `json:"currency"`
	Distributor string  `json:"distributor"`
	Recipient   string  `json:"recipient"`
	StartDate   int64   `json:"start_date"`
	EndDate     int64   `json:"end_date"`
	HashRate    float64 `json:"hash_rate"`
}
type HashRateDistributeResp struct {
	Order HashRateDistributionOrder `json:"order"`
}

type HashRateDistributionTerminateReq struct {
	OrderID int64 `json:"order_id"`
}
type HashRateDistributionTerminateResp struct {
	Reason string `json:"reason"`
}
