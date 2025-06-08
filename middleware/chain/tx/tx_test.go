package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.Tx{
	CoinUnit:      "BTC",
	CoinLogo:      uuid.NewString(),
	CoinENV:       "test",
	FromAccountID: uuid.NewString(),
	ToAccountID:   uuid.NewString(),
	Amount:        "123.1",
	FeeAmount:     "2.01",
	State:         basetypes.TxState_TxStateCreated,
	StateStr:      basetypes.TxState_TxStateCreated.String(),
	Type:          basetypes.TxType_TxWithdraw,
	TypeStr:       basetypes.TxType_TxWithdraw.String(),
	Extra:         uuid.NewString(),
}

var req = &npool.TxReq{
	FromAccountID: &ret.FromAccountID,
	ToAccountID:   &ret.ToAccountID,
	Amount:        &ret.Amount,
	FeeAmount:     &ret.FeeAmount,
	State:         &ret.State,
	Type:          &ret.Type,
	Extra:         &ret.Extra,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithEntID(&ret.CoinTypeID, true),
		coin1.WithName(&ret.CoinName, true),
		coin1.WithUnit(&ret.CoinUnit, true),
		coin1.WithLogo(&ret.CoinLogo, true),
		coin1.WithENV(&ret.CoinENV, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithCoinTypeID(req.CoinTypeID, true),
		WithFromAccountID(req.FromAccountID, true),
		WithToAccountID(req.ToAccountID, true),
		WithAmount(req.Amount, true),
		WithFeeAmount(req.FeeAmount, true),
		WithChainTxID(req.ChainTxID, false),
		WithState(req.State, true),
		WithExtra(req.Extra, true),
		WithType(req.Type, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.State = basetypes.TxState_TxStateCreatedCheck
	ret.StateStr = ret.State.String()
	req.State = &ret.State

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithChainTxID(req.ChainTxID, false),
		WithState(req.State, false),
		WithExtra(req.Extra, false),
		WithType(req.Type, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.State = basetypes.TxState_TxStateWait
	ret.StateStr = ret.State.String()
	req.State = &ret.State

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithChainTxID(req.ChainTxID, false),
		WithState(req.State, false),
		WithExtra(req.Extra, false),
		WithType(req.Type, false),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateTx(context.Background())
	assert.Nil(t, err)
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
}
