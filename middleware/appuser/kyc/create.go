package kyc

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/kyc"
	kyccrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/kyc"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	user1 "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateKyc(ctx context.Context) (*npool.Kyc, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	// TODO: deduplicate

	userID := h.UserID.String()
	appID := h.AppID.String()

	h1, err := user1.NewHandler(
		ctx,
		user1.WithAppID(&appID, true),
		user1.WithEntID(&userID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err := h1.ExistUser(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := kyccrud.SetQueryConds(
			cli.Kyc.Query(),
			&kyccrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			return fmt.Errorf("kyc exist")
		}

		if _, err := kyccrud.CreateSet(
			cli.Kyc.Create(),
			&kyccrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				UserID:       h.UserID,
				DocumentType: h.DocumentType,
				IDNumber:     h.IDNumber,
				FrontImg:     h.FrontImg,
				BackImg:      h.BackImg,
				SelfieImg:    h.SelfieImg,
				EntityType:   h.EntityType,
				ReviewID:     h.ReviewID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetKyc(ctx)
}
