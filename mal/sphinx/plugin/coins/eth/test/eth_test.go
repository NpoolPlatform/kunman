package test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/getter"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	npool "github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/test-go/testify/assert"
)

var (
	amount1        = "5.123"
	amount2        = 1.23
	amount3        = 7.23
	EthOnchainTime = coins.SyncTime[plugin.CoinType_CoinTypetethereum]
)

// run in dev-box or docker environment
func TestEthAndTokens(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	ctx := context.Background()
	coinInfo, err := env.GetCoinInfo()
	assert.Nil(t, err)
	assert.Equal(t, coinInfo.NetworkType, coins.CoinNetTest)

	coinType := coins.CoinStr2CoinType(coinInfo.NetworkType, coinInfo.CoinType)
	tokens := getter.GetTokenInfos(coinType)
	assert.NotEqual(t, 0, len(tokens))

	bcServer, ok := env.LookupEnv(env.ENVBUIILDCHIANSERVER)
	assert.Equal(t, ok, true)

	// create address and transfer gas
	// one tokentype has 2 addresses,named address1 and address2 by the following commants
	addrMap := make(map[*coins.TokenInfo][2]string, len(tokens))
	for _, token := range tokens {
		info1, err := CreateAddress(ctx, token.Name)
		assert.Nil(t, err)

		info2, err := CreateAddress(ctx, token.Name)
		assert.Nil(t, err)

		addrMap[token] = [2]string{info1.Address, info2.Address}

		// skip when its tokentype is  Ethereum，next step will transfer to
		if token.OfficialContract == string(coins.Ethereum) {
			continue
		}
	}

	// wait msg on chain
	time.Sleep(EthOnchainTime)
	for k, v := range addrMap {
		info, err := GetBalance(ctx, &npool.GetBalanceRequest{
			Name:    "tethereum",
			Address: v[0],
		})
		assert.Nil(t, err)
		assert.Equal(t, info.BalanceStr, amount1)

		info, err = GetBalance(ctx, &npool.GetBalanceRequest{
			Name:    k.Name,
			Address: v[0],
		})
		assert.Nil(t, err)
		assert.Equal(t, info.BalanceStr, amount1)
	}

	for token, addrs := range addrMap {
		// transfer token to address2 from address1,its balance is enough
		err = CreateTransaction(ctx, &npool.CreateTransactionRequest{
			Name:          token.Name,
			TransactionID: addrs[0], // make address1 as transactionid
			Amount:        amount2,
			From:          addrs[0],
			To:            addrs[1],
		})
		assert.Nil(t, err)
	}

	time.Sleep(EthOnchainTime * 3)
	for token, addrs := range addrMap {
		// transfer token to address2 from address1,its balance is not enough
		err = CreateTransaction(ctx, &npool.CreateTransactionRequest{
			Name:          token.Name,
			TransactionID: addrs[1], // make address2 as transactionid
			Amount:        amount3,
			From:          addrs[0],
			To:            addrs[1],
		})
		assert.Nil(t, err)
	}

	// wait msg on chain
	time.Sleep(EthOnchainTime * 3)
	for token, addrs := range addrMap {
		info, err := GetBalance(ctx, &npool.GetBalanceRequest{
			Name:    token.Name,
			Address: addrs[1],
		})
		assert.Nil(t, err)
		assert.Equal(t, info.Balance, amount2)

		tInfo, err := GetTransaction(ctx, addrs[0])
		assert.Nil(t, err)
		assert.Equal(t, tInfo.Amount, amount2)
		assert.Equal(t, tInfo.From, addrs[0])
		assert.Equal(t, tInfo.TransactionState, npool.TransactionState_TransactionStateDone)

		tInfo, err = GetTransaction(ctx, addrs[1])
		assert.Nil(t, err)
		assert.Equal(t, tInfo.Amount, amount3)
		assert.Equal(t, tInfo.From, addrs[0])
		assert.Equal(t, tInfo.TransactionState, npool.TransactionState_TransactionStateFail)
	}
}

// rewrite sphinx-proxy client do func
func do(ctx context.Context, fn func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	host, ok := env.LookupEnv(env.ENVPROXY)
	if !ok {
		return nil, env.ErrENVProxyInvalid
	}
	conn, err := grpc.DialContext(_ctx, host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connState := conn.GetState()
	if connState != connectivity.Idle && connState != connectivity.Ready {
		return nil, err
	}

	cli := npool.NewSphinxProxyClient(conn)

	return fn(_ctx, cli)
}

func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (*npool.BalanceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.GetBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail get balance: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get balance: %v", err)
	}
	return info.(*npool.BalanceInfo), nil
}

func CreateTransaction(ctx context.Context, in *npool.CreateTransactionRequest) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		_, err := cli.CreateTransaction(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail get balances: %v", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("fail get balances: %v", err)
	}
	return nil
}

func GetTransaction(ctx context.Context, id string) (*npool.TransactionInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.GetTransaction(ctx, &npool.GetTransactionRequest{
			TransactionID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transaction: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transaction: %v", err)
	}
	return info.(*npool.TransactionInfo), nil
}

func CreateAddress(ctx context.Context, coinName string) (*npool.WalletInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.CreateWallet(ctx, &npool.CreateWalletRequest{
			Name: coinName,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create wallet: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create wallet: %v", err)
	}
	return info.(*npool.WalletInfo), nil
}
