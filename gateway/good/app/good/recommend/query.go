package recommend

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/recommend"
	recommendmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	recommends []*recommendmwpb.Recommend
	infos      []*npool.Recommend
	apps       map[string]*appmwpb.App
	users      map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, recommend := range h.recommends {
			appIDs = append(appIDs, recommend.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = goodgwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, recommend := range h.recommends {
			userIDs = append(userIDs, recommend.RecommenderID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, recommend := range h.recommends {
		info := &npool.Recommend{
			ID:             recommend.ID,
			EntID:          recommend.EntID,
			AppID:          recommend.AppID,
			RecommenderID:  recommend.RecommenderID,
			AppGoodID:      recommend.AppGoodID,
			GoodName:       recommend.GoodName,
			RecommendIndex: recommend.RecommendIndex,
			Message:        recommend.Message,
			Hide:           recommend.Hide,
			HideReason:     recommend.HideReason,
			CreatedAt:      recommend.CreatedAt,
			UpdatedAt:      recommend.UpdatedAt,
		}

		app, ok := h.apps[recommend.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[recommend.RecommenderID]
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

func (h *Handler) GetRecommend(ctx context.Context) (*npool.Recommend, error) {
	recommendHandler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	recommend, err := recommendHandler.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}
	if recommend == nil {
		return nil, fmt.Errorf("invalid recommend")
	}

	handler := &queryHandler{
		Handler:    h,
		recommends: []*recommendmwpb.Recommend{recommend},
		apps:       map[string]*appmwpb.App{},
		users:      map[string]*usermwpb.User{},
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

func (h *Handler) GetRecommends(ctx context.Context) ([]*npool.Recommend, uint32, error) {
	conds := &recommendmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.RecommenderID != nil {
		conds.RecommenderID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID}
	}

	recommendHandler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithConds(conds),
		recommendmw.WithOffset(h.Offset),
		recommendmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	recommends, total, err := recommendHandler.GetRecommends(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(recommends) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:    h,
		recommends: recommends,
		apps:       map[string]*appmwpb.App{},
		users:      map[string]*usermwpb.User{},
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
