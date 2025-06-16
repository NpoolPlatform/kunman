package trc20

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	tronclient "github.com/Geapefurit/gotron-sdk/pkg/client"
	"github.com/Geapefurit/gotron-sdk/pkg/proto/api"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron"
	tron_plugin "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpGetBalance,
		WalletBalance,
	)
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpPreSign,
		BuildTransaciton,
	)
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpBroadcast,
		tron_plugin.BroadcastTransaction,
	)
	register.RegisteTokenHandler(
		coins.Trc20,
		register.OpSyncTx,
		tron_plugin.SyncTxState,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypeusdttrc20, tron.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetusdttrc20, tron.TxFailErr)
	if err != nil {
		panic(err)
	}
}

func WalletBalance(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	wbReq := &ct.WalletBalanceRequest{}
	err = json.Unmarshal(in, wbReq)
	if err != nil {
		return nil, err
	}

	v, ok := env.LookupEnv(env.ENVCOINNET)
	if !ok {
		return nil, env.ErrEVNCoinNet
	}
	if !coins.CheckSupportNet(v) {
		return nil, env.ErrEVNCoinNetValue
	}

	contract := tron.USDTContract(v)
	err = tron.ValidAddress(contract)
	if err != nil {
		return nil, fmt.Errorf("contract %v, %v, %v", contract, tron.AddressInvalid, err)
	}

	bl := tron.EmptyTRC20
	if err := tron.ValidAddress(wbReq.Address); err != nil {
		return nil, err
	}

	client := tron.Client()
	err = client.WithClient(func(c *tronclient.GrpcClient) (bool, error) {
		bl, err = c.TRC20ContractBalance(wbReq.Address, contract)
		if err != nil && strings.Contains(err.Error(), tron.AddressNotActive) {
			bl = tron.EmptyTRC20
			return false, nil
		}
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	f := tron.TRC20ToBigFloat(bl)
	wbResp := &ct.WalletBalanceResponse{}

	wbResp.Balance, _ = f.Float64()
	wbResp.BalanceStr = f.Text('f', tron.TRC20ACCURACY)

	out, err = json.Marshal(wbResp)

	return out, err
}

func BuildTransaciton(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	baseInfo := &ct.BaseInfo{}
	err = json.Unmarshal(in, baseInfo)
	if err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(baseInfo.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	err = tron.ValidAddress(baseInfo.From)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", tron.AddressInvalid, err)
	}

	err = tron.ValidAddress(baseInfo.To)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", tron.AddressInvalid, err)
	}

	contract := tron.USDTContract(baseInfo.ENV)
	err = tron.ValidAddress(contract)
	if err != nil {
		return nil, fmt.Errorf("contract %v, %v, %v", contract, tron.AddressInvalid, err)
	}

	var txExtension *api.TransactionExtention
	client := tron.Client()
	err = client.WithClient(func(c *tronclient.GrpcClient) (bool, error) {
		txExtension, err = c.TRC20Send(
			baseInfo.From,
			baseInfo.To,
			contract,
			tron.TRC20ToBigInt(baseInfo.Value),
			tron.TRC20FeeLimit,
		)
		return false, err
	})
	if err != nil {
		return nil, err
	}
	signTx := &tron.SignMsgTx{
		Base:        *baseInfo,
		TxExtension: txExtension,
	}

	return json.Marshal(signTx)
}
