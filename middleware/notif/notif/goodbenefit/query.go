package goodbenefit

import (
	"context"
	"fmt"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entgoodbenefit "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/goodbenefit"
)

type queryHandler struct {
	*Handler
	stm   *ent.GoodBenefitSelect
	infos []*npool.GoodBenefit
	total uint32
}

func (h *queryHandler) selectGoodBenefit(stm *ent.GoodBenefitQuery) {
	h.stm = stm.Select(
		entgoodbenefit.FieldID,
		entgoodbenefit.FieldEntID,
		entgoodbenefit.FieldGoodID,
		entgoodbenefit.FieldGoodType,
		entgoodbenefit.FieldGoodName,
		entgoodbenefit.FieldCoinTypeID,
		entgoodbenefit.FieldAmount,
		entgoodbenefit.FieldState,
		entgoodbenefit.FieldMessage,
		entgoodbenefit.FieldBenefitDate,
		entgoodbenefit.FieldTxID,
		entgoodbenefit.FieldGenerated,
		entgoodbenefit.FieldCreatedAt,
		entgoodbenefit.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.State = basetypes.Result(basetypes.Result_value[info.StateStr])
		info.GoodType = goodtypes.GoodType(goodtypes.GoodType_value[info.GoodTypeStr])
	}
}

func (h *queryHandler) queryGoodBenefit(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.GoodBenefit.Query().Where(entgoodbenefit.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entgoodbenefit.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodbenefit.EntID(*h.EntID))
	}
	h.selectGoodBenefit(stm)
	return nil
}

func (h *queryHandler) queryGoodBenefitsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.GoodBenefit.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectGoodBenefit(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetGoodBenefits(ctx context.Context) ([]*npool.GoodBenefit, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBenefitsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetGoodBenefit(ctx context.Context) (info *npool.GoodBenefit, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBenefit(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
