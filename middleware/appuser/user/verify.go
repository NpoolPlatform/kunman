//nolint:dupl
package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/appuser/encrypt"

	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entappusersecret "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appusersecret"

	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type verifyHandler struct {
	*Handler
	stm *ent.AppUserSelect
}

func (h *verifyHandler) queryAppUserByAccount(cli *ent.Client) error {
	if h.EmailAddress == nil && h.PhoneNO == nil {
		return fmt.Errorf("invalid account")
	}

	conds := &usercrud.Conds{
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
	}
	if h.EmailAddress != nil {
		conds.EmailAddress = &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress}
	}
	if h.PhoneNO != nil {
		conds.PhoneNO = &cruder.Cond{Op: cruder.EQ, Val: *h.PhoneNO}
	}

	stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), conds)
	if err != nil {
		return err
	}
	h.stm = stm.Select(
		entappuser.FieldEntID,
		entappuser.FieldAppID,
	)
	return nil
}

func (h *verifyHandler) queryAppUserByID(cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(
		cli.AppUser.Query(),
		&usercrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
		},
	)
	if err != nil {
		return err
	}

	h.stm = stm.Select(
		entappuser.FieldEntID,
		entappuser.FieldAppID,
	)
	return nil
}

func (h *verifyHandler) queryJoinAppUserSecret() {
	h.stm.Modify(func(s *sql.Selector) {
		t := sql.Table(entappusersecret.Table)
		s.LeftJoin(t).
			On(
				s.C(entappuser.FieldEntID),
				t.C(entappusersecret.FieldUserID),
			).
			On(
				s.C(entappuser.FieldAppID),
				t.C(entappusersecret.FieldAppID),
			).
			AppendSelect(
				sql.As(t.C(entappusersecret.FieldPasswordHash), "password_hash"),
				sql.As(t.C(entappusersecret.FieldSalt), "salt"),
				sql.As(t.C(entappusersecret.FieldUserID), "user_id"),
			)
	})
}

type r struct {
	EntID        uuid.UUID `sql:"ent_id"`
	AppID        uuid.UUID `sql:"app_id"`
	UserID       uuid.UUID `sql:"user_id"`
	PasswordHash string    `sql:"password_hash"`
	Salt         string    `sql:"salt"`
}

func (h *Handler) VerifyAccount(ctx context.Context) (*npool.User, error) {
	var infos []*r

	handler := &verifyHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUserByAccount(cli); err != nil {
			return err
		}
		handler.queryJoinAppUserSecret()
		return handler.stm.Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("invalid user")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	if h.PasswordHash == nil {
		return nil, fmt.Errorf("invalid password")
	}

	if err := encrypt.VerifyWithSalt(
		*h.PasswordHash,
		infos[0].PasswordHash,
		infos[0].Salt,
	); err != nil {
		return nil, err
	}

	h.EntID = &infos[0].UserID
	return h.GetUser(ctx)
}

func (h *Handler) VerifyUser(ctx context.Context) (*npool.User, error) {
	var infos []*r

	handler := &verifyHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUserByID(cli); err != nil {
			return err
		}
		handler.queryJoinAppUserSecret()
		return handler.stm.Scan(_ctx, &infos)
	})

	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("invalid user")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	if h.PasswordHash == nil {
		return nil, fmt.Errorf("invalid password")
	}

	if err := encrypt.VerifyWithSalt(
		*h.PasswordHash,
		infos[0].PasswordHash,
		infos[0].Salt,
	); err != nil {
		return nil, err
	}

	h.EntID = &infos[0].UserID
	return h.GetUser(ctx)
}
