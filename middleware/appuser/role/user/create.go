package user

import (
	"context"
	"fmt"

	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role/user"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	role1 "github.com/NpoolPlatform/kunman/middleware/appuser/role"
	user1 "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

//nolint:gocyclo
func (h *Handler) CreateUser(ctx context.Context) (*npool.User, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	if h.RoleID == nil || h.UserID == nil {
		return nil, fmt.Errorf("invalid roleid or userid")
	}

	userID := h.UserID.String()
	appID := h.AppID.String()
	roleID := h.RoleID.String()

	h1, err := role1.NewHandler(
		ctx,
		role1.WithAppID(&appID, true),
		role1.WithEntID(&roleID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err := h1.ExistRole(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid role")
	}

	h2, err := user1.NewHandler(
		ctx,
		user1.WithAppID(&appID, true),
		user1.WithEntID(&userID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err = h2.ExistUser(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateRoleUser, *h.AppID, *h.RoleID, *h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(
			cli.AppRoleUser.Query(),
			&usercrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				RoleID: &cruder.Cond{Op: cruder.EQ, Val: *h.RoleID},
				UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.EntID = &info.EntID
			return nil
		}

		if _, err := usercrud.CreateSet(
			cli.AppRoleUser.Create(),
			&usercrud.Req{
				EntID:  h.EntID,
				AppID:  h.AppID,
				RoleID: h.RoleID,
				UserID: h.UserID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
