package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entappuserthirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuserthirdparty"

	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm            *ent.AppUserSelect
	joinThirdParty bool
}

func (h *existHandler) selectAppUser(stm *ent.AppUserQuery) {
	h.stm = stm.Select(entappuser.FieldID)
}

func (h *existHandler) queryAppUser(cli *ent.Client) {
	stm := cli.AppUser.
		Query().
		Where(entappuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappuser.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entappuser.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entappuser.EntID(*h.EntID))
	}
	h.selectAppUser(stm)
}

func (h *existHandler) queryAppUserByConds(cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectAppUser(stm)
	return nil
}

func (h *existHandler) queryJoinAppUserThirdParty(s *sql.Selector) error {
	if !h.joinThirdParty {
		return nil
	}

	t := sql.Table(entappuserthirdparty.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldEntID),
			t.C(entappuserthirdparty.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entappuserthirdparty.FieldAppID),
		).
		On(
			s.C(entappuser.FieldDeletedAt),
			t.C(entappuserthirdparty.FieldDeletedAt),
		)

	if h.Conds != nil && h.Conds.ThirdPartyID != nil {
		thirdPartyID, ok := h.Conds.ThirdPartyID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyID), thirdPartyID),
		)
	}

	if h.Conds != nil && h.Conds.ThirdPartyUserID != nil {
		thirdPartyUserID, ok := h.Conds.ThirdPartyUserID.Val.(string)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyuserid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyUserID), thirdPartyUserID),
		)
	}

	return nil
}

func (h *existHandler) queryJoin() error {
	var err error
	h.stm.Modify(func(s *sql.Selector) {
		err = h.queryJoinAppUserThirdParty(s)
	})
	return err
}

func (h *Handler) ExistUser(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		Handler:        h,
		joinThirdParty: false,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppUser(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		count, err := handler.stm.Limit(1).Count(_ctx)
		exist = count > 0
		return err
	})
	return exist, err
}

func (h *Handler) ExistUserConds(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		Handler:        h,
		joinThirdParty: false,
	}
	if h.Conds != nil && (h.Conds.ThirdPartyID != nil || h.Conds.ThirdPartyUserID != nil) {
		handler.joinThirdParty = true
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUserByConds(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		count, err := handler.stm.Limit(1).Count(_ctx)
		exist = count > 0
		return err
	})
	return exist, err
}
