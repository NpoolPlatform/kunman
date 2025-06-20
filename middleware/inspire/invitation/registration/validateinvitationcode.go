package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	invitationcodemwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/invitationcode"
	invitationcode1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/invitationcode"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) validateInvitationCode(ctx context.Context) error {
	h1, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithConds(&invitationcodemwpb.Conds{
			AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.InviterID.String()},
			Disabled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		}),
		invitationcode1.WithLimit(0),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	exist, err := h1.ExistInvitationCodeConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invatationcode not exist")
	}

	return nil
}
