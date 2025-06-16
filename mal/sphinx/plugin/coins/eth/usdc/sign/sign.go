package sign

import (
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	eth_sign "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/eth"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
)

func init() {
	// main
	register.RegisteTokenHandler(
		coins.USDC,
		register.OpWalletNew,
		eth_sign.CreateEthAccount,
	)
	register.RegisteTokenHandler(
		coins.USDC,
		register.OpSign,
		eth_sign.Msg,
	)
}
