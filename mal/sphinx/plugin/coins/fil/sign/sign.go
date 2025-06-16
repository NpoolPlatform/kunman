package sign

import (
	"context"
	"encoding/json"

	"github.com/NpoolPlatform/kunman/framework/oss"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/fil"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	ftypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

func init() {
	register.RegisteTokenHandler(
		coins.Filecoin,
		register.OpWalletNew,
		createAccount,
	)
	register.RegisteTokenHandler(
		coins.Filecoin,
		register.OpSign,
		signTx,
	)
}

const s3KeyPrxfix = "filecoin/"

// createAccount create new account address
func createAccount(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := ct.NewAccountRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	// set current net type main or test
	address.CurrentNetwork = fil.FILNetMap[info.ENV]

	ki, _addr, err := local.WalletNew(types.KTSecp256k1)
	if err != nil {
		return nil, err
	}

	addr := _addr.String()
	_out := ct.NewAccountResponse{
		Address: addr,
	}

	out, err = json.Marshal(_out)
	if err != nil {
		return nil, err
	}

	err = oss.PutObject(ctx, s3KeyPrxfix+addr, ki.PrivateKey, true)
	return out, err
}

// signTx sign a raw transaction
func signTx(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := fil.SignRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	raw := info.Info

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	// set current net type main or test
	address.CurrentNetwork = fil.FILNetMap[info.ENV]

	to, err := address.NewFromString(raw.To)
	if err != nil {
		return nil, err
	}
	from, err := address.NewFromString(raw.From)
	if err != nil {
		return nil, err
	}

	// big.NewFloat(raw.Value).String()/Text()
	val, err := ftypes.ParseFIL(decimal.NewFromFloat(raw.Value).String())
	if err != nil {
		return nil, err
	}

	pk, err := oss.GetObject(ctx, s3KeyPrxfix+raw.From, true)
	if err != nil {
		return
	}

	msg := types.Message{
		Version:    raw.Version,
		To:         to,
		From:       from,
		Nonce:      raw.Nonce,
		Value:      abi.TokenAmount(val),
		GasLimit:   raw.GasLimit,
		GasFeeCap:  abi.NewTokenAmount(raw.GasFeeCap),
		GasPremium: abi.NewTokenAmount(raw.GasPremium),
		Method:     raw.Method,
		Params:     raw.Params,
	}

	s, err := local.WalletSignMessage(types.KTSecp256k1, pk, &msg)
	if err != nil {
		return nil, err
	}

	err = local.WalletVerifyMessage(s)
	if err != nil {
		return nil, err
	}

	signType, err := s.Signature.Type.Name()
	if err != nil {
		return nil, err
	}

	_out := fil.BroadcastRequest{
		ENV: info.ENV,
		Raw: info.Info,
		Signature: fil.Signature{
			SignType: signType,
			Data:     s.Signature.Data,
		},
	}

	return json.Marshal(_out)
}
