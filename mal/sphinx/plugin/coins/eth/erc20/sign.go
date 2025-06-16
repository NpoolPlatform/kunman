package erc20

import (
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/eth"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
)

func init() {
	register.RegisteTokenHandler(
		coins.Erc20,
		register.OpWalletNew,
		eth.CreateEthAccount,
	)
	register.RegisteTokenHandler(
		coins.Erc20,
		register.OpSign,
		eth.Msg,
	)
}
