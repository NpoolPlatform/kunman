package bsc

import "github.com/NpoolPlatform/kunman/message/sphinx/plugin"

type PreSignData struct {
	CoinType   plugin.CoinType `json:"coin_type"`
	From       string                `json:"from"`
	To         string                `json:"to"`
	Value      float64               `json:"value"`
	ChainID    int64                 `json:"chain_id"`
	Nonce      uint64                `json:"nonce"`
	GasPrice   int64                 `json:"gas_price"`
	ContractID string                `json:"contract_id"`
	GasLimit   int64                 `json:"gas_limit"`
}

type SignedData struct {
	SignedTx []byte `json:"signed_tx"`
}

type BroadcastedData struct {
	TxHash string `json:"tx_hash"`
}
