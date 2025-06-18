package types

import (
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
)

type PersistentAccount struct {
	*depositaccmwpb.Account
	CollectOutcoming *string
	Error            error
}
