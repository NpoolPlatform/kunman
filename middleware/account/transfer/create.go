package transfer

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/transfer"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	transfercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/transfer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateTransfer(ctx context.Context) (*npool.Transfer, error) {
	// TODO: deduplicate

	handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
			TargetUserID: &basetypes.StringVal{Op: cruder.EQ, Value: h.TargetUserID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistTransferConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("transfer exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := transfercrud.CreateSet(
			tx.Transfer.Create(),
			&transfercrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				UserID:       h.UserID,
				TargetUserID: h.TargetUserID,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTransfer(ctx)
}
