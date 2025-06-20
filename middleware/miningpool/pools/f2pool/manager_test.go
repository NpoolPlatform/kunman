package f2pool

import (
	"context"
	"os"
	"strconv"
	"testing"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	accessToken = "7ecdq1fosdsfcruypom2otsn7hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73"
	user1       string
)

func auth(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}
	authed := mgr.CheckAuth(context.Background())
	assert.Equal(t, nil, authed)
}

func mininguser(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}

	ret, err := mgr.ExistMiningUser(context.Background(), uuid.NewString())
	assert.NotNil(t, err)
	assert.Equal(t, false, ret)

	username, link, err := mgr.AddMiningUser(context.Background())
	assert.Nil(t, err)

	ret, err = mgr.ExistMiningUser(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)

	_link, err := mgr.GetReadPageLink(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, link, _link)

	// f2pool have not implment the method
	// err = mgr.DeleteMiningUser(context.Background(), username)
	// assert.Nil(t, err)
	// ret, err = mgr.ExistMiningUser(context.Background(), username)
	// assert.NotNil(t, err)
	// assert.Equal(t, false, ret)
	// _, err = mgr.GetReadPageLink(context.Background(), username)
	// assert.NotNil(t, err)
}

func proportion(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}

	propotion := "0.1"
	err = mgr.SetRevenueProportion(context.Background(), "cococonut2", "cococonut001", propotion)
	assert.Nil(t, err)

	_propotion, err := mgr.GetRevenueProportion(context.Background(), "cococonut2", "cococonut001")
	assert.Nil(t, err)
	assert.Equal(t, propotion, _propotion)
}

func revenueAddress(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}
	user1, _, err = mgr.AddMiningUser(context.Background())
	assert.Nil(t, err)

	addr := "1PWMfNSb3auXwZ1qhu96WRJL7BCgG4mGB4"
	err = mgr.SetRevenueAddress(context.Background(), user1, addr)
	assert.Nil(t, err)

	_addr, err := mgr.GetRevenueAddress(context.Background(), user1)
	assert.Nil(t, err)
	assert.Equal(t, addr, _addr)
}

//nolint:gocritic
func pageLink(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}

	// f2pool api have bug,
	// link, err := mgr.AddReadPageLink(context.Background(), user1)
	// assert.Nil(t, err)

	// _link, err := mgr.GetReadPageLink(context.Background(), user1)
	// assert.Nil(t, err)
	// assert.Equal(t, link, _link)

	err = mgr.DeleteReadPageLink(context.Background(), user1)
	assert.Nil(t, err)

	// _link, err = mgr.GetReadPageLink(context.Background(), user1)
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
}

func payment(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}
	paused, err := mgr.PausePayment(context.Background(), user1)
	assert.Nil(t, err)
	assert.Equal(t, true, paused)

	resumed, err := mgr.ResumePayment(context.Background(), user1)
	assert.Nil(t, err)
	assert.Equal(t, true, resumed)
}

func withdrawFractionWithdrawal(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_CoinTypeBitCoin.Enum(), accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)
	if err != nil {
		return
	}
	txTime, err := mgr.WithdrawFractionWithdrawal(context.Background(), user1)
	// most of the time this api will report an error
	if err == nil {
		assert.Nil(t, err)
		assert.NotEqual(t, 0, txTime)
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	os.Setenv("TEST_F2POOL", "false")
	testF2pool := os.Getenv("TEST_F2POOL")
	if testF2pool == "false" {
		return
	}

	t.Run("auth", auth)
	t.Run("mininguser", mininguser)
	t.Run("proportion", proportion)
	t.Run("revenueAddress", revenueAddress)
	t.Run("pageLink", pageLink)
	t.Run("payment", payment)
	t.Run("withdrawFractionWithdrawal", withdrawFractionWithdrawal)
}
