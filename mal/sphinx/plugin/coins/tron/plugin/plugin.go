package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/NpoolPlatform/kunman/framework/logger"

	tronclient "github.com/Geapefurit/gotron-sdk/pkg/client"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"

	"github.com/Geapefurit/gotron-sdk/pkg/common"
	"github.com/Geapefurit/gotron-sdk/pkg/proto/api"
	"github.com/Geapefurit/gotron-sdk/pkg/proto/core"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpGetBalance,
		WalletBalance,
	)
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpPreSign,
		BuildTransaciton,
	)
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpBroadcast,
		BroadcastTransaction,
	)
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpSyncTx,
		SyncTxState,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetron, tron.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypettron, tron.TxFailErr)
	if err != nil {
		panic(err)
	}
}

// redefine Code ,because github.com/Geapefurit/gotron-sdk/pkg/proto/core/Tron.pb.go line 564 spelling err
const (
	TransactionInfoSUCCESS = 0
	TransactionInfoFAILED  = 1
)

func WalletBalance(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	wbReq := &ct.WalletBalanceRequest{}
	err = json.Unmarshal(in, wbReq)
	if err != nil {
		return in, err
	}

	if err := tron.ValidAddress(wbReq.Address); err != nil {
		return in, err
	}

	client := tron.Client()

	var bl int64
	err = client.WithClient(func(cli *tronclient.GrpcClient) (bool, error) {
		acc, err := cli.GetAccount(wbReq.Address)
		if err != nil && strings.Contains(err.Error(), tron.AddressNotActive) {
			bl = tron.EmptyTRX
			return false, nil
		}
		if err != nil {
			return true, err
		}
		if acc == nil {
			return true, errors.New(tron.GetAccountFailed)
		}
		bl = acc.GetBalance()
		return false, nil
	})
	if err != nil {
		return in, err
	}

	wbResp := &ct.WalletBalanceResponse{}
	f := tron.TRXToBigFloat(bl)

	wbResp.Balance, _ = f.Float64()
	wbResp.BalanceStr = f.Text('f', tron.TRXACCURACY)

	return json.Marshal(wbResp)
}

func BuildTransaciton(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	baseInfo := &ct.BaseInfo{}
	err = json.Unmarshal(in, baseInfo)
	if err != nil {
		return in, err
	}

	err = tron.ValidAddress(baseInfo.From)
	if err != nil {
		return in, fmt.Errorf("%v,%v", tron.AddressInvalid, err)
	}
	err = tron.ValidAddress(baseInfo.To)
	if err != nil {
		return in, fmt.Errorf("%v,%v", tron.AddressInvalid, err)
	}

	from := baseInfo.From
	to := baseInfo.To
	amount := tron.TRXToInt(baseInfo.Value)

	client := tron.Client()

	var txExtension *api.TransactionExtention
	err = client.WithClient(func(cli *tronclient.GrpcClient) (bool, error) {
		_, err := cli.GetAccount(from)
		if err != nil {
			return true, err
		}
		if tron.TxFailErr(err) {
			return false, err
		}
		txExtension, err = cli.Transfer(from, to, amount)
		if err != nil {
			return true, err
		}
		if txExtension == nil {
			return false, errors.New(tron.BuildTransactionFailed)
		}
		return false, err
	})
	if err != nil {
		return in, err
	}

	signTx := &tron.SignMsgTx{
		Base:        *baseInfo,
		TxExtension: txExtension,
	}

	return json.Marshal(signTx)
}

func BroadcastTransaction(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	bReq := &tron.BroadcastRequest{}
	err = json.Unmarshal(in, bReq)
	if err != nil {
		return in, err
	}

	client := tron.Client()
	transaction := bReq.TxExtension.Transaction
	bReq.TxExtension.GetTxid()
	var result *api.Return
	err = client.WithClient(func(cli *tronclient.GrpcClient) (bool, error) {
		result, err = cli.Broadcast(transaction)
		if err != nil && result != nil && result.GetCode() == api.Return_TRANSACTION_EXPIRATION_ERROR {
			return false, err
		}
		if err != nil || result == nil {
			return true, err
		}
		return false, err
	})

	if err != nil {
		return in, err
	}
	if result == nil {
		return in, fmt.Errorf("get result failed")
	}

	if api.Return_SUCCESS == result.Code {
		bResp := &ct.BroadcastInfo{TxID: common.BytesToHexString(bReq.TxExtension.GetTxid())}
		if result.Result {
			return json.Marshal(bResp)
		}
	}

	failCodes := []api.ReturnResponseCode{
		// api.Return_SUCCESS,
		api.Return_SIGERROR,
		api.Return_CONTRACT_VALIDATE_ERROR,
		api.Return_CONTRACT_EXE_ERROR,
		// api.Return_BANDWIDTH_ERROR=4,
		4,
		api.Return_DUP_TRANSACTION_ERROR,
		api.Return_TAPOS_ERROR,
		api.Return_TOO_BIG_TRANSACTION_ERROR,
		api.Return_TRANSACTION_EXPIRATION_ERROR,
		// api.Return_SERVER_BUSY,
		// api.Return_NO_CONNECTION,
		// api.Return_NOT_ENOUGH_EFFECTIVE_CONNECTION,
		api.Return_OTHER_ERROR,
	}
	for _, v := range failCodes {
		if v == result.Code {
			return in, env.ErrTransactionFail
		}
	}

	return in, errors.New(string(result.GetMessage()))
}

// done(on chain) => true
func SyncTxState(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	syncReq := &ct.SyncRequest{}
	err = json.Unmarshal(in, syncReq)
	if err != nil {
		logger.Sugar().Errorw("SyncTxState", "Req", syncReq, "Error", err)
		return in, err
	}
	client := tron.Client()

	var txInfo *core.TransactionInfo
	err = client.WithClient(func(cli *tronclient.GrpcClient) (bool, error) {
		txInfo, err = cli.GetTransactionInfoByID(syncReq.TxID)
		if err != nil {
			logger.Sugar().Errorw("SyncTxState", "Req", syncReq, "Error", err)
			return true, err
		}
		return false, err
	})

	if txInfo == nil || err != nil {
		logger.Sugar().Errorw("SyncTxState", "Req", syncReq, "Info", txInfo, "Error", err)
		return in, env.ErrWaitMessageOnChain
	}

	if txInfo.GetResult() != TransactionInfoSUCCESS {
		logger.Sugar().Errorw("SyncTxState", "Req", syncReq, "Info", txInfo, "Result", txInfo.GetResult())
		return in, env.ErrTransactionFail
	}

	if txInfo.Receipt.GetResult() != core.Transaction_Result_SUCCESS && txInfo.Receipt.GetResult() != core.Transaction_Result_DEFAULT {
		logger.Sugar().Errorw("SyncTxState", "Req", syncReq, "Info", txInfo, "Result", txInfo.GetResult())
		return in, env.ErrTransactionFail
	}

	syncResp := &ct.SyncResponse{ExitCode: 0}
	return json.Marshal(syncResp)
}
