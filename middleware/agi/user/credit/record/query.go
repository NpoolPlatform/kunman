package record

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/user/credit/record"
)

type queryHandler struct {
	*baseQueryHandler
	infos []*npool.Record
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.OperationType = types.OperationType(types.OperationType_value[info.OperationTypeStr])
	}
}

func (h *Handler) GetRecord(ctx context.Context) (*npool.Record, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRecord(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
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

func (h *Handler) GetRecords(ctx context.Context) ([]*npool.Record, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.queryRecords(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, nil
}
