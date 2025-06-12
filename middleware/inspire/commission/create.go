package commission

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
	commissioncrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/commission"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommission "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/commission"

	"github.com/google/uuid"
)

func (h *Handler) CreateCommission(ctx context.Context) (*npool.Commission, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			Commission.
			Update().
			Where(
				entcommission.AppID(*h.AppID),
				entcommission.UserID(*h.UserID),
				entcommission.AppGoodID(*h.AppGoodID),
				entcommission.SettleType(h.SettleType.String()),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}

		if _, err := commissioncrud.CreateSet(
			tx.Commission.Create(),
			&commissioncrud.Req{
				EntID:            h.EntID,
				AppID:            h.AppID,
				UserID:           h.UserID,
				GoodID:           h.GoodID,
				AppGoodID:        h.AppGoodID,
				SettleType:       h.SettleType,
				SettleMode:       h.SettleMode,
				SettleAmountType: h.SettleAmountType,
				SettleInterval:   h.SettleInterval,
				AmountOrPercent:  h.AmountOrPercent,
				StartAt:          h.StartAt,
				Threshold:        h.Threshold,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetCommission(ctx)
}
