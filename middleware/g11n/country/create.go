package country

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"
	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	countryIDs map[string]*uuid.UUID
}

func (h *createHandler) createCountry(ctx context.Context, tx *ent.Tx, req *countrycrud.Req) error {
	_id, ok := h.countryIDs[*req.Country]
	if ok {
		h.EntID = _id
		return nil
	}

	h.Conds = &countrycrud.Conds{
		Country: &cruder.Cond{Op: cruder.EQ, Val: *req.Country},
	}
	h.Limit = 2
	infos, _, err := h.GetCountriesWithTx(ctx, tx)
	if err != nil {
		return err
	}
	if infos != nil {
		id := uuid.MustParse(infos[0].EntID)
		h.EntID = &id
		return nil
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := countrycrud.CreateSet(
		tx.Country.Create(),
		&countrycrud.Req{
			EntID:   req.EntID,
			Country: req.Country,
			Flag:    req.Flag,
			Code:    req.Code,
			Short:   req.Short,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	h.countryIDs[*req.Country] = h.EntID

	return nil
}

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	handler := &createHandler{
		Handler:    h,
		countryIDs: map[string]*uuid.UUID{},
	}
	h.Conds = &countrycrud.Conds{
		Country: &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
	}
	exist, err := h.ExistCountryConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("country exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &countrycrud.Req{
			EntID:   h.EntID,
			Country: h.Country,
			Flag:    h.Flag,
			Code:    h.Code,
			Short:   h.Short,
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
		Handler:    h,
		countryIDs: map[string]*uuid.UUID{},
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
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

	h.Conds = &countrycrud.Conds{
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
