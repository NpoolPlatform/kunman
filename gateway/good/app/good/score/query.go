package score

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
	scoremwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/score"
	scoremw "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	scores []*scoremwpb.Score
	infos  []*npool.Score
	apps   map[string]*appmwpb.App
	users  map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, score := range h.scores {
			appIDs = append(appIDs, score.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = goodgwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, score := range h.scores {
			userIDs = append(userIDs, score.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, score := range h.scores {
		info := &npool.Score{
			ID:        score.ID,
			EntID:     score.EntID,
			AppID:     score.AppID,
			UserID:    score.UserID,
			GoodID:    score.GoodID,
			AppGoodID: score.AppGoodID,
			GoodName:  score.GoodName,
			Score:     score.Score,
			CreatedAt: score.CreatedAt,
		}

		app, ok := h.apps[score.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[score.UserID]
		if ok {
			if user.Username != "" {
				info.Username = &user.Username
			}
			if user.EmailAddress != "" {
				info.EmailAddress = &user.EmailAddress
			}
			if user.PhoneNO != "" {
				info.PhoneNO = &user.PhoneNO
			}
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetScore(ctx context.Context) (*npool.Score, error) {
	scoreHandler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	score, err := scoreHandler.GetScore(ctx)
	if err != nil {
		return nil, err
	}
	if score == nil {
		return nil, fmt.Errorf("invalid score")
	}

	handler := &queryHandler{
		Handler: h,
		scores:  []*scoremwpb.Score{score},
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}
	return handler.infos[0], nil
}

func (h *Handler) GetScores(ctx context.Context) ([]*npool.Score, uint32, error) {
	if h.UserID != nil {
		if err := h.CheckUser(ctx); err != nil {
			return nil, 0, err
		}
	}

	conds := &scoremwpb.Conds{}
	conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}

	scoreHandler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithConds(conds),
		scoremw.WithOffset(h.Offset),
		scoremw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	scores, total, err := scoreHandler.GetScores(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(scores) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		scores:  scores,
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
