package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/depinc"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/depinc/depc/rpcclient"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/log"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

// here register plugin func
func init() {
	register.RegisteTokenHandler(
		coins.Depinc,
		register.OpGetBalance,
		walletBalance,
	)
	register.RegisteTokenHandler(
		coins.Depinc,
		register.OpPreSign,
		preSign,
	)
	register.RegisteTokenHandler(
		coins.Depinc,
		register.OpBroadcast,
		broadcast,
	)
	register.RegisteTokenHandler(
		coins.Depinc,
		register.OpSyncTx,
		syncTx,
	)

	err := register.RegisteAbortFuncErr(plugin.CoinType_CoinTypedepinc, depinc.TxFailErr)
	if err != nil {
		panic(err)
	}

	err = register.RegisteAbortFuncErr(plugin.CoinType_CoinTypetdepinc, depinc.TxFailErr)
	if err != nil {
		panic(err)
	}
}

// walletBalance ..
func walletBalance(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := ct.WalletBalanceRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	v, ok := env.LookupEnv(env.ENVCOINNET)
	if !ok {
		return nil, env.ErrEVNCoinNet
	}
	if !coins.CheckSupportNet(v) {
		return nil, env.ErrEVNCoinNetValue
	}

	_addr, err := btcutil.DecodeAddress(info.Address, depinc.DEPCNetMap[v])
	if err != nil {
		return nil, err
	}

	client := depinc.Client()
	err = client.WithClient(ctx, func(cli *rpcclient.Client) (bool, error) {
		err := cli.ImportAddressRescan(info.Address, "", false)
		if err != nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	var unspents []btcjson.ListUnspentResult
	err = client.WithClient(ctx, func(cli *rpcclient.Client) (bool, error) {
		unspents, err = cli.ListUnspentMinMaxAddresses(depinc.DefaultMinConfirms, depinc.DefaultMaxConfirms, []btcutil.Address{_addr})
		if err != nil || unspents == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	accountAmount := float64(0)
	for _, sp := range unspents {
		if sp.Address == info.Address {
			// TODO: bug not accuracy, use bigint
			accountAmount += sp.Amount
		}
	}

	balance, err := btcutil.NewAmount(accountAmount)
	if err != nil {
		return nil, err
	}

	_out := ct.WalletBalanceResponse{
		Balance:    balance.ToBTC(),
		BalanceStr: strconv.FormatFloat(balance.ToUnit(btcutil.AmountBTC), 'f', -int(btcutil.AmountBTC+8), 64), // process reference is ‘balance.String()’
	}

	return json.Marshal(_out)
}

// preSign ..
func preSign(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) ([]byte, error) {
	info := ct.BaseInfo{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	if info.From == "" {
		return nil, env.ErrAddressInvalid
	}
	if info.To == "" {
		return nil, env.ErrAddressInvalid
	}
	if info.Value <= 0 {
		return nil, env.ErrAmountInvalid
	}

	var (
		from   = info.From
		to     = info.To
		amount = info.Value
	)

	fromAddr, err := btcutil.DecodeAddress(from, depinc.DEPCNetMap[info.ENV])
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAddressInvalid.Error(), err)
	}

	client := depinc.Client()

	var listUnspentResult []btcjson.ListUnspentResult
	err = client.WithClient(ctx, func(cli *rpcclient.Client) (bool, error) {
		listUnspentResult, err = cli.ListUnspentMinMaxAddresses(
			depinc.DefaultMinConfirms,
			depinc.DefaultMaxConfirms,
			[]btcutil.Address{fromAddr},
		)
		if err != nil || listUnspentResult == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	// new transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)

	// TODO: optimizing the utxos selected
	enoughUTXOAmount := float64(0)
	// sign and check need this info
	// btcutil.Amount is alias of int64
	inputAccount := make([]btcutil.Amount, 0)
	amountflag := false
	for _, txIn := range listUnspentResult {
		txHash, err := chainhash.NewHashFromStr(txIn.TxID)
		if err != nil {
			return nil, err
		}
		// get in amount
		inAmount, err := btcutil.NewAmount(txIn.Amount)
		if err != nil {
			return nil, fmt.Errorf("%v,%v", env.ErrAmountInvalid.Error(), err)
		}

		inputAccount = append(inputAccount, inAmount)
		msgTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txHash, txIn.Vout), nil, nil))

		// util sufficient funds are available
		enoughUTXOAmount += txIn.Amount
		if enoughUTXOAmount >= amount+depinc.DEPCGas {
			amountflag = true
			break
		}
	}

	if !amountflag {
		// TODO: think how to use same error
		log.Errorf(
			"insufficient balance: total: %v, transfer: %v, gas: %v",
			enoughUTXOAmount,
			amount,
			depinc.DEPCGas,
		)
		return nil, env.ErrInsufficientBalance
	}

	fromScript, err := txscript.PayToAddrScript(fromAddr)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAddressInvalid, err)
	}

	// cal change amount
	changeAmount, err := btcutil.NewAmount(enoughUTXOAmount - amount - depinc.DEPCGas)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAmountInvalid, err)
	}

	if changeAmount.ToBTC() > 0 {
		msgTx.AddTxOut(wire.NewTxOut(int64(changeAmount), fromScript))
	}

	toAddr, err := btcutil.DecodeAddress(to, depinc.DEPCNetMap[info.ENV])
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAddressInvalid, err)
	}

	toScript, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAddressInvalid, err)
	}

	outAmount, err := btcutil.NewAmount(amount)
	if err != nil {
		return nil, fmt.Errorf("%v,%v", env.ErrAmountInvalid, err)
	}

	msgTx.AddTxOut(wire.NewTxOut(int64(outAmount), toScript))

	_out := depinc.SignMsgTx{
		BaseInfo:        info,
		PayToAddrScript: fromScript,
		Amounts:         inputAccount,
		MsgTx:           msgTx,
	}

	return json.Marshal(_out)
}

// SendRawTransaction ..
func broadcast(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := &wire.MsgTx{}
	if err := json.Unmarshal(in, info); err != nil {
		return nil, err
	}

	client := depinc.Client()
	var _hash *chainhash.Hash
	err = client.WithClient(ctx, func(cli *rpcclient.Client) (bool, error) {
		_hash, err = cli.SendRawTransaction(info, false)
		if err != nil || _hash == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	_out := ct.SyncRequest{
		TxID: _hash.String(),
	}

	return json.Marshal(_out)
}

// syncTx ..
func syncTx(_ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := ct.SyncRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	var txHash *chainhash.Hash
	txHash, err = chainhash.NewHashFromStr(info.TxID)
	if err != nil {
		return nil, err
	}

	client := depinc.Client()
	var transactionResult *btcjson.GetTransactionResult
	err = client.WithClient(_ctx, func(cli *rpcclient.Client) (bool, error) {
		transactionResult, err = cli.GetTransaction(txHash)
		if err != nil || transactionResult == nil {
			return true, err
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}

	if transactionResult.Confirmations < depinc.DefaultMinConfirms {
		return nil, depinc.ErrWaitMessageOnChainMinConfirms
	}

	sResp := &ct.SyncResponse{ExitCode: 0}
	return json.Marshal(sResp)
}
