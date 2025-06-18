package types

import (
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
)

type PersistentTx struct {
	*txmwpb.Tx
}
