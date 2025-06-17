package devicetype

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	manufacturercommon "github.com/NpoolPlatform/kunman/gateway/good/device/manufacturer/common"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	Type  *string
	manufacturercommon.ManufacturerCheckHandler
	PowerConsumption *uint32
	ShipmentAt       *uint32
	Offset           int32
	Limit            int32
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithType(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid type")
			}
			return nil
		}
		const leastTypeLen = 3
		if len(*s) < leastTypeLen {
			return fmt.Errorf("invalid type")
		}
		h.Type = s
		return nil
	}
}

func WithManufacturerID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid manufacturerid")
			}
			return nil
		}
		if err := h.CheckManufacturerWithManufacturerID(ctx, *s); err != nil {
			return err
		}
		h.ManufacturerID = s
		return nil
	}
}

func WithPowerConsumption(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid powercunsumption")
			}
			return nil
		}
		h.PowerConsumption = n
		return nil
	}
}

func WithShipmentAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid powercunsumption")
			}
			return nil
		}
		h.ShipmentAt = n
		return nil
	}
}

func WithOffset(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = n
		return nil
	}
}

func WithLimit(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == 0 {
			n = constant.DefaultRowLimit
		}
		h.Limit = n
		return nil
	}
}
