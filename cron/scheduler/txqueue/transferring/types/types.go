package types

import (
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
)

type PersistentTx struct {
	*txmwpb.Tx
	NewTxState basetypes.TxState
	TxExtra    string
	TxCID      *string
}
