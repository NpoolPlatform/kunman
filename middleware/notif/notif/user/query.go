package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/user"
	entnotifuser "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notifuser"

	usercrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifUserSelect
	infos []*npool.NotifUser
	total uint32
}

func (h *queryHandler) selectNotifUser(stm *ent.NotifUserQuery) {
	h.stm = stm.Select(
		entnotifuser.FieldID,
		entnotifuser.FieldEntID,
		entnotifuser.FieldAppID,
		entnotifuser.FieldUserID,
		entnotifuser.FieldEventType,
		entnotifuser.FieldCreatedAt,
		entnotifuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryNotifUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.NotifUser.Query().Where(entnotifuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entnotifuser.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entnotifuser.EntID(*h.EntID))
	}
	h.selectNotifUser(stm)
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
	}
}

func (h *queryHandler) queryNotifUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.NotifUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectNotifUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUser(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetNotifUsers(ctx context.Context) ([]*npool.NotifUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUsers(ctx, cli); err != nil {
			return err
		}
		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetNotifUserOnly(ctx context.Context) (info *npool.NotifUser, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUsers(_ctx, cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
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
	handler.formalize()

	return handler.infos[0], nil
}
