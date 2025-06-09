package types

// request url
const (
	AssetsBalance          = "/v2/assets/balance"
	AssetsTransactionsList = "/v2/assets/transactions/list"
	AssetsSettleModeSwitch = "/v2/assets/settle_mode/switch"
)

// public struct
type BalanceInfo struct {
	Balance              float64 `json:"balance"`
	Paid                 float64 `json:"paid"`
	TotalIncome          float64 `json:"total_income"`
	YesterdayIncome      float64 `json:"yesterday_income"`
	EstimatedTodayIncome float64 `json:"estimated_today_income"`
}

type MiningExtra struct {
	MiningDate  int64   `json:"mining_date"`
	SettleDate  int64   `json:"settle_date"`
	Pps         float64 `json:"pps"`
	PpsFeeRatio float32 `json:"pps_fee_rate"`
	TxFee       float64 `json:"tx_fee"`
	TxFeeRatio  float32 `json:"tx_fee_rate"`
	HashRate    float64 `json:"hash_rate"`
}

type PayoutExtra struct {
	Value      float64 `json:"value"`
	Address    string  `json:"address"`
	TxID       string  `json:"tx_id"`
	PaidTime   int64   `json:"paid_time"`
	WalletType string  `json:"wallet_type"`
	Tip        float64 `json:"tip"`
}

type TransactionItem struct {
	ID             int64       `json:"id"`
	Type           string      `json:"type"`
	ChangedBalance float64     `json:"changed_balance"`
	CreatedAt      int64       `json:"created_at"`
	MiningExtra    MiningExtra `json:"mining_extra"`
	PayoutExtra    PayoutExtra `json:"payout_extra"`
}

// request and response

type AssetsBalanceReq struct {
	Currency       string `json:"currency"`
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
}

type AssetsBalanceResp struct {
	BalanceInfo BalanceInfo `json:"balance_info"`
}

type AssetsTransactionsListReq struct {
	Currency       string `json:"currency"`
	MiningUserName string `json:"mining_user_name"`
	Address        string `json:"address"`
	TxType         string `json:"type"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
}

type AssetsTransactionsListResp struct {
	Transactions []TransactionItem `json:"transactions"`
}

type AssetsSettleModeSwitchReq struct {
	Currency       string `json:"currency"`
	Mode           string `json:"mode"`
	MiningUserName string `json:"mining_user_name"`
	ActivatedAt    int64  `json:"activated_at"`
}

type AssetsSettleModeSwitchResp struct {
	Mode        string `json:"mode"`
	ActivatedAt int64  `json:"activated_at"`
}
