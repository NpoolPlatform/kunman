package types

// request url
const (
	MiningUserGet                       = "/v2/mining_user/get"
	MiningUserAdd                       = "/v2/mining_user/add"
	MiningUserList                      = "/v2/mining_user/list"
	MiningUserWalletUpdate              = "/v2/mining_user/wallet/update"
	MiningUserReadOnlyPageAdd           = "/v2/mining_user/read_only_page/add"
	MiningUserReadOnlyPageDelete        = "/v2/mining_user/read_only_page/delete"
	MiningUserBalanceWithdraw           = "/v2/mining_user/balance/withdraw"
	MiningUserPaymentPause              = "/v2/mining_user/payment/pause"
	MiningUserPaymentResume             = "/v2/mining_user/payment/resume"
	MiningUserLeastTransferAmountUpdate = "/v2/mining_user/least_transfer_amount/update"
)

// error response
type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// public struct

type ReadOnlyPage struct {
	Key         string `json:"key"`
	PageName    string `json:"page_name"`
	Permissions string `json:"permissions"`
}
type Permission string

const (
	MinerInfo Permission = "1"
	PayInfo   Permission = "2"
	NoPayInfo Permission = "3"
)

type Wallet struct {
	Currency            string  `json:"currency"`
	Address             string  `json:"address"`
	LeastTransferAmount float64 `json:"least_transfer_amount,string"`
}

type ExtraEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MiningUserInfo struct {
	MiningUserName string         `json:"mining_user_name"`
	Pages          []ReadOnlyPage `json:"pages"`
	Wallets        []Wallet       `json:"wallets"`
	Description    string         `json:"description"`
}

type WalletParams struct {
	MiningUserName string   `json:"mining_user_name"`
	Wallets        []Wallet `json:"wallets"`
}

// type ResultsEntry struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

type WalletResp struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
	Result   string `json:"result"`
	Msg      string `json:"msg"`
}

type MiningUserWalletResp struct {
	MiningUserName string       `json:"mining_user_name"`
	WalletResult   []WalletResp `json:"wallet_result"`
}

type ResultsEntry struct {
	Key   string               `json:"key"`
	Value MiningUserWalletResp `json:"value"`
}

// request and response
type MiningUserGetReq struct {
	MiningUserName string `json:"mining_user_name"`
}
type MiningUserGetResp struct {
	MiningUserName string         `json:"mining_user_name"`
	Pages          []ReadOnlyPage `json:"pages"`
	Wallets        []Wallet       `json:"wallets"`
	Description    string         `json:"description"`
}

type MiningUserAddReq struct {
	MiningUserName string `json:"mining_user_name"`
}
type MiningUserAddResp struct {
	MiningUserName string         `json:"mining_user_name"`
	Pages          []ReadOnlyPage `json:"pages"`
	Wallets        []Wallet       `json:"wallets"`
	Description    string         `json:"description"`
}

type MiningUserListReq struct{}

type MiningUserListResp struct {
	MiningUserList []MiningUserInfo `json:"mining_user_list"`
}

type MiningUserWalletUpdateReq struct {
	Params []WalletParams `json:"params"`
}
type MiningUserWalletUpdateResp struct {
	// Results  []ResultsEntry                    `json:"results"`
	Results map[string]MiningUserWalletResp `json:"results"`
}

type MiningUserReadOnlyPageAddReq struct {
	MiningUserName string `json:"mining_user_name"`
	PageName       string `json:"page_name"`
	Permissions    string `json:"permissions"`
}
type MiningUserReadOnlyPageAddResp struct {
	MiningUserName string       `json:"mining_user_name"`
	Page           ReadOnlyPage `json:"page"`
}

type MiningUserReadOnlyPageDeleteReq struct {
	MiningUserName string `json:"mining_user_name"`
	Key            string `json:"key"`
}
type MiningUserReadOnlyPageDeleteResp struct {
	MiningUserName string `json:"mining_user_name"`
}

type MiningUserBalanceWithdrawReq struct {
	MiningUserName string `json:"mining_user_name"`
	Currency       string `json:"currency"`
}
type MiningUserBalanceWithdrawResp struct {
	PaidTime int64 `json:"paid_time"`
}

type MiningUserPaymentPauseReq struct {
	Currency        string   `json:"currency"`
	MiningUserNames []string `json:"mining_user_names"`
}
type MiningUserPaymentPauseResp struct {
	Results map[string]string `json:"results"`
}

type MiningUserPaymentResumeReq struct {
	Currency        string   `json:"currency"`
	MiningUserNames []string `json:"mining_user_names"`
}
type MiningUserPaymentResumeResp struct {
	Results map[string]string `json:"results"`
}

type MiningUserLeastTransferAmountUpdateReq struct {
	MiningUserName      string  `json:"mining_user_name"`
	Currency            string  `json:"currency"`
	LeastTransferAmount float64 `json:"least_transfer_amount"`
}
type MiningUserLeastTransferAmountUpdateResp struct {
	LeastTransferAmount float64 `json:"least_transfer_amount"`
}
