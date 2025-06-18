package appcountry

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/appcountry"
	countrymw "github.com/NpoolPlatform/kunman/middleware/g11n/country"
	appcountrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/appcountry"
	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkRepeat() error {
	countryMap := map[string]*uuid.UUID{}
	for _, req := range h.Reqs {
		_, ok := countryMap[req.AppID.String()+req.CountryID.String()]
		if ok {
			return fmt.Errorf("duplicate countryid")
		}
		countryMap[req.AppID.String()+req.CountryID.String()] = req.CountryID
	}
	return nil
}

func (h *createHandler) createCountry(ctx context.Context, tx *ent.Tx, req *appcountrycrud.Req) error {
	handler, err := countrymw.NewHandler(
		ctx,
	)
	if err != nil {
		return err
	}
	handler.Conds = &countrycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *req.CountryID},
	}
	existCountry, err := handler.ExistCountryCondsWithTx(ctx, tx)
	if err != nil {
		return err
	}
	if !existCountry {
		return fmt.Errorf("country not exist")
	}

	h.Conds = &appcountrycrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		CountryID: &cruder.Cond{Op: cruder.EQ, Val: *req.CountryID},
	}
	exist, err := h.ExistAppCountryCondsWithTx(ctx, tx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("appcountry exist")
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := appcountrycrud.CreateSet(
		tx.AppCountry.Create(),
		&appcountrycrud.Req{
			EntID:     req.EntID,
			AppID:     req.AppID,
			CountryID: req.CountryID,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &appcountrycrud.Req{
			EntID:     h.EntID,
			AppID:     h.AppID,
			CountryID: h.CountryID,
		}
		if err := handler.createCountry(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCountry(ctx)
}

func (h *Handler) CreateCountries(ctx context.Context) ([]*npool.Country, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.checkRepeat(); err != nil {
				return err
			}
			if err := handler.createCountry(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &appcountrycrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
