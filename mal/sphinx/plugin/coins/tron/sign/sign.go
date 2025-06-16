package tron

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	addr "github.com/Geapefurit/gotron-sdk/pkg/address"
	"github.com/NpoolPlatform/kunman/framework/oss"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"

	"google.golang.org/protobuf/proto"
)

func init() {
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpWalletNew,
		CreateTrxAccount,
	)
	register.RegisteTokenHandler(
		coins.Tron,
		register.OpSign,
		SignTrxMSG,
	)
}

const s3KeyPrxfix = "tron/"

func SignTrxMSG(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	return SignTronMSG(ctx, s3KeyPrxfix, in)
}

func CreateTrxAccount(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	return CreateTronAccount(ctx, s3KeyPrxfix, in)
}

func SignTronMSG(ctx context.Context, s3Strore string, in []byte) (out []byte, err error) {
	signMsgTx := &tron.SignMsgTx{}
	err = json.Unmarshal(in, signMsgTx)
	if err != nil {
		return in, err
	}
	if signMsgTx.TxExtension == nil {
		return in, errors.New(tron.BuildTransactionFailed)
	}

	pk, err := oss.GetObject(ctx, s3Strore+signMsgTx.Base.From, true)
	if err != nil {
		return in, err
	}

	privateBytes, err := hex.DecodeString(string(pk))
	if err != nil {
		return in, err
	}
	transaction := signMsgTx.TxExtension.Transaction
	priv := crypto.ToECDSAUnsafe(privateBytes)
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, fmt.Errorf("proto marshal tx raw data error: %v", err)
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, priv)
	if err != nil {
		return nil, fmt.Errorf("sign error: %v", err)
	}

	transaction.Signature = append(transaction.Signature, signature)
	signedMsg := &tron.BroadcastRequest{TxExtension: signMsgTx.TxExtension}
	return json.Marshal(signedMsg)
}

func CreateTronAccount(ctx context.Context, s3Strore string, in []byte) (out []byte, err error) {
	info := ct.NewAccountRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	if len(priv.D.Bytes()) != 32 {
		for {
			priv, err := btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(priv.D.Bytes()) == 32 {
				break
			}
		}
	}

	a := addr.PubkeyToAddress(priv.ToECDSA().PublicKey)
	pubkey := a.String()
	prikey := hex.EncodeToString(priv.D.Bytes())

	err = oss.PutObject(ctx, s3Strore+pubkey, []byte(prikey), true)
	if err != nil {
		return nil, err
	}

	naResp := &ct.NewAccountResponse{Address: pubkey}
	return json.Marshal(naResp)
}
