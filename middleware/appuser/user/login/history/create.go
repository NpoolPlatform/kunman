package history

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/login/history"
	historycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/login/history"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	user1 "github.com/NpoolPlatform/kunman/middleware/appuser/user"

	"github.com/google/uuid"
)

func (h *Handler) CreateHistory(ctx context.Context) (*npool.History, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	userID := h.UserID.String()
	appID := h.AppID.String()

	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&userID, true),
		user1.WithAppID(&appID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistUser(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := historycrud.CreateSet(
			cli.LoginHistory.Create(),
			&historycrud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				UserID:    h.UserID,
				ClientIP:  h.ClientIP,
				UserAgent: h.UserAgent,
				Location:  h.Location,
				LoginType: h.LoginType,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetHistory(ctx)
}
