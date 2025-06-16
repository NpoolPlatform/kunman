package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/big"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc"
	bsc_plugin "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/log"
	plugin_types "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Bep20,
		register.OpGetBalance,
		walletBalance,
	)
	register.RegisteTokenHandler(
		coins.Bep20,
		register.OpPreSign,
		bsc_plugin.PreSign,
	)
	register.RegisteTokenHandler(
		coins.Bep20,
		register.OpBroadcast,
		bsc_plugin.SendRawTransaction,
	)
	register.RegisteTokenHandler(
		coins.Bep20,
		register.OpSyncTx,
		bsc_plugin.SyncTxState,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypebinanceusd, bsc.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetbinanceusd, bsc.TxFailErr)
	if err != nil {
		panic(err)
	}
}

var (
	ErrContractAddrInvalid = errors.New("contract address is invalid")
	ErrAccountAddrInvalid  = errors.New("account address is invalid")
)

func walletBalance(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	wbReq := &plugin_types.WalletBalanceRequest{}
	err = json.Unmarshal(in, wbReq)
	if err != nil {
		return nil, err
	}

	bl, err := _walletBalance(ctx, wbReq.Address, tokenInfo)
	if err != nil {
		return nil, err
	}

	balance, ok := big.NewFloat(0).SetString(bl.String())
	if !ok {
		return nil, errors.New("convert balance string to float64 error")
	}

	balance.Quo(balance, big.NewFloat(math.Pow10(tokenInfo.Decimal)))
	f, exact := balance.Float64()
	if exact != big.Exact {
		log.Warnf("wallet balance transfer warning balance from->to %v-%v", balance.String(), f)
	}

	wbResp := &plugin_types.WalletBalanceResponse{
		Balance:    f,
		BalanceStr: balance.String(),
	}
	out, err = json.Marshal(wbResp)

	return out, err
}

func _walletBalance(ctx context.Context, addr string, tokenInfo *coins.TokenInfo) (*big.Int, error) {
	var ret *big.Int
	var err error
	client := bsc.Client()

	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		chainID, err := c.ChainID(ctx)
		if err != nil || chainID == nil {
			return true, err
		}
		contract := bsc.GetContract(chainID.Int64(), tokenInfo)
		ret, err = bep20Balance(ctx, contract, addr, c)
		if err != nil || ret == nil {
			return true, err
		}
		return false, err
	})

	return ret, err
}

func bep20Balance(ctx context.Context, contract, addr string, client bind.ContractBackend) (*big.Int, error) {
	if !common.IsHexAddress(contract) {
		return nil, ErrContractAddrInvalid
	}

	if !common.IsHexAddress(addr) {
		return nil, ErrAccountAddrInvalid
	}

	usdt, err := NewBEP20Token(common.HexToAddress(contract), client)
	if err != nil {
		return nil, err
	}

	return usdt.BalanceOf(&bind.CallOpts{
		Pending: true,
		Context: ctx,
	}, common.HexToAddress(addr))
}
