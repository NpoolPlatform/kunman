package eth

import (
	"math/big"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/ethereum/go-ethereum/core/types"
)

// ----------use for transaction process
type PreSignData struct {
	CoinType plugin.CoinType `json:"coin_type"`
	From     string                `json:"from"`
	ChainID  *big.Int              `json:"chain_id"`
	Tx       *types.Transaction    `json:"tx"`
}

type SignedData struct {
	SignedTx []byte `json:"signed_tx"`
}

type BroadcastedData struct {
	TxHash string `json:"tx_hash"`
}
