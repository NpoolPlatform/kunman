package delegatedstaking

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDelegatedStaking(ctx context.Context) error {
	handler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithID(h.ID, true),
		appdelegatedstakingmw.WithEntID(h.EntID, true),
		appdelegatedstakingmw.WithAppID(h.AppID, true),
		appdelegatedstakingmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistDelegatedStaking(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appdelegatedstaking")
	}
	return nil
}
