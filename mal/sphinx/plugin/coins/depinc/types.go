package depinc

import (
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type SignMsgTx struct {
	ct.BaseInfo
	// from address script
	PayToAddrScript []byte
	// all used utxo amount
	Amounts []btcutil.Amount
	MsgTx   *wire.MsgTx
}
