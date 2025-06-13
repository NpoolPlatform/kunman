package ledger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	statement1 "github.com/NpoolPlatform/kunman/middleware/ledger/simulate/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/testinit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger/statement"
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

var (
	appID      = uuid.NewString()
	userID     = uuid.NewString()
	coinTypeID = uuid.NewString()

	miningBenefit = statementmwpb.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "100",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_MiningBenefit,
		IOSubTypeStr: basetypes.IOSubType_MiningBenefit.String(),
		IOExtra:      fmt.Sprintf(`{"GoodID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	ledgerResult = ledgermwpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "100",
		Outcoming:  "0",
	}
)

func setup(t *testing.T) func(*testing.T) {
	reqs1 := []*statementmwpb.StatementReq{
		{
			EntID:      &miningBenefit.EntID,
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &miningBenefit.Amount,
			IOType:     &miningBenefit.IOType,
			IOSubType:  &miningBenefit.IOSubType,
			IOExtra:    &miningBenefit.IOExtra,
		},
	}

	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithReqs(reqs1, true),
	)
	assert.Nil(t, err)

	deposits, err := handler.CreateStatements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(deposits))
		miningBenefit.CreatedAt = deposits[0].CreatedAt
		miningBenefit.UpdatedAt = deposits[0].UpdatedAt
		miningBenefit.ID = deposits[0].ID
		assert.Equal(t, &miningBenefit, deposits[0])
	}

	st1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&miningBenefit.EntID, true),
	)
	assert.Nil(t, err)

	return func(t *testing.T) {
		_, _ = st1.DeleteStatement(context.Background())
	}
}

func getLedgerOnly(t *testing.T) {
	conds := ledgermwpb.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(&conds),
	)
	assert.Nil(t, err)

	info, err := handler.GetLedgerOnly(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult.ID = info.ID
		ledgerResult.EntID = info.EntID
		ledgerResult.CreatedAt = info.CreatedAt
		ledgerResult.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult, info)
	}
}
func TestLedger(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getLedgerOnly", getLedgerOnly)
}
