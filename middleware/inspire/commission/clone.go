package commission

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommission "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/commission"

	"github.com/shopspring/decimal"
)

func (h *Handler) CloneCommissions(ctx context.Context) error {
	if *h.FromAppGoodID == *h.ToAppGoodID {
		return nil
	}

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Commission.
			Query().
			Where(
				entcommission.AppID(*h.AppID),
				entcommission.GoodID(*h.FromGoodID),
				entcommission.AppGoodID(*h.FromAppGoodID),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			).
			All(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(infos) == 0 {
			return wlog.Errorf("commission not found")
		}

		percent := decimal.NewFromInt(1)
		if h.ScalePercent != nil {
			percent = h.ScalePercent.Div(decimal.NewFromInt(100)) //nolint
		}

		cs := []*ent.CommissionCreate{}
		for _, info := range infos {
			info1, err := cli.
				Commission.
				Query().
				Where(
					entcommission.AppID(*h.AppID),
					entcommission.UserID(info.UserID),
					entcommission.GoodID(*h.ToGoodID),
					entcommission.AppGoodID(*h.ToAppGoodID),
					entcommission.SettleType(info.SettleType),
					entcommission.EndAt(0),
					entcommission.DeletedAt(0),
				).
				Only(_ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return wlog.WrapError(err)
				}
			}
			if info1 != nil {
				if _, err := cli.
					Commission.
					UpdateOneID(info1.ID).
					SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
					SetSettleType(info.SettleType).
					SetSettleMode(info.SettleMode).
					SetSettleAmountType(info.SettleAmountType).
					SetSettleInterval(info.SettleInterval).
					SetThreshold(info.Threshold).
					Save(_ctx); err != nil {
					return wlog.WrapError(err)
				}
				continue
			}

			c := cli.
				Commission.
				Create().
				SetAppID(info.AppID).
				SetUserID(info.UserID).
				SetGoodID(*h.ToGoodID).
				SetAppGoodID(*h.ToAppGoodID).
				SetSettleType(info.SettleType).
				SetSettleMode(info.SettleMode).
				SetSettleAmountType(info.SettleAmountType).
				SetSettleInterval(info.SettleInterval).
				SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
				SetStartAt(now).
				SetThreshold(info.Threshold)
			cs = append(cs, c)
		}
		if _, err := cli.
			Commission.
			CreateBulk(cs...).
			Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
