package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/rlp"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpGetBalance,
		walletBalance,
	)
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpEstimateGas,
		estimateGas,
	)
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpPreSign,
		PreSign,
	)
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpBroadcast,
		SendRawTransaction,
	)
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpSyncTx,
		SyncTxState,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypeethereum, eth.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetethereum, eth.TxFailErr)
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

	client := eth.Client()

	var bl *big.Int
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		bl, err = c.BalanceAt(ctx, common.HexToAddress(wbReq.Address), nil)
		if err == nil && bl != nil {
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

	balance := eth.ToEth(bl)

	wbResp := &ct.WalletBalanceResponse{
		Balance:    balance.InexactFloat64(),
		BalanceStr: balance.String(),
	}
	out, err = json.Marshal(wbResp)

	return out, err
}

func estimateGas(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	esGasReq := &proxy.GetEstimateGasRequest{}
	err = json.Unmarshal(in, esGasReq)
	if err != nil {
		return nil, err
	}

	client := eth.Client()
	mockFrom := common.HexToAddress("")
	mockTo := common.HexToAddress("")
	var mockValueInt int64
	mockValue := big.NewInt(mockValueInt)

	var gasLimit uint64
	var blockHeight uint64
	var gasPrice *big.Int
	var gasTips *big.Int
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		blockHeight, err = c.BlockNumber(ctx)
		if err != nil {
			return true, err
		}

		gasLimit, err = c.EstimateGas(ctx, ethereum.CallMsg{
			From:  mockFrom,
			To:    &mockTo,
			Value: mockValue,
		})
		if err != nil {
			return true, err
		}

		gasPrice, err = c.SuggestGasPrice(ctx)
		if err != nil {
			return true, err
		}

		gasTips, err = c.SuggestGasTipCap(ctx)
		if err != nil {
			return true, err
		}

		return false, err
	})

	if err != nil {
		return nil, err
	}

	gasLimitBig := big.NewInt(int64(gasLimit))
	estimateFee := big.NewInt(0).Mul(gasPrice, gasLimitBig)

	wbResp := &proxy.GetEstimateGasResponse{
		GasLimit:  fmt.Sprint(gasLimit),
		GasPrice:  gasPrice.String(),
		Fee:       eth.ToEth(estimateFee).String(),
		TipsPrice: gasTips.String(),
		BlockNum:  blockHeight,
	}
	return json.Marshal(wbResp)
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

	amountBig, ok := eth.ToWei(baseInfo.Value)
	if !ok {
		return nil, errors.New(eth.AmountInvalid)
	}

	client := eth.Client()

	var (
		chainID     *big.Int
		nonce       uint64
		estimateGas uint64
		bl          *big.Int
		gasPrice    *big.Int
		gasLimit    = uint64(21_000)
	)

	err = client.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) (bool, error) {
		chainID, err = cli.NetworkID(ctx)
		if err != nil || chainID == nil {
			return true, err
		}

		nonce, err = cli.PendingNonceAt(ctx, common.HexToAddress(baseInfo.From))
		if err != nil {
			return true, err
		}

		gasPrice, err = cli.SuggestGasPrice(ctx)
		if err != nil || gasPrice == nil {
			return true, err
		}

		to := common.HexToAddress(baseInfo.To)
		estimateGas, err = cli.EstimateGas(ctx, ethereum.CallMsg{
			From:  common.HexToAddress(baseInfo.From),
			To:    &to,
			Value: amountBig,
		})
		if err != nil {
			return true, err
		}

		bl, err = cli.BalanceAt(ctx, common.HexToAddress(baseInfo.From), nil)
		if err != nil {
			return true, err
		}

		return false, err
	})

	if err != nil {
		return nil, err
	}

	if bl == nil {
		return nil, errors.New(eth.GetInfoFailed)
	}

	if amountBig.Cmp(common.Big0) <= 0 {
		return nil, errors.New(eth.AmountInvalid)
	}

	estimateGasBig := big.NewInt(int64(estimateGas))
	estimateFee := big.NewInt(0).Mul(gasPrice, estimateGasBig)

	totalFunds := big.NewInt(0).Add(estimateFee, amountBig)

	if bl.Cmp(totalFunds) <= 0 {
		logger.Sugar().Warnf("from: %v, estimate fee + amount = totalFunds >= balance: %v + %v = %v >= %v",
			baseInfo.From,
			eth.ToEth(estimateFee),
			eth.ToEth(amountBig),
			eth.ToEth(totalFunds),
			eth.ToEth(bl),
		)
	}

	logger.Sugar().Infof("from: %v, estimate fee + amount = totalFunds: %v + %v = %v, balance: %v",
		baseInfo.From,
		eth.ToEth(estimateFee),
		eth.ToEth(amountBig),
		eth.ToEth(totalFunds),
		eth.ToEth(bl),
	)

	// build tx
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(baseInfo.To),
		amountBig,
		gasLimit,
		big.NewInt(gasPrice.Int64()),
		nil,
	)

	info := &eth.PreSignData{
		From:    baseInfo.From,
		Tx:      tx,
		ChainID: chainID,
	}

	out, err = json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return out, err
}

// SendRawTransaction eth/usdt
func SendRawTransaction(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) ([]byte, error) {
	signedData := &eth.SignedData{}
	err := json.Unmarshal(in, signedData)
	if err != nil {
		return nil, err
	}
	tx := new(types.Transaction)
	if err := rlp.Decode(bytes.NewReader(signedData.SignedTx), tx); err != nil {
		return nil, err
	}

	client := eth.Client()
	err = client.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		err = c.SendTransaction(ctx, tx)
		if err != nil && eth.TxFailErr(err) {
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

	return json.Marshal(broadcastedData)
}

// SyncTxState done(on chain) => true
func SyncTxState(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	broadcastedData := &ct.BroadcastInfo{}
	err = json.Unmarshal(in, broadcastedData)
	if err != nil {
		return nil, err
	}
	client := eth.Client()
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
