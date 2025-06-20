package withdraw

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	statementcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"
	entwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/withdraw"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	withdraw *ent.Withdraw
}

var (
	stateMap    map[types.WithdrawState][]types.WithdrawState
	rollbackMap map[types.WithdrawState]types.WithdrawState
)

func init() {
	stateMap = map[types.WithdrawState][]types.WithdrawState{
		types.WithdrawState_Created:                {types.WithdrawState_Reviewing},
		types.WithdrawState_Reviewing:              {types.WithdrawState_Approved, types.WithdrawState_PreRejected},
		types.WithdrawState_Approved:               {types.WithdrawState_Transferring},
		types.WithdrawState_Transferring:           {types.WithdrawState_PreFail, types.WithdrawState_PreSuccessful},
		types.WithdrawState_PreRejected:            {types.WithdrawState_ReturnRejectedBalance},
		types.WithdrawState_ReturnRejectedBalance:  {types.WithdrawState_Rejected},
		types.WithdrawState_PreFail:                {types.WithdrawState_ReturnFailBalance},
		types.WithdrawState_ReturnFailBalance:      {types.WithdrawState_TransactionFail},
		types.WithdrawState_PreSuccessful:          {types.WithdrawState_SpendSuccessfulBalance},
		types.WithdrawState_SpendSuccessfulBalance: {types.WithdrawState_Successful},
	}
	rollbackMap = map[types.WithdrawState]types.WithdrawState{
		types.WithdrawState_TransactionFail: types.WithdrawState_ReturnFailBalance,
		types.WithdrawState_Rejected:        types.WithdrawState_ReturnRejectedBalance,
		types.WithdrawState_Successful:      types.WithdrawState_SpendSuccessfulBalance,
	}
}

// nolint
func (h *updateHandler) checkWithdrawState(ctx context.Context) error {
	if h.State == nil && h.PlatformTransactionID == nil && h.ReviewID == nil {
		return nil
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Withdraw.
			Query().
			Where(
				entwithdraw.ID(*h.ID),
				entwithdraw.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.withdraw = info
		return nil
	})
	if err != nil {
		return err
	}
	if h.State == nil {
		return nil
	}

	state := types.WithdrawState(types.WithdrawState_value[h.withdraw.State])

	switch state {
	case types.WithdrawState_TransactionFail:
		fallthrough //nolint
	case types.WithdrawState_Successful:
		return fmt.Errorf("permission denied")
	}

	if h.Rollback == nil || !*h.Rollback {
		switch state {
		case types.WithdrawState_Rejected:
			fallthrough //nolint
		case types.WithdrawState_TransactionFail:
			fallthrough //nolint
		case types.WithdrawState_Successful:
			return fmt.Errorf("permission denied")
		}
	}

	toStates := []types.WithdrawState{}
	if h.Rollback != nil && *h.Rollback &&
		*h.State == types.WithdrawState(types.WithdrawState_value[h.withdraw.State]) {
		toState, ok := rollbackMap[*h.State]
		if !ok {
			return fmt.Errorf("invalid rollback state")
		}
		h.State = &toState
		return nil
	}

	switch *h.State {
	case types.WithdrawState_Reviewing:
		if h.ReviewID == nil {
			return fmt.Errorf("invalid review id")
		}
	case types.WithdrawState_Transferring:
		if h.PlatformTransactionID == nil {
			return fmt.Errorf("invalid platform transaction id")
		}
	}
	toStates = stateMap[state]
	allow := false
	for _, state := range toStates {
		if state == *h.State {
			allow = true
			break
		}
	}
	if !allow {
		return fmt.Errorf("permission denied %v -> %v", h.withdraw.State, *h.State)
	}

	return nil
}

func (h *updateHandler) updateLedger(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
		return nil
	}

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.withdraw.AppID),
			entledger.UserID(h.withdraw.UserID),
			entledger.CurrencyID(h.withdraw.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	spendable := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *h.State {
	case types.WithdrawState_Successful:
		outcoming = h.withdraw.Amount
	case types.WithdrawState_TransactionFail:
		fallthrough //nolint
	case types.WithdrawState_Rejected:
		spendable = h.withdraw.Amount
	default:
		return nil
	}

	locked := decimal.NewFromInt(0).Sub(h.withdraw.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Spendable: &spendable,
			Outcoming: &outcoming,
			Locked:    &locked,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) checkWithdraw(ctx context.Context) error {
	if h.PlatformTransactionID == nil && h.ReviewID == nil {
		return nil
	}

	if h.withdraw.PlatformTransactionID != uuid.Nil &&
		h.PlatformTransactionID != nil &&
		*h.PlatformTransactionID != h.withdraw.PlatformTransactionID {
		return fmt.Errorf("permission denied")
	}
	if h.withdraw.ReviewID != uuid.Nil &&
		h.ReviewID != nil &&
		*h.ReviewID != h.withdraw.ReviewID {
		return fmt.Errorf("permission denied")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Withdraw.Query().Where(entwithdraw.IDNEQ(*h.ID))
		if h.PlatformTransactionID != nil {
			stm.Where(
				entwithdraw.PlatformTransactionID(*h.PlatformTransactionID),
			)
		}
		if h.ReviewID != nil {
			stm.Where(
				entwithdraw.ReviewID(*h.ReviewID),
			)
		}
		exist, err := stm.Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("already exists")
		}
		return nil
	})
}

func (h *updateHandler) updateWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.UpdateSet(
		tx.Withdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
		return nil
	}
	if h.State.String() != types.WithdrawState_Successful.String() {
		return nil
	}
	if h.FeeAmount == nil {
		return fmt.Errorf("invalid feeamount")
	}

	ioExtra := fmt.Sprintf(
		`{"WithdrawID":"%v","TransactionID":"%v","CID":"%v","TransactionFee":"%v","AccountID":"%v"}`,
		h.withdraw.EntID,
		h.withdraw.PlatformTransactionID.String(),
		h.withdraw.ChainTransactionID,
		h.FeeAmount.String(),
		h.withdraw.AccountID,
	)

	ioType := types.IOType_Outcoming
	ioSubType := types.IOSubType_Withdrawal
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:        &h.withdraw.AppID,
			UserID:       &h.withdraw.UserID,
			CurrencyID:   &h.withdraw.CoinTypeID,
			CurrencyType: types.CurrencyType_CurrencyCrypto.Enum(),
			IOType:       &ioType,
			IOSubType:    &ioSubType,
			IOExtra:      &ioExtra,
			Amount:       &h.withdraw.Amount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkWithdrawState(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkWithdraw(ctx); err != nil {
		return nil, err
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetWithdraw(ctx)
}
