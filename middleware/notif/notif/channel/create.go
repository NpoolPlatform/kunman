package channel

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

func (h *Handler) CreateChannel(ctx context.Context) (info *npool.Channel, err error) {
	h.Conds = &crud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		Channel:   &cruder.Cond{Op: cruder.EQ, Val: basetypes.NotifChannel(basetypes.NotifChannel_value[h.Channel.String()])},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: basetypes.UsedFor(basetypes.UsedFor_value[h.EventType.String()])},
	}
	exist, err := h.ExistChannelConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("channel exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.NotifChannel.Create(),
			&crud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				Channel:   h.Channel,
				EventType: h.EventType,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetChannel(ctx)
}
