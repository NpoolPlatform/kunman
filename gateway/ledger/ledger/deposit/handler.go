package deposit

import (
	"context"
	"fmt"

	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	AppID        *string
	UserID       *string
	CoinTypeID   *string
	TargetAppID  *string
	TargetUserID *string
	Amount       *string
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithAppID(appID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			return fmt.Errorf("invalid app id")
		}
		_, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		handler, err := appmw.NewHandler(
			ctx,
			appmw.WithEntID(appID, true),
		)
		if err != nil {
			return err
		}

		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}

		h.AppID = appID
		return nil
	}
}

func WithUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid user id")
			}
			return nil
		}
		_, err := uuid.Parse(*userID)
		if err != nil {
			return err
		}
		h.UserID = userID
		return nil
	}
}

func WithTargetAppID(appID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			if must {
				return fmt.Errorf("invalid target app id")
			}
			return nil
		}
		_, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		handler, err := appmw.NewHandler(
			ctx,
			appmw.WithEntID(appID, true),
		)
		if err != nil {
			return err
		}

		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}

		h.TargetAppID = appID
		return nil
	}
}

func WithTargetUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid target user id")
			}
			return nil
		}
		h.TargetUserID = userID
		return nil
	}
}

func WithCoinTypeID(coinTypeID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if coinTypeID == nil {
			if must {
				return fmt.Errorf("invalid coin type id")
			}
			return nil
		}
		_, err := uuid.Parse(*coinTypeID)
		if err != nil {
			return err
		}
		h.CoinTypeID = coinTypeID
		return nil
	}
}

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}

		h.Amount = amount
		return nil
	}
}
