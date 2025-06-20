package ledger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	statement1 "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/testinit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
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

	deposit = statementmwpb.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "100",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_Deposit,
		IOSubTypeStr: basetypes.IOSubType_Deposit.String(),
		IOExtra:      fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	payment = statementmwpb.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "10",
		IOType:       basetypes.IOType_Outcoming,
		IOTypeStr:    basetypes.IOType_Outcoming.String(),
		IOSubType:    basetypes.IOSubType_Payment,
		IOSubTypeStr: basetypes.IOSubType_Payment.String(),
		IOExtra:      fmt.Sprintf(`{"PaymentID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}

	ledgerResult = ledgermwpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "100",
		Outcoming:  "10",
		Locked:     "0",
		Spendable:  "90",
	}
)

func setup(t *testing.T) func(*testing.T) {
	reqs1 := []*statementmwpb.StatementReq{
		{
			EntID:      &deposit.EntID,
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &deposit.Amount,
			IOType:     &deposit.IOType,
			IOSubType:  &deposit.IOSubType,
			IOExtra:    &deposit.IOExtra,
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
		deposit.CreatedAt = deposits[0].CreatedAt
		deposit.UpdatedAt = deposits[0].UpdatedAt
		deposit.ID = deposits[0].ID
		assert.Equal(t, &deposit, deposits[0])
	}

	reqs2 := []*statementmwpb.StatementReq{
		{
			EntID:      &payment.EntID,
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &payment.Amount,
			IOType:     &payment.IOType,
			IOSubType:  &payment.IOSubType,
			IOExtra:    &payment.IOExtra,
		},
	}

	handler2, err := statement1.NewHandler(
		context.Background(),
		statement1.WithReqs(reqs2, true),
	)
	assert.Nil(t, err)

	payments, err := handler2.CreateStatements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(payments))
		payment.CreatedAt = payments[0].CreatedAt
		payment.UpdatedAt = payments[0].UpdatedAt
		payment.ID = payments[0].ID
		assert.Equal(t, &payment, payments[0])
	}

	st1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&deposit.EntID, true),
	)
	assert.Nil(t, err)
	st2, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&payment.EntID, true),
	)
	assert.Nil(t, err)

	return func(t *testing.T) {
		_, _ = st1.DeleteStatement(context.Background())
		_, _ = st2.DeleteStatement(context.Background())
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
