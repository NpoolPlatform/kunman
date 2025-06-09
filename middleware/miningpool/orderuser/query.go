package orderuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	entgooduser "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/gooduser"
	entorderuser "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/orderuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/rootuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"

	orderusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/orderuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.OrderUserSelect
	infos []*npool.OrderUser
	total uint32
}

func (h *queryHandler) selectOrderUser(stm *ent.OrderUserQuery) {
	h.stm = stm.Select(
		entorderuser.FieldID,
		entorderuser.FieldEntID,
		entorderuser.FieldName,
		entorderuser.FieldGoodUserID,
		entorderuser.FieldAppID,
		entorderuser.FieldUserID,
		entorderuser.FieldReadPageLink,
		entorderuser.FieldCreatedAt,
		entorderuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOrderUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id or ent_id")
	}
	stm := cli.OrderUser.Query().Where(entorderuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderuser.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderuser.EntID(*h.EntID))
	}
	h.selectOrderUser(stm)
	return nil
}

func (h *queryHandler) queryOrderUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinCoinAndPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectOrderUser(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(h.queryJoinCoinAndPool)
}

func (h *queryHandler) queryJoinCoinAndPool(s *sql.Selector) {
	guT := sql.Table(entgooduser.Table)
	ruT := sql.Table(rootuser.Table)
	poolT := sql.Table(pool.Table)

	s.Join(guT).On(
		s.C(entorderuser.FieldGoodUserID),
		guT.C(entgooduser.FieldEntID),
	).OnP(
		sql.EQ(guT.C(entgooduser.FieldDeletedAt), 0),
	).Join(ruT).On(
		guT.C(entgooduser.FieldRootUserID),
		ruT.C(rootuser.FieldEntID),
	).OnP(
		sql.EQ(ruT.C(rootuser.FieldDeletedAt), 0),
	).Join(poolT).On(
		ruT.C(rootuser.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		pool.FieldMiningPoolType,
		sql.As(poolT.C(pool.FieldName), "mining_pool_name"),
		sql.As(poolT.C(pool.FieldLogo), "mining_pool_logo"),
		sql.As(poolT.C(pool.FieldSite), "mining_pool_site"),
		sql.As(poolT.C(pool.FieldEntID), "pool_id"),
		entgooduser.FieldRootUserID,
	)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningPoolType = mpbasetypes.MiningPoolType(mpbasetypes.MiningPoolType_value[info.MiningPoolTypeStr])
	}
}

func (h *Handler) GetOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUser(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetOrderUsers(ctx context.Context) ([]*npool.OrderUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUsers(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entorderuser.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetOrderUserProportion(ctx context.Context) (string, error) {
	info, err := h.GetOrderUser(ctx)
	zeroStr := decimal.Zero.String()
	if err != nil {
		return zeroStr, wlog.WrapError(err)
	}

	if info == nil {
		return zeroStr, wlog.Errorf("invalid entid")
	}

	handle := &baseInfoHandle{Handler: h}

	err = handle.getBaseInfo(ctx)
	if err != nil {
		return zeroStr, wlog.WrapError(err)
	}

	mgr, err := pools.NewPoolManager(handle.baseInfo.MiningPoolType, &handle.baseInfo.CoinType, handle.baseInfo.AuthToken)
	if err != nil {
		return zeroStr, wlog.WrapError(err)
	}

	proportion, err := mgr.GetRevenueProportion(ctx, handle.baseInfo.Distributor, handle.baseInfo.Recipient)
	if err != nil {
		return zeroStr, wlog.WrapError(err)
	}

	if proportion == nil {
		return zeroStr, nil
	}

	return decimal.NewFromFloat(*proportion).String(), nil
}

func (h *Handler) GetOrderUserBalance(ctx context.Context) (*npool.BalanceInfo, error) {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if info == nil {
		return nil, wlog.Errorf("invalid entid")
	}

	handle := &baseInfoHandle{Handler: h}

	err = handle.getBaseInfo(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	mgr, err := pools.NewPoolManager(handle.baseInfo.MiningPoolType, &handle.baseInfo.CoinType, handle.baseInfo.AuthToken)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	assetsBalance, err := mgr.GetAssetsBalance(ctx, handle.baseInfo.Recipient)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return &npool.BalanceInfo{
		Balance:              decimal.NewFromFloat(assetsBalance.Balance).String(),
		Paid:                 decimal.NewFromFloat(assetsBalance.Paid).String(),
		TotalIncome:          decimal.NewFromFloat(assetsBalance.TotalIncome).String(),
		YesterdayIncome:      decimal.NewFromFloat(assetsBalance.YesterdayIncome).String(),
		EstimatedTodayIncome: decimal.NewFromFloat(assetsBalance.EstimatedTodayIncome).String(),
	}, nil
}
