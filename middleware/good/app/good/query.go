package good

import (
	"context"
	"sort"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaynamecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/name"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgooddisplayname "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddisplayname"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/name"
	"github.com/google/uuid"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount     *ent.AppGoodBaseSelect
	infos        []*npool.Good
	total        uint32
	displayNames []*appgooddisplaynamemwpb.DisplayNameInfo
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {
		if err := h.queryJoinGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodBase", "Error", err)
		}
	})
}

func (h *queryHandler) getDisplayNames(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.EntID))
		}
		return
	}()

	stm, err := appgooddisplaynamecrud.SetQueryConds(
		cli.AppGoodDisplayName.Query(),
		&appgooddisplaynamecrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgooddisplayname.FieldAppGoodID,
		entappgooddisplayname.FieldName,
		entappgooddisplayname.FieldIndex,
	).Scan(ctx, &h.displayNames)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	displayNames := map[string][]string{}
	displayNameMaps := map[string][]*appgooddisplaynamemwpb.DisplayNameInfo{}
	for _, displayName := range h.displayNames {
		displayNameMaps[displayName.AppGoodID] = append(displayNameMaps[displayName.AppGoodID], displayName)
	}
	for appGoodID, nameInfos := range displayNameMaps {
		sort.Slice(nameInfos, func(i, j int) bool {
			return nameInfos[i].Index < nameInfos[j].Index
		})

		names := []string{}
		for _, displayName := range nameInfos {
			names = append(names, displayName.Name)
		}
		displayNames[appGoodID] = names
	}

	for _, info := range h.infos {
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.State = types.GoodState(types.GoodState_value[info.StateStr])
		info.DisplayNames = displayNames[info.EntID]
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryGood(cli)
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getDisplayNames(_ctx, cli)
	}); err != nil {
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

func (h *Handler) GetGoods(ctx context.Context) (infos []*npool.Good, total uint32, err error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryGoods(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryGoods(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getDisplayNames(_ctx, cli)
	}); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
