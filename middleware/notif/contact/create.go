package contact

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

func (h *Handler) CreateContact(ctx context.Context) (info *npool.Contact, err error) {
	h.Conds = &crud.Conds{
		AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		AccountType: &cruder.Cond{Op: cruder.EQ, Val: basetypes.SignMethod(basetypes.SignMethod_value[h.AccountType.String()])},
		UsedFor:     &cruder.Cond{Op: cruder.EQ, Val: basetypes.UsedFor(basetypes.UsedFor_value[h.UsedFor.String()])},
	}
	exist, err := h.ExistContactConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("contact exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.Contact.Create(),
			&crud.Req{
				EntID:       h.EntID,
				AppID:       h.AppID,
				Account:     h.Account,
				AccountType: h.AccountType,
				UsedFor:     h.UsedFor,
				Sender:      h.Sender,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetContact(ctx)
}
