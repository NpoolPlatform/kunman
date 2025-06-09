package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	orderusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/orderuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	orderuserent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/orderuser"
)

type existHandler struct {
	*Handler
	checkAppAuthSQL string
}

func (h *existHandler) existJoinCoinAndAppPool() error {
	if h.AppID == nil || h.GoodUserID == nil {
		return wlog.Errorf("invalid appid or gooduserid")
	}

	sql := fmt.Sprintf(`SELECT DISTINCT 
	t1.ent_id,t2.ent_id,t3.ent_id
	FROM good_users AS t1 
	LEFT JOIN root_users AS t2 ON t1.root_user_id = t2.ent_id 
	LEFT JOIN app_pools AS t3 ON t2.pool_id = t3.pool_id
	WHERE (t1.ent_id = '%v' AND t3.app_id = '%v') AND t1.deleted_at = 0 AND t2.deleted_at = 0 AND t3.deleted_at = 0;`,
		h.GoodUserID.String(),
		h.AppID.String(),
	)

	h.checkAppAuthSQL = sql
	return nil
}

func (h *Handler) checkAppAuth(ctx context.Context) (bool, error) {
	exist := false
	var err error
	existH := &existHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		err := existH.existJoinCoinAndAppPool()
		if err != nil {
			return wlog.WrapError(err)
		}

		rows, err := cli.QueryContext(ctx, existH.checkAppAuthSQL)
		if err != nil {
			return wlog.WrapError(err)
		}
		defer rows.Close()

		if err := rows.Err(); err != nil {
			return wlog.WrapError(err)
		}

		exist = false
		for rows.Next() {
			exist = true
			break
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistOrderUser(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			OrderUser.
			Query().
			Where(
				orderuserent.EntID(*h.EntID),
				orderuserent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistOrderUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
