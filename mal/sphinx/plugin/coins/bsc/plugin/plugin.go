package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/big"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/log"

	bsc "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/rlp"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Binancecoin,
		register.OpGetBalance,
		walletBalance,
	)
	register.RegisteTokenHandler(
		coins.Binancecoin,
		register.OpPreSign,
		PreSign,
	)
	register.RegisteTokenHandler(
		coins.Binancecoin,
		register.OpBroadcast,
		SendRawTransaction,
	)
	register.RegisteTokenHandler(
		coins.Binancecoin,
		register.OpSyncTx,
		SyncTxState,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypebinancecoin, bsc.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetbinancecoin, bsc.TxFailErr)
	if err != nil {
		panic(err)
	}
}

func walletBalance(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	wbReq := &ct.WalletBalanceRequest{}
	err = json.Unmarshal(in, wbReq)
	if err != nil {
		return nil, err
	}

	if !common.IsHexAddress(wbReq.Address) {
		return nil, env.ErrAddressInvalid
	}

	client := bsc.Client()
	var bl *big.Int
	err = client.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) (bool, error) {
		bl, err = cli.BalanceAt(ctx, common.HexToAddress(wbReq.Address), nil)
		if err != nil || bl == nil {
			return true, err
		}
		return false, err
	})
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

	wbResp := &ct.WalletBalanceResponse{
		Balance:    f,
		BalanceStr: balance.String(),
	}
	out, err = json.Marshal(wbResp)

	return out, err
}

func PreSign(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	baseInfo := &ct.BaseInfo{}
	err = json.Unmarshal(in, baseInfo)
	if err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(baseInfo.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	if !common.IsHexAddress(baseInfo.From) {
		return nil, env.ErrAddressInvalid
	}

	client := bsc.Client()

	var chainID *big.Int
	err = client.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) (bool, error) {
		chainID, err = cli.NetworkID(ctx)
		if err != nil || chainID == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	var nonce uint64
	err = client.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) (bool, error) {
		nonce, err = cli.PendingNonceAt(ctx, common.HexToAddress(baseInfo.From))
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	var gasPrice *big.Int
	err = client.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) (bool, error) {
		gasPrice, err = cli.SuggestGasPrice(ctx)
		if err != nil || gasPrice == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	info := &bsc.PreSignData{
		ChainID:  chainID.Int64(),
		Nonce:    nonce,
		GasPrice: gasPrice.Int64(),
		From:     baseInfo.From,
		To:       baseInfo.To,
		Value:    baseInfo.Value,
	}

	switch baseInfo.CoinType {
	case plugin.CoinType_CoinTypebinancecoin, plugin.CoinType_CoinTypetbinancecoin:
		info.GasLimit = 21_000
	case plugin.CoinType_CoinTypebinanceusd,
		plugin.CoinType_CoinTypetbinanceusd,
		plugin.CoinType_CoinTypeusdtbep20,
		plugin.CoinType_CoinTypetusdtbep20:
		info.ContractID = bsc.GetContract(chainID.Int64(), tokenInfo)
		info.GasLimit = 300_000
	}

	return json.Marshal(info)
}

// SendRawTransaction bsc
func SendRawTransaction(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	signedData := &bsc.SignedData{}
	err = json.Unmarshal(in, signedData)
	if err != nil {
		return nil, err
	}

	tx := new(types.Transaction)

	if err := rlp.Decode(bytes.NewReader(signedData.SignedTx), tx); err != nil {
		return nil, err
	}

	client := bsc.Client()
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		err = c.SendTransaction(ctx, tx)
		if err != nil && bsc.TxFailErr(err) {
			return false, err
		}
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	broadcastedData := ct.BroadcastInfo{
		TxID: tx.Hash().Hex(),
	}
	out, err = json.Marshal(broadcastedData)
	return out, err
}

// done(on chain) => true
func SyncTxState(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	broadcastedData := &ct.BroadcastInfo{}
	err = json.Unmarshal(in, broadcastedData)
	if err != nil {
		return nil, err
	}

	client := bsc.Client()
	var isPending bool
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		_, isPending, err = c.TransactionByHash(ctx, common.HexToHash(broadcastedData.TxID))
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}
	if isPending {
		return nil, env.ErrWaitMessageOnChain
	}

	var receipt *types.Receipt
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		receipt, err = c.TransactionReceipt(ctx, common.HexToHash(broadcastedData.TxID))
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}
	if receipt == nil {
		return nil, env.ErrWaitMessageOnChain
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		sResp := &ct.SyncResponse{ExitCode: 0}
		out, err = json.Marshal(sResp)
		if err != nil {
			return nil, err
		}

		return out, nil
	}

	return nil, env.ErrTransactionFail
}
