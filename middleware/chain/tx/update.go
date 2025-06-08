package tx

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	txcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/tx"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	enttran "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/tran"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.Tran) error {
	if h.State == nil {
		return nil
	}

	switch info.State {
	case basetypes.TxState_TxStateCreated.String():
		switch *h.State {
		case basetypes.TxState_TxStateCreatedCheck:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateCreatedCheck.String():
		switch *h.State {
		case basetypes.TxState_TxStateWait:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateWait.String():
		switch *h.State {
		case basetypes.TxState_TxStateWaitCheck:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateWaitCheck.String():
		switch *h.State {
		case basetypes.TxState_TxStateTransferring:
		case basetypes.TxState_TxStateFail:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateTransferring.String():
		switch *h.State {
		case basetypes.TxState_TxStateSuccessful:
		case basetypes.TxState_TxStateFail:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateFail.String():
		switch *h.State {
		case basetypes.TxState_TxStateTransferring:
		default:
			return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
		}
	case basetypes.TxState_TxStateSuccessful.String():
		fallthrough //nolint
	default:
		return fmt.Errorf("state is invalid: %v -> %v", info.State, h.State)
	}

	return nil
}

func (h *Handler) UpdateTx(ctx context.Context) (*npool.Tx, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Tran.
			Query().
			Where(
				enttran.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := txcrud.UpdateSet(
			info.Update(),
			&txcrud.Req{
				ChainTxID: h.ChainTxID,
				State:     h.State,
				Extra:     h.Extra,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}
