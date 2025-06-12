package user

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/user"
	achievementusercrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/user"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entachievementuser "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/achievementuser"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AchievementUserSelect
	infos     []*npool.AchievementUser
	total     uint32
}

func (h *queryHandler) selectAchievementUser(stm *ent.AchievementUserQuery) {
	h.stmSelect = stm.Select(
		entachievementuser.FieldID,
		entachievementuser.FieldEntID,
		entachievementuser.FieldAppID,
		entachievementuser.FieldUserID,
		entachievementuser.FieldTotalCommission,
		entachievementuser.FieldSelfCommission,
		entachievementuser.FieldDirectInvitees,
		entachievementuser.FieldIndirectInvitees,
		entachievementuser.FieldDirectConsumeAmount,
		entachievementuser.FieldInviteeConsumeAmount,
		entachievementuser.FieldCreatedAt,
		entachievementuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAchievementUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AchievementUser.Query().Where(entachievementuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entachievementuser.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entachievementuser.EntID(*h.EntID))
	}
	h.selectAchievementUser(stm)
	return nil
}

func (h *queryHandler) queryAchievementUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := achievementusercrud.SetQueryConds(cli.AchievementUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectAchievementUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.TotalCommission)
		if err != nil {
			info.TotalCommission = decimal.NewFromInt(0).String()
		} else {
			info.TotalCommission = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfCommission)
		if err != nil {
			info.SelfCommission = decimal.NewFromInt(0).String()
		} else {
			info.SelfCommission = amount.String()
		}
		amount, err = decimal.NewFromString(info.DirectConsumeAmount)
		if err != nil {
			info.DirectConsumeAmount = decimal.NewFromInt(0).String()
		} else {
			info.DirectConsumeAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.InviteeConsumeAmount)
		if err != nil {
			info.InviteeConsumeAmount = decimal.NewFromInt(0).String()
		} else {
			info.InviteeConsumeAmount = amount.String()
		}
	}
}

func (h *Handler) GetAchievementUser(ctx context.Context) (*npool.AchievementUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAchievementUser(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetAchievementUsers(ctx context.Context) ([]*npool.AchievementUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AchievementUser{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAchievementUsers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
