package trc20

import (
	"context"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	tron "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/sign"
)

func init() {
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpWalletNew,
		CreateTrc20Account,
	)
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpSign,
		SignTrc20MSG,
	)
}

const s3KeyPrxfix = "usdttrc20/"

func SignTrc20MSG(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	return tron.SignTronMSG(ctx, s3KeyPrxfix, in)
}

func CreateTrc20Account(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	return tron.CreateTronAccount(ctx, s3KeyPrxfix, in)
}
