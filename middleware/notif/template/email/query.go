package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	entemailtemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/emailtemplate"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	emailtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/email"
)

type queryHandler struct {
	*Handler
	stm   *ent.EmailTemplateSelect
	infos []*npool.EmailTemplate
	total uint32
}

func (h *queryHandler) selectEmailTemplate(stm *ent.EmailTemplateQuery) {
	h.stm = stm.Select(
		entemailtemplate.FieldID,
		entemailtemplate.FieldEntID,
		entemailtemplate.FieldAppID,
		entemailtemplate.FieldLangID,
		entemailtemplate.FieldUsedFor,
		entemailtemplate.FieldSender,
		entemailtemplate.FieldReplyTos,
		entemailtemplate.FieldCcTos,
		entemailtemplate.FieldSubject,
		entemailtemplate.FieldBody,
		entemailtemplate.FieldDefaultToUsername,
		entemailtemplate.FieldCreatedAt,
		entemailtemplate.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryEmailTemplate(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.EmailTemplate.Query().Where(entemailtemplate.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entemailtemplate.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entemailtemplate.EntID(*h.EntID))
	}
	h.selectEmailTemplate(stm)
	return nil
}

func (h *queryHandler) formalize() error {
	for _, info := range h.infos {
		info.UsedFor = basetypes.UsedFor(basetypes.UsedFor_value[info.UsedForStr])
		if err := json.Unmarshal([]byte(info.ReplyTosStr), &info.ReplyTos); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(info.CCTosStr), &info.CCTos); err != nil {
			return err
		}
	}
	return nil
}

func (h *queryHandler) queryEmailTemplates(ctx context.Context, cli *ent.Client) error {
	stm, err := emailtemplatecrud.SetQueryConds(cli.EmailTemplate.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectEmailTemplate(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplate(cli); err != nil {
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
	if err := handler.formalize(); err != nil {
		return nil, err
	}

	return handler.infos[0], nil
}

func (h *Handler) GetEmailTemplates(ctx context.Context) ([]*npool.EmailTemplate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplates(ctx, cli); err != nil {
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
	if err := handler.formalize(); err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetEmailTemplateOnly(ctx context.Context) (info *npool.EmailTemplate, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplates(_ctx, cli); err != nil {
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
	if err := handler.formalize(); err != nil {
		return nil, err
	}

	return handler.infos[0], nil
}
