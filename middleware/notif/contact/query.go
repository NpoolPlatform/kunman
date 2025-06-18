package contact

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/contact"
)

type queryHandler struct {
	*Handler
	stm   *ent.ContactSelect
	infos []*npool.Contact
	total uint32
}

func (h *queryHandler) selectContact(stm *ent.ContactQuery) {
	h.stm = stm.Select(
		entamt.FieldID,
		entamt.FieldEntID,
		entamt.FieldAppID,
		entamt.FieldAccount,
		entamt.FieldAccountType,
		entamt.FieldUsedFor,
		entamt.FieldSender,
		entamt.FieldCreatedAt,
		entamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.AccountType = basetypes.SignMethod(basetypes.SignMethod_value[info.AccountTypeStr])
		info.UsedFor = basetypes.UsedFor(basetypes.UsedFor_value[info.UsedForStr])
	}
}

func (h *queryHandler) queryContact(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Contact.Query().Where(entamt.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entamt.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entamt.EntID(*h.EntID))
	}
	h.selectContact(stm)
	return nil
}

func (h *queryHandler) queryContactsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.Contact.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectContact(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetContacts(ctx context.Context) ([]*npool.Contact, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContactsByConds(_ctx, cli); err != nil {
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

func (h *Handler) GetContact(ctx context.Context) (info *npool.Contact, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContact(cli); err != nil {
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

func (h *Handler) GetContactOnly(ctx context.Context) (*npool.Contact, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContactsByConds(_ctx, cli); err != nil {
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
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("to many record")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
