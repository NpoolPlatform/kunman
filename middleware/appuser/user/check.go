package user

import (
	"context"
	"fmt"

	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) checkAccountExist(ctx context.Context) error {
	if h.PhoneNO == nil && h.EmailAddress == nil && (h.ThirdPartyID == nil || h.ThirdPartyUserID == nil) {
		return nil
	}

	conds := &usercrud.Conds{
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
	}
	if h.EmailAddress != nil {
		conds.EmailAddress = &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress}
	}
	if h.PhoneNO != nil {
		conds.PhoneNO = &cruder.Cond{Op: cruder.EQ, Val: *h.PhoneNO}
	}
	if h.ThirdPartyID != nil {
		conds.ThirdPartyID = &cruder.Cond{Op: cruder.EQ, Val: *h.ThirdPartyID}
	}
	if h.ThirdPartyUserID != nil {
		conds.ThirdPartyUserID = &cruder.Cond{Op: cruder.EQ, Val: *h.ThirdPartyUserID}
	}

	h1, err := NewHandler(ctx)
	if err != nil {
		return err
	}
	h1.Conds = conds

	exist, err := h1.ExistUserConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("already exists")
	}

	return nil
}
