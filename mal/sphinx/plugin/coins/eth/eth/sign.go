package eth

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	ct "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/types"

	"github.com/NpoolPlatform/kunman/framework/oss"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpWalletNew,
		CreateEthAccount,
	)
	register.RegisteTokenHandler(
		coins.Ethereum,
		register.OpSign,
		Msg,
	)
}

func Msg(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	s3KeyPrxfix := coins.GetS3KeyPrxfix(tokenInfo)
	return Message(ctx, s3KeyPrxfix, in)
}

func CreateEthAccount(ctx context.Context, in []byte, tokenInfo *coins.TokenInfo) (out []byte, err error) {
	s3KeyPrxfix := coins.GetS3KeyPrxfix(tokenInfo)
	return CreateAccount(ctx, s3KeyPrxfix, in)
}

func Message(ctx context.Context, s3Store string, in []byte) ([]byte, error) {
	preSignData := &eth.PreSignData{}
	if err := json.Unmarshal(in, preSignData); err != nil {
		return nil, err
	}

	pk, err := oss.GetObject(ctx, s3Store+preSignData.From, true)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(string(pk))
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(preSignData.Tx, types.NewEIP155Signer(preSignData.ChainID), privateKey)
	if err != nil {
		return nil, err
	}

	signedTxBuf := bytes.Buffer{}
	err = signedTx.EncodeRLP(&signedTxBuf)
	if err != nil {
		return nil, err
	}

	signedData := eth.SignedData{
		SignedTx: signedTxBuf.Bytes(),
	}

	return json.Marshal(signedData)
}

func CreateAccount(ctx context.Context, s3Store string, in []byte) (out []byte, err error) {
	info := ct.NewAccountRequest{}
	if err := json.Unmarshal(in, &info); err != nil {
		return nil, err
	}

	if !coins.CheckSupportNet(info.ENV) {
		return nil, env.ErrEVNCoinNetValue
	}

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyBytesHex := make([]byte, len(privateKeyBytes)*2)
	hex.Encode(privateKeyBytesHex, privateKeyBytes)

	// privateKey.PublicKey
	publicKeyECDSA := privateKey.PublicKey

	// crypto.PubkeyToAddress
	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// hash := sha3.NewKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// hexutil.Encode(hash.Sum(nil)[12:])
	address := crypto.PubkeyToAddress(publicKeyECDSA).Hex() // Hex String
	err = oss.PutObject(ctx, s3Store+address, privateKeyBytesHex, true)
	if err != nil {
		return nil, err
	}

	_out := &ct.NewAccountResponse{Address: address}
	out, err = json.Marshal(_out)
	if err != nil {
		return nil, err
	}

	return out, err
}
