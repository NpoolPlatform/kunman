package invitationcode

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/invitationcode"
	invitationcodecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	codegenerator "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/invitationcode/generator"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/google/uuid"
)

func (h *Handler) CreateInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	h.Conds = &invitationcodecrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	exist, err := h.ExistInvitationCodeConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("already exists")
	}

	var code string
	for {
		code, err = codegenerator.Generate()
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		h.Conds = &invitationcodecrud.Conds{
			AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			InvitationCode: &cruder.Cond{Op: cruder.EQ, Val: code},
		}
		exist, err := h.ExistInvitationCodeConds(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if exist {
			continue
		}
		break
	}

	exist, err = h.ExistInvitationCodeConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("already exists")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.CreateSet(
			cli.InvitationCode.Create(),
			&invitationcodecrud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				InvitationCode: &code,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetInvitationCode(ctx)
}
