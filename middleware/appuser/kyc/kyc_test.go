package kyc

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/kyc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	user "github.com/NpoolPlatform/kunman/middleware/appuser/user"
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
	ret = npool.Kyc{
		EntID:           uuid.NewString(),
		AppID:           uuid.NewString(),
		UserID:          uuid.NewString(),
		DocumentType:    basetypes.KycDocumentType_IDCard,
		DocumentTypeStr: basetypes.KycDocumentType_IDCard.String(),
		IDNumber:        uuid.NewString(),
		FrontImg:        uuid.NewString(),
		BackImg:         uuid.NewString(),
		SelfieImg:       uuid.NewString(),
		EntityType:      basetypes.KycEntityType_Individual,
		EntityTypeStr:   basetypes.KycEntityType_Individual.String(),
		ReviewID:        uuid.NewString(),
		State:           basetypes.KycState_Reviewing,
		StateStr:        basetypes.KycState_Reviewing.String(),
	}
)

func setupKyc(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&ret.UserID, true),
		app.WithName(&ret.AppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ah, err = app.NewHandler(
		context.Background(),
		app.WithID(&app1.ID, true),
	)
	assert.Nil(t, err)

	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+1000000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+5000000) //nolint
	passwordHash := uuid.NewString()

	ret.AppName = ret.AppID

	uh, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.UserID, true),
		user.WithAppID(&ret.AppID, true),
		user.WithPhoneNO(&ret.PhoneNO, true),
		user.WithEmailAddress(&ret.EmailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh)
	user1, err := uh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user1)

	uh, err = user.NewHandler(
		context.Background(),
		user.WithID(&user1.ID, true),
	)
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = uh.DeleteUser(context.Background())
	}
}

func createKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithDocumentType(&ret.DocumentType, true),
		WithIDNumber(&ret.IDNumber, true),
		WithFrontImg(&ret.FrontImg, true),
		WithBackImg(&ret.BackImg, true),
		WithSelfieImg(&ret.SelfieImg, true),
		WithEntityType(&ret.EntityType, true),
		WithReviewID(&ret.ReviewID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateKyc(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateKyc(t *testing.T) {
	ret.State = basetypes.KycState_Approved
	ret.StateStr = basetypes.KycState_Approved.String()
	ret.FrontImg = uuid.NewString()
	ret.BackImg = uuid.NewString()
	ret.SelfieImg = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithDocumentType(&ret.DocumentType, false),
		WithIDNumber(&ret.IDNumber, false),
		WithFrontImg(&ret.FrontImg, false),
		WithBackImg(&ret.BackImg, false),
		WithSelfieImg(&ret.SelfieImg, false),
		WithEntityType(&ret.EntityType, false),
		WithReviewID(&ret.ReviewID, false),
		WithState(&ret.State, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateKyc(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetKyc(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getKycs(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetKycs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteKyc(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetKyc(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestKyc(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupKyc(t)
	defer teardown(t)

	t.Run("createKyc", createKyc)
	t.Run("updateKyc", updateKyc)
	t.Run("getKyc", getKyc)
	t.Run("getKycs", getKycs)
	t.Run("deleteKyc", deleteKyc)
}
