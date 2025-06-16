package depc

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

type DepincAccount struct {
	*btcutil.WIF
	netParams        *chaincfg.Params
	CompressedPubKey []byte
	PayAddress       *btcutil.AddressScriptHash
	PayAddressStr    string
	RedeemScript     []byte
	ScriptPubKey     []byte
}

func New(net *chaincfg.Params) (*DepincAccount, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	wif, err := btcutil.NewWIF(secret, net, true)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return initDepincAccount(wif, net)
}

func NewFromWIFString(wifStr string) (*DepincAccount, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	// parse network ID from privateKey
	decoded := base58.Decode(wifStr)
	netID := decoded[0]

	var net *chaincfg.Params
	switch netID {
	case chaincfg.MainNetParams.PrivateKeyID:
		net = &chaincfg.MainNetParams
	case chaincfg.RegressionNetParams.PrivateKeyID:
		net = &chaincfg.RegressionNetParams
	case chaincfg.TestNet3Params.PrivateKeyID:
		net = &chaincfg.TestNet3Params
	case chaincfg.SigNetParams.PrivateKeyID:
		net = &chaincfg.SigNetParams
	case chaincfg.SimNetParams.PrivateKeyID:
		net = &chaincfg.SimNetParams
	default:
		return nil, wlog.Errorf("not support network format")
	}
	return initDepincAccount(wif, net)
}

func initDepincAccount(wif *btcutil.WIF, net *chaincfg.Params) (*DepincAccount, error) {
	da := &DepincAccount{WIF: wif, netParams: net}
	// compressed pub key
	compressedPubKey := da.PrivKey.PubKey().SerializeCompressed()

	// pub key hash
	pkHash := btcutil.Hash160(compressedPubKey)

	// redeemScript
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(pkHash).Script()
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	// redeemScript hash
	pkhsHash, err := btcutil.NewAddressScriptHash(redeemScript, net)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	// script pub key,locking script
	scriptPubKey, err := txscript.PayToAddrScript(pkhsHash)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	da.CompressedPubKey = compressedPubKey
	da.PayAddress = pkhsHash
	da.PayAddressStr = pkhsHash.EncodeAddress()
	da.RedeemScript = redeemScript
	da.ScriptPubKey = scriptPubKey
	return da, nil
}
