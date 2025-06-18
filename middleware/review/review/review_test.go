package review

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	reviewmwpb "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"
	testinit "github.com/NpoolPlatform/kunman/middleware/review/testinit"
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
	ret = reviewmwpb.Review{
		EntID:         uuid.NewString(),
		AppID:         uuid.NewString(),
		Domain:        uuid.NewString(),
		ObjectType:    reviewtypes.ReviewObjectType_ObjectKyc,
		ObjectTypeStr: reviewtypes.ReviewObjectType_ObjectKyc.String(),
		ObjectID:      uuid.NewString(),
		ReviewerID:    uuid.Nil.String(),
		State:         reviewtypes.ReviewState_Wait,
		StateStr:      reviewtypes.ReviewState_Wait.String(),
		Trigger:       reviewtypes.ReviewTriggerType_DefaultTriggerType,
		TriggerStr:    reviewtypes.ReviewTriggerType_DefaultTriggerType.String(),
		Message:       "",
	}
)

func createReview(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithAppID(&ret.AppID, true),
		WithDomain(&ret.Domain, true),
		WithObjectID(&ret.ObjectID, true),
		WithObjectType(&ret.ObjectType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateReview(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateReview(t *testing.T) {
	ret.State = reviewtypes.ReviewState_Rejected
	ret.StateStr = reviewtypes.ReviewState_Rejected.String()
	ret.Message = uuid.NewString()
	ret.ReviewerID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithState(&ret.State, false),
		WithMessage(&ret.Message, false),
		WithReviewerID(&ret.ReviewerID, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateReview(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getReview(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetReview(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

var (
	conds = &reviewmwpb.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		ObjectID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.ObjectID},
		Domain:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Domain},
		Trigger:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewTriggerType_DefaultTriggerType)},
		ObjectType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewObjectType_ObjectKyc)},
		State:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewState_Rejected)},
	}
)

func getReviews(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetReviews(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, &ret, infos[0])
	}
}

func existReviewConds(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistReviewConds(context.Background())
	if assert.Nil(t, err) {
		assert.True(t, exist)
	}
}

func deleteReview(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteReview(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = handler.GetReview(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestReview(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createReviw", createReview)
	t.Run("updateReview", updateReview)
	t.Run("getReview", getReview)
	t.Run("getReviews", getReviews)
	t.Run("existReviewConds", existReviewConds)
	t.Run("deleteReview", deleteReview)
}
