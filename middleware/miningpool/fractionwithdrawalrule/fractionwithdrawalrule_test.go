package fractionwithdrawalrule

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawalrule"
	testinit "github.com/NpoolPlatform/kunman/middleware/miningpool/testinit"
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

var ret = &npool.FractionWithdrawalRule{
	EntID:                 uuid.NewString(),
	PoolCoinTypeID:        uuid.NewString(),
	WithdrawInterval:      1000,
	PayoutThreshold:       "5.0",
	LeastWithdrawalAmount: "2.0",
	WithdrawFee:           "0.03",
}

var req = &npool.FractionWithdrawalRuleReq{
	EntID:                 &ret.EntID,
	PoolCoinTypeID:        &ret.PoolCoinTypeID,
	WithdrawInterval:      &ret.WithdrawInterval,
	PayoutThreshold:       &ret.PayoutThreshold,
	LeastWithdrawalAmount: &ret.LeastWithdrawalAmount,
	WithdrawFee:           &ret.WithdrawFee,
}

func create(t *testing.T) {
	getH, err := NewHandler(context.Background(), WithConds(nil), WithOffset(0), WithLimit(2))
	assert.Nil(t, err)

	infos, _, err := getH.GetFractionWithdrawalRules(context.Background())
	assert.Nil(t, err)
	if len(infos) == 0 {
		return
	}

	ret.PoolCoinTypeID = infos[0].PoolCoinTypeID
	req.PoolCoinTypeID = &infos[0].PoolCoinTypeID
	ret.CoinType = infos[0].CoinType
	ret.EntID = infos[0].EntID
	req.EntID = &infos[0].EntID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithPoolCoinTypeID(req.PoolCoinTypeID, true),
		WithWithdrawInterval(req.WithdrawInterval, true),
		WithPayoutThreshold(req.PayoutThreshold, true),
		WithLeastWithdrawalAmount(req.LeastWithdrawalAmount, true),
		WithWithdrawFee(req.WithdrawFee, true),
	)
	assert.Nil(t, err)

	err = handler.CreateFractionWithdrawalRule(context.Background())
	assert.NotNil(t, err)
}

func TestFractionWithdrawalRule(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
}
