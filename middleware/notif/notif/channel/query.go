package channel

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entchannel "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notifchannel"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifChannelSelect
	infos []*npool.Channel
	total uint32
}

func (h *queryHandler) selectChannel(stm *ent.NotifChannelQuery) {
	h.stm = stm.Select(
		entchannel.FieldID,
		entchannel.FieldEntID,
		entchannel.FieldAppID,
		entchannel.FieldChannel,
		entchannel.FieldEventType,
		entchannel.FieldCreatedAt,
		entchannel.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
	}
}

func (h *queryHandler) queryChannel(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.NotifChannel.Query().Where(entchannel.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entchannel.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entchannel.EntID(*h.EntID))
	}
	h.selectChannel(stm)
	return nil
}

func (h *queryHandler) queryChannelsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.NotifChannel.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectChannel(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetChannels(ctx context.Context) ([]*npool.Channel, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannelsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetChannel(ctx context.Context) (info *npool.Channel, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannel(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetChannelOnly(ctx context.Context) (*npool.Channel, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannelsByConds(_ctx, cli); err != nil {
			return err
		}

		_, err := handler.stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("to many record")
	}

	return handler.infos[0], nil
}
