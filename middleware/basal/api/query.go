package api

import (
	"context"
	"encoding/json"
	"fmt"

	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

type queryHandler struct {
	*Handler
	stm   *ent.APISelect
	infos []*npool.API
	total uint32
}

func (h *queryHandler) selectAPI(stm *ent.APIQuery) {
	h.stm = stm.Select(
		entapi.FieldID,
		entapi.FieldEntID,
		entapi.FieldProtocol,
		entapi.FieldServiceName,
		entapi.FieldMethod,
		entapi.FieldMethodName,
		entapi.FieldPath,
		entapi.FieldPathPrefix,
		entapi.FieldDomains,
		entapi.FieldExported,
		entapi.FieldDeprecated,
		entapi.FieldCreatedAt,
		entapi.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAPI(cli *ent.Client) error {
	if h.EntID == nil && h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.API.Query().Where(entapi.DeletedAt(0))
	if h.EntID != nil {
		stm.Where(entapi.EntID(*h.EntID))
	}
	if h.ID != nil {
		stm.Where(entapi.ID(*h.ID))
	}
	h.selectAPI(stm)
	return nil
}

func (h *queryHandler) queryAPIsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.API.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAPI(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetDomains(ctx context.Context) (domains []string, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		domains, err = cli.
			API.
			Query().
			GroupBy(entapi.FieldServiceName).
			Strings(_ctx)
		return err
	})
	return
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Protocol = npool.Protocol(npool.Protocol_value[info.ProtocolStr])
		info.Method = npool.Method(npool.Method_value[info.MethodStr])
		_ = json.Unmarshal([]byte(info.DomainsStr), &info.Domains)
	}
}

func (h *Handler) GetAPIs(ctx context.Context) ([]*npool.API, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAPIsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.stm.Offset(int(h.Offset)).Limit(int(h.Limit))

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

func (h *Handler) GetAPI(ctx context.Context) (*npool.API, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAPI(cli); err != nil {
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

	return handler.infos[0], nil
}

func (h *Handler) GetAPIOnly(ctx context.Context) (*npool.API, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAPIsByConds(_ctx, cli); err != nil {
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

	return handler.infos[0], nil
}
