package sign

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	proxypb "github.com/NpoolPlatform/kunman/message/sphinx/proxy"
)

var (
	ErrCoinSignTypeAlreadyRegister = errors.New("coin sign type already register")
	ErrOpSignTypeAlreadyRegister   = errors.New("op sign type already register")

	ErrCoinSignTypeNotRegister = errors.New("coin sign type not register")
	ErrOpSignTypeNotRegister   = errors.New("op sign type not register")

	coinSignHandles       = make(map[plugin.CoinType]map[proxypb.TransactionState]Handlef)
	coinWalletSignHandles = make(map[plugin.CoinType]map[proxypb.TransactionType]Handlef)
)

type Handlef func(ctx context.Context, payload []byte) ([]byte, error)

func Register(coinType plugin.CoinType, opType proxypb.TransactionState, handle Handlef) {
	coinPluginHandle, ok := coinSignHandles[coinType]
	if !ok {
		coinSignHandles[coinType] = make(map[proxypb.TransactionState]Handlef)
	}
	if _, ok := coinPluginHandle[opType]; ok {
		panic(fmt.Errorf("coin type: %v for transaction: %v already registered", coinType, opType))
	}
	coinSignHandles[coinType][opType] = handle
}

func GetCoinSign(coinType plugin.CoinType, opType proxypb.TransactionState) (Handlef, error) {
	// TODO: check nested map exist
	if _, ok := coinSignHandles[coinType]; !ok {
		return nil, ErrCoinSignTypeNotRegister
	}
	if _, ok := coinSignHandles[coinType][opType]; !ok {
		return nil, ErrOpSignTypeNotRegister
	}
	return coinSignHandles[coinType][opType], nil
}

func RegisterWallet(coinType plugin.CoinType, opType proxypb.TransactionType, handle Handlef) {
	coinWalletPluginHandle, ok := coinWalletSignHandles[coinType]
	if !ok {
		coinWalletSignHandles[coinType] = make(map[proxypb.TransactionType]Handlef)
	}
	if _, ok := coinWalletPluginHandle[opType]; ok {
		panic(fmt.Errorf("coin type: %v for transaction: %v already registered", coinType, opType))
	}
	coinWalletSignHandles[coinType][opType] = handle
}

func GetCoinWalletSign(coinType plugin.CoinType, opType proxypb.TransactionType) (Handlef, error) {
	// TODO: check nested map exist
	if _, ok := coinWalletSignHandles[coinType]; !ok {
		return nil, ErrCoinSignTypeNotRegister
	}
	if _, ok := coinWalletSignHandles[coinType][opType]; !ok {
		return nil, ErrOpSignTypeNotRegister
	}
	return coinWalletSignHandles[coinType][opType], nil
}
