package auth

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/auth"
	rolemwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	authcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/authing/auth"
	rolemw "github.com/NpoolPlatform/kunman/middleware/appuser/role"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *Handler) CreateAuth(ctx context.Context) (*npool.Auth, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	if h.UserID != nil {
		exist, err := h.ExistAuth(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("auth exist")
		}
	}
	if h.RoleID != nil {
		handler, err := rolemw.NewHandler(
			ctx,
			rolemw.WithConds(&rolemwpb.Conds{
				AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
				EntID: &basetypes.StringVal{Op: cruder.EQ, Value: h.RoleID.String()},
			}),
		)
		if err != nil {
			return nil, err
		}

		exist, err := handler.ExistRoleConds(ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, fmt.Errorf("role not exists")
		}

		h.Conds = &authcrud.Conds{
			AppID:    &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			RoleID:   &cruder.Cond{Op: cruder.EQ, Val: *h.RoleID},
			Resource: &cruder.Cond{Op: cruder.EQ, Val: *h.Resource},
			Method:   &cruder.Cond{Op: cruder.EQ, Val: *h.Method},
		}
		exist, err = h.ExistAuthConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("auth exist")
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := authcrud.CreateSet(
			cli.Auth.Create(),
			&authcrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				RoleID:   h.RoleID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAuth(ctx)
}

func (h *Handler) CreateAuths(ctx context.Context) ([]*npool.Auth, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			handler := &createHandler{
				Handler: h,
			}
			handler.EntID = req.EntID
			handler.AppID = req.AppID
			handler.Resource = req.Resource
			handler.Method = req.Method
			handler.UserID = req.UserID
			handler.RoleID = req.RoleID
			exist, err := handler.ExistAuth(ctx)
			if err != nil {
				return err
			}
			if exist {
				continue
			}
			id := uuid.New()
			if req.EntID == nil {
				req.EntID = &id
			}
			if _, err := authcrud.CreateSet(
				tx.Auth.Create(),
				req,
			).Save(_ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &authcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetAuths(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
