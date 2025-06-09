package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	testinit "github.com/NpoolPlatform/kunman/middleware/miningpool/testinit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

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

var ret = &npool.Coin{
	EntID:               uuid.NewString(),
	CoinType:            basetypes.CoinType_CoinTypeBitCoin,
	PoolID:              uuid.NewString(),
	CoinTypeID:          uuid.NewString(),
	FixedRevenueAble:    false,
	LeastTransferAmount: decimal.NewFromFloat(5.4).String(),
	Remark:              "asdfaf",
}

var req = &npool.CoinReq{
	EntID:               &ret.EntID,
	CoinTypeID:          &ret.CoinTypeID,
	CoinType:            &ret.CoinType,
	PoolID:              &ret.PoolID,
	FixedRevenueAble:    &ret.FixedRevenueAble,
	LeastTransferAmount: &ret.LeastTransferAmount,
	Remark:              &ret.Remark,
}

func create(t *testing.T) {
	getH, err := NewHandler(context.Background(), WithConds(nil), WithOffset(0), WithLimit(2))
	assert.Nil(t, err)

	infos, _, err := getH.GetCoins(context.Background())
	assert.Nil(t, err)

	if len(infos) == 0 {
		return
	}

	ret.PoolID = infos[0].PoolID
	req.PoolID = &infos[0].PoolID
	ret.CoinType = infos[0].CoinType
	req.CoinType = &infos[0].CoinType

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithCoinType(req.CoinType, true),
		WithPoolID(req.PoolID, true),
		WithCoinTypeID(req.CoinTypeID, true),
		WithFixedRevenueAble(req.FixedRevenueAble, true),
		WithLeastTransferAmount(req.LeastTransferAmount, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoin(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithCoinType(req.CoinType, true),
		WithPoolID(req.PoolID, true),
		WithCoinTypeID(req.CoinTypeID, true),
		WithFixedRevenueAble(req.FixedRevenueAble, true),
		WithLeastTransferAmount(req.LeastTransferAmount, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoin(context.Background())
	assert.NotNil(t, err)
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
}
