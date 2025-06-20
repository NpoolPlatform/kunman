package app

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	appcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entappctrl "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appcontrol"
	entbanapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	appusertypes "github.com/NpoolPlatform/kunman/message/basetypes/appuser/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppSelect
	stmCount  *ent.AppSelect
	infos     []*npool.App
	total     uint32
}

func (h *queryHandler) selectApp(stm *ent.AppQuery) *ent.AppSelect {
	return stm.Select(entapp.FieldID)
}

func (h *queryHandler) queryApp(cli *ent.Client) {
	stm := cli.App.
		Query().
		Where(entapp.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapp.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entapp.EntID(*h.EntID))
	}
	h.stmSelect = h.selectApp(stm)
}

func (h *queryHandler) queryApps(cli *ent.Client) (*ent.AppSelect, error) {
	stm, err := appcrud.SetQueryConds(cli.App.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectApp(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapp.FieldID),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			t.C(entapp.FieldEntID),
			t.C(entapp.FieldCreatedBy),
			t.C(entapp.FieldLogo),
			t.C(entapp.FieldName),
			t.C(entapp.FieldDescription),
			t.C(entapp.FieldCreatedAt),
		)
}

func (h *queryHandler) queryJoinAppCtrl(s *sql.Selector) {
	t := sql.Table(entappctrl.Table)
	s.LeftJoin(t).
		On(
			s.C(entapp.FieldEntID),
			t.C(entappctrl.FieldAppID),
		).
		AppendSelect(
			t.C(entappctrl.FieldSignupMethods),
			t.C(entappctrl.FieldExternSigninMethods),
			t.C(entappctrl.FieldRecaptchaMethod),
			t.C(entappctrl.FieldKycEnable),
			t.C(entappctrl.FieldSigninVerifyEnable),
			t.C(entappctrl.FieldInvitationCodeMust),
			t.C(entappctrl.FieldCreateInvitationCodeWhen),
			t.C(entappctrl.FieldMaxTypedCouponsPerOrder),
			t.C(entappctrl.FieldMaintaining),
			t.C(entappctrl.FieldCommitButtonTargets),
			t.C(entappctrl.FieldCouponWithdrawEnable),
			t.C(entappctrl.FieldResetUserMethod),
		)
}

func (h *queryHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapp.FieldEntID),
			t.C(entbanapp.FieldAppID),
		).
		On(
			s.C(entapp.FieldDeletedAt),
			t.C(entbanapp.FieldDeletedAt),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldAppID), "ban_app_id"),
			sql.As(t.C(entbanapp.FieldMessage), "ban_message"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppCtrl(s)
		h.queryJoinBanApp(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppCtrl(s)
		h.queryJoinBanApp(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CreateInvitationCodeWhen = basetypes.CreateInvitationCodeWhen(
			basetypes.CreateInvitationCodeWhen_value[info.CreateInvitationCodeWhenStr],
		)
		_ = json.Unmarshal([]byte(info.CommitButtonTargetsStr), &info.CommitButtonTargets)

		methods := []string{}
		_methods := []basetypes.SignMethod{}

		_ = json.Unmarshal([]byte(info.SignupMethodsStr), &methods)
		for _, m := range methods {
			_methods = append(_methods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
		}

		emethods := []string{}
		_emethods := []basetypes.SignMethod{}

		_ = json.Unmarshal([]byte(info.ExtSigninMethodsStr), &emethods)
		for _, m := range emethods {
			_emethods = append(_emethods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
		}

		info.SignupMethods = _methods
		info.ExtSigninMethods = _emethods
		info.RecaptchaMethod = basetypes.RecaptchaMethod(basetypes.RecaptchaMethod_value[info.RecaptchaMethodStr])

		info.Banned = info.BanAppID != ""
		info.ResetUserMethod = appusertypes.ResetUserMethod(appusertypes.ResetUserMethod_value[info.ResetUserMethodStr])
	}
}

func (h *Handler) GetApp(ctx context.Context) (*npool.App, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryApp(cli)
		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetApps(ctx context.Context) ([]*npool.App, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		var err error
		if handler.stmSelect, err = handler.queryApps(cli); err != nil {
			return err
		}
		if handler.stmCount, err = handler.queryApps(cli); err != nil {
			return err
		}

		total, err := handler.stmCount.Count(ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
