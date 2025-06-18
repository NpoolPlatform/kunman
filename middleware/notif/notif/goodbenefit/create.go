package goodbenefit

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.GoodID == nil {
		return fmt.Errorf("good id is empty")
	}
	if h.GoodName == nil {
		return fmt.Errorf("good name id is empty")
	}
	if h.State == nil {
		return fmt.Errorf("state is empty")
	}
	if h.BenefitDate == nil {
		return fmt.Errorf("benefit date is empty")
	}
	if *h.State == basetypes.Result_Success {
		if h.Amount != nil && h.Amount.Cmp(decimal.NewFromInt(0)) > 0 && h.TxID == nil {
			return fmt.Errorf("amount or tx id can not be empty")
		}
	}

	return nil
}

func (h *Handler) CreateGoodBenefit(ctx context.Context) (*npool.GoodBenefit, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		var amount *string
		if h.Amount != nil {
			_amount := h.Amount.String()
			amount = &_amount
		}
		if _, err := crud.CreateSet(
			cli.GoodBenefit.Create(),
			&crud.Req{
				EntID:       h.EntID,
				GoodID:      h.GoodID,
				GoodType:    h.GoodType,
				GoodName:    h.GoodName,
				CoinTypeID:  h.CoinTypeID,
				Amount:      amount,
				State:       h.State,
				Message:     h.Message,
				BenefitDate: h.BenefitDate,
				TxID:        h.TxID,
				Generated:   h.Generated,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodBenefit(ctx)
}
