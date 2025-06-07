package appsimulatepowerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappsimulatepowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsimulatepowerrental"
	apppowerrental1 "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
)

type appPowerRentalHandler struct {
	*Handler
	appPowerRental apppowerrental1.PowerRental
	simulate       *ent.AppSimulatePowerRental
}

func (h *appPowerRentalHandler) queryAppPowerRentalEnt(ctx context.Context) (err error) {
	handler, err := apppowerrental1.NewHandler(
		ctx,
		apppowerrental1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.appPowerRental, err = handler.QueryPowerRentalEnt(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *appPowerRentalHandler) queryAppPowerRental(ctx context.Context) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid simulateid")
	}

	if h.AppGoodID != nil {
		return h.queryAppPowerRentalEnt(ctx)
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppSimulatePowerRental.Query()
		if h.ID != nil {
			stm.Where(entappsimulatepowerrental.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappsimulatepowerrental.EntID(*h.EntID))
		}
		h.simulate, err = stm.Only(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &h.simulate.AppGoodID
		return nil
	}); err != nil {
		return wlog.WrapError(err)
	}

	return h.queryAppPowerRentalEnt(ctx)
}
