package history

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	historycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/authing/history"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/history"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateHistory(ctx context.Context) (*npool.History, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	if h.UserID != nil {
		handler, err := usermw.NewHandler(
			ctx,
			usermw.WithConds(&usermwpb.Conds{
				AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
				EntID: &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
			}),
		)
		if err != nil {
			return nil, err
		}

		exist, err := handler.ExistUserConds(ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, fmt.Errorf("user not exists")
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := historycrud.CreateSet(
			cli.AuthHistory.Create(),
			&historycrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
				Allowed:  h.Allowed,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetHistory(ctx)
}
