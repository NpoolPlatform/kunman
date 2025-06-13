package ledger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	unsoldstatement "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger/unsold"
	goodstatement1 "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/statement"
	unsold1 "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/unsold"
	"github.com/NpoolPlatform/kunman/middleware/ledger/testinit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
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
	id                        = uint32(0)
	entID                     = uuid.NewString()
	goodID                    = uuid.NewString()
	coinTypeID                = uuid.NewString()
	totalAmount               = "400"
	unsoldAmount              = "100"
	techniqueServiceFeeAmount = "2"
	benefitDate               = uint32(time.Now().Unix())
)

func setup(t *testing.T) func(*testing.T) {
	reqs := []*goodstatementmwpb.GoodStatementReq{
		{
			EntID:                     &entID,
			GoodID:                    &goodID,
			CoinTypeID:                &coinTypeID,
			TotalAmount:               &totalAmount,
			UnsoldAmount:              &unsoldAmount,
			TechniqueServiceFeeAmount: &techniqueServiceFeeAmount,
			BenefitDate:               &benefitDate,
		},
	}
	handler, err := goodstatement1.NewHandler(
		context.Background(),
		goodstatement1.WithReqs(reqs, true),
	)
	assert.Nil(t, err)

	infos, err := handler.CreateGoodStatements(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
		assert.Equal(t, entID, infos[0].EntID)
		id = infos[0].ID
	}

	handler1, err := goodstatement1.NewHandler(context.Background(), goodstatement1.WithID(&id, true))
	assert.Nil(t, err)

	return func(t *testing.T) {
		_, _ = handler1.DeleteGoodStatement(context.Background())
	}
}

func getUnsoldStatements(t *testing.T) {
	handler, err := unsold1.NewHandler(context.Background())
	assert.Nil(t, err)
	handler.Conds = &unsoldstatement.Conds{
		StatementID: &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(entID)},
	}
	info, err := handler.GetUnsoldStatementOnly(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		assert.Equal(t, "100", info.Amount)
	}
}

func TestGoodLedger(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	teardown := setup(t)
	defer teardown(t)
	t.Run("getUnsoldStatements", getUnsoldStatements)
}
