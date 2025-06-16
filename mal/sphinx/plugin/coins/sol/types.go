package sol

import (
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
)

type SignMsgTx struct {
	BaseInfo        ct.BaseInfo `json:"base_info"`
	RecentBlockHash string      `json:"recent_block_hash"`
}

type BroadcastRequest struct {
	Signature []byte `json:"signature"`
}
