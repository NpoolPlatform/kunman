package notif

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	entnotif "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notif"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifSelect
	infos []*npool.Notif
	total uint32
}

func (h *queryHandler) selectNotif(stm *ent.NotifQuery) {
	h.stm = stm.Select(
		entnotif.FieldID,
		entnotif.FieldEntID,
		entnotif.FieldAppID,
		entnotif.FieldLangID,
		entnotif.FieldUserID,
		entnotif.FieldEventID,
		entnotif.FieldNotified,
		entnotif.FieldEventType,
		entnotif.FieldUseTemplate,
		entnotif.FieldTitle,
		entnotif.FieldContent,
		entnotif.FieldChannel,
		entnotif.FieldExtra,
		entnotif.FieldType,
		entnotif.FieldCreatedAt,
		entnotif.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryNotif(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Notif.Query().Where(entnotif.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entnotif.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entnotif.EntID(*h.EntID))
	}
	h.selectNotif(stm)
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
		info.NotifType = basetypes.NotifType(basetypes.NotifType_value[info.NotifTypeStr])
	}
}

func (h *queryHandler) queryNotifs(ctx context.Context, cli *ent.Client) error {
	stm, err := notifcrud.SetQueryConds(cli.Notif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectNotif(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetNotif(ctx context.Context) (*npool.Notif, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotif(cli); err != nil {
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

func (h *Handler) GetNotifs(ctx context.Context) ([]*npool.Notif, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(ctx, cli); err != nil {
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

func (h *Handler) GetNotifOnly(ctx context.Context) (info *npool.Notif, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(_ctx, cli); err != nil {
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
