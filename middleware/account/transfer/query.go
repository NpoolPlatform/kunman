package transfer

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	enttransfer "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/transfer"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/transfer"
	transfercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/transfer"
)

type queryHandler struct {
	*Handler
	stm   *ent.TransferSelect
	infos []*npool.Transfer
	total uint32
}

func (h *queryHandler) selectTransfer(stm *ent.TransferQuery) {
	h.stm = stm.Select(
		enttransfer.FieldID,
		enttransfer.FieldEntID,
		enttransfer.FieldAppID,
		enttransfer.FieldUserID,
		enttransfer.FieldTargetUserID,
		enttransfer.FieldCreatedAt,
		enttransfer.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTransfer(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Transfer.Query().Where(enttransfer.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttransfer.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttransfer.EntID(*h.EntID))
	}
	h.selectTransfer(stm)
	return nil
}

func (h *queryHandler) queryTransfers(ctx context.Context, cli *ent.Client) error {
	stm, err := transfercrud.SetQueryConds(cli.Transfer.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectTransfer(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTransfer(ctx context.Context) (*npool.Transfer, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfer(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
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

	return handler.infos[0], nil
}

func (h *Handler) GetTransfers(ctx context.Context) ([]*npool.Transfer, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfers(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(enttransfer.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
