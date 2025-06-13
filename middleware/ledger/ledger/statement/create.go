package statement

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	profitcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/profit"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createOrUpdateProfit(ctx context.Context, tx *ent.Tx, req *crud.Req) error {
	if *req.IOSubType != types.IOSubType_MiningBenefit && *req.IOSubType != types.IOSubType_SimulateMiningBenefit {
		return nil
	}

	stm, err := profitcrud.SetQueryConds(
		tx.Profit.Query(),
		&profitcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	// create
	if info == nil {
		if _, err := profitcrud.CreateSet(
			tx.Profit.Create(),
			&profitcrud.Req{
				AppID:      req.AppID,
				UserID:     req.UserID,
				CoinTypeID: req.CoinTypeID,
				Incoming:   req.Amount,
			}).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	stm1, err := profitcrud.UpdateSetWithValidate(
		info,
		&profitcrud.Req{
			AppID:      req.AppID,
			UserID:     req.UserID,
			CoinTypeID: req.CoinTypeID,
			Incoming:   req.Amount,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createStatement(ctx context.Context, tx *ent.Tx, req *crud.Req) error {
	stm, err := crud.SetQueryConds(tx.Statement.Query(), &crud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		IOType:     &cruder.Cond{Op: cruder.EQ, Val: *req.IOType},
		IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *req.IOSubType},
		IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *req.IOExtra},
	})
	if err != nil {
		return err
	}
	exist, err := stm.Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("statement already exist")
	}
	if _, err := crud.CreateSet(
		tx.Statement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createOrUpdateLedger(ctx context.Context, tx *ent.Tx, req *crud.Req) error {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
	})
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	incoming := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *req.IOType {
	case types.IOType_Incoming:
		incoming = decimal.RequireFromString(req.Amount.String())
	case types.IOType_Outcoming:
		outcoming = decimal.RequireFromString(req.Amount.String())
	default:
		return fmt.Errorf("invalid io type %v", *req.IOType)
	}

	spendable := incoming.Sub(outcoming)
	locked := decimal.NewFromInt(0)

	if info == nil {
		if incoming.Cmp(decimal.NewFromInt(0)) < 0 || spendable.Cmp(decimal.NewFromInt(0)) < 0 ||
			outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("insufficient funds")
		}
		if _, err := ledgercrud.CreateSet(tx.Ledger.Create(), &ledgercrud.Req{
			AppID:      req.AppID,
			UserID:     req.UserID,
			CoinTypeID: req.CoinTypeID,
			Incoming:   &incoming,
			Outcoming:  &outcoming,
			Locked:     &locked,
			Spendable:  &spendable,
		}).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	stm1, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming:  &incoming,
			Outcoming: &outcoming,
			Spendable: &spendable,
			Locked:    &locked,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				id := uuid.New()
				if req.EntID == nil {
					req.EntID = &id
				}
				if err := handler.createStatement(ctx, tx, req); err != nil {
					return err
				}
				if err := handler.createOrUpdateProfit(ctx, tx, req); err != nil {
					return err
				}
				if err := handler.createOrUpdateLedger(ctx, tx, req); err != nil {
					return err
				}
				ids = append(ids, *req.EntID)
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &crud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (h *Handler) validate() error { //nolint:gocyclo
	switch *h.IOType {
	case types.IOType_Incoming:
		switch *h.IOSubType {
		case types.IOSubType_Payment:
		case types.IOSubType_MiningBenefit:
		case types.IOSubType_Commission:
		case types.IOSubType_TechniqueFeeCommission:
		case types.IOSubType_Deposit:
		case types.IOSubType_Transfer:
		case types.IOSubType_OrderRevoke:
		case types.IOSubType_SimulateMiningBenefit:
		case types.IOSubType_ObseletePayment:
		case types.IOSubType_EventReward:
		default:
			return fmt.Errorf("io subtype not match io type")
		}
	case types.IOType_Outcoming:
		switch *h.IOSubType {
		case types.IOSubType_Payment:
		case types.IOSubType_Withdrawal:
		case types.IOSubType_Transfer:
		case types.IOSubType_CommissionRevoke:
		default:
			return fmt.Errorf("io subtype not match io type")
		}
	default:
		return fmt.Errorf("invalid io type %v", *h.IOType)
	}
	return nil
}

func (h *Handler) CreateStatement(ctx context.Context) (*npool.Statement, error) {
	if err := h.validate(); err != nil {
		return nil, err
	}
	h.Reqs = []*crud.Req{&h.Req}
	infos, err := h.CreateStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos[0], nil
}
