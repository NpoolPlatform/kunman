package sign

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"

	"github.com/NpoolPlatform/kunman/framework/oss"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/sol"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/log"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
)

func init() {
	register.RegisteTokenHandler(
		coins.Solana,
		register.OpWalletNew,
		createAccount,
	)
	register.RegisteTokenHandler(
		coins.Solana,
		register.OpSign,
		signTx,
	)
}

const s3KeyPrxfix = "solana/"

// createAccount ..
func createAccount(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := ct.NewAccountRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	// if account equal nil will panic
	account := solana.NewWallet()
	addr := account.PublicKey().String()
	_out := ct.NewAccountResponse{
		Address: addr,
	}

	out, err = json.Marshal(_out)
	if err != nil {
		return nil, err
	}

	err = oss.PutObject(ctx, s3KeyPrxfix+addr, account.PrivateKey, true)
	if err != nil {
		return nil, err
	}

	return out, err
}

// signTx ..
func signTx(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	info := sol.SignMsgTx{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	var (
		from   = info.BaseInfo.From
		to     = info.BaseInfo.To
		value  = info.BaseInfo.Value
		rbHash = info.RecentBlockHash
	)

	fPublicKey, err := solana.PublicKeyFromBase58(from)
	if err != nil {
		return nil, err
	}

	tPublicKey, err := solana.PublicKeyFromBase58(to)
	if err != nil {
		return nil, err
	}

	rhash, err := solana.HashFromBase58(rbHash)
	if err != nil {
		return nil, err
	}

	lamports, accuracy := sol.ToLarm(value)
	if accuracy != big.Exact {
		log.Warnf("transafer sol amount not accuracy: from %v-> to %v", value, lamports)
	}

	// build tx
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				lamports,
				fPublicKey,
				tPublicKey,
			).Build(),
		},
		rhash,
		solana.TransactionPayer(fPublicKey),
	)
	if err != nil {
		return nil, err
	}

	pk, err := oss.GetObject(ctx, s3KeyPrxfix+from, true)
	if err != nil {
		return nil, err
	}

	accountFrom := solana.PrivateKey(pk)
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	err = tx.VerifySignatures()
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	if err := tx.MarshalWithEncoder(bin.NewBinEncoder(&buf)); err != nil {
		return nil, err
	}

	_out := sol.BroadcastRequest{
		Signature: buf.Bytes(),
	}

	return json.Marshal(_out)
}
