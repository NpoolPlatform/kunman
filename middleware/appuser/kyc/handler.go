package kyc

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/kyc"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	kyccrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/kyc"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	DocumentType *basetypes.KycDocumentType
	IDNumber     *string
	FrontImg     *string
	BackImg      *string
	SelfieImg    *string
	EntityType   *basetypes.KycEntityType
	ReviewID     *uuid.UUID
	State        *basetypes.KycState
	Conds        *kyccrud.Conds
	Offset       int32
	Limit        int32
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
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := app.NewHandler(
			ctx,
			app.WithEntID(id, true),
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithDocumentType(docType *basetypes.KycDocumentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if docType == nil {
			if must {
				return fmt.Errorf("invalid documenttype")
			}
			return nil
		}
		switch *docType {
		case basetypes.KycDocumentType_IDCard:
		case basetypes.KycDocumentType_DriverLicense:
		case basetypes.KycDocumentType_Passport:
		default:
			return fmt.Errorf("invalid document type")
		}
		h.DocumentType = docType
		return nil
	}
}

func WithIDNumber(idNumber *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if idNumber == nil {
			if must {
				return fmt.Errorf("invalid idnumber")
			}
			return nil
		}
		const leastIDNumberLen = 8
		if len(*idNumber) < leastIDNumberLen {
			return fmt.Errorf("invalid id number")
		}
		h.IDNumber = idNumber
		return nil
	}
}

func WithFrontImg(img *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.FrontImg = img
		return nil
	}
}

func WithBackImg(img *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.BackImg = img
		return nil
	}
}

func WithSelfieImg(img *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SelfieImg = img
		return nil
	}
}

func WithEntityType(entType *basetypes.KycEntityType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if entType == nil {
			if must {
				return fmt.Errorf("invalid entitytype")
			}
			return nil
		}
		switch *entType {
		case basetypes.KycEntityType_Individual:
		case basetypes.KycEntityType_Organization:
		default:
			return fmt.Errorf("invalid entity type")
		}
		h.EntityType = entType
		return nil
	}
}

func WithReviewID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid reviewid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ReviewID = &_id
		return nil
	}
}

func WithState(state *basetypes.KycState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case basetypes.KycState_Approved:
		case basetypes.KycState_Reviewing:
		case basetypes.KycState_Rejected:
		default:
			return fmt.Errorf("invalid state")
		}
		h.State = state
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &kyccrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op: conds.GetAppID().GetOp(), Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op: conds.GetUserID().GetOp(), Val: id,
			}
		}
		if conds.DocumentType != nil {
			docType := conds.GetDocumentType().GetValue()
			h.Conds.DocumentType = &cruder.Cond{
				Op:  conds.GetDocumentType().GetOp(),
				Val: basetypes.KycDocumentType(docType),
			}
		}
		if conds.IDNumber != nil {
			h.Conds.IDNumber = &cruder.Cond{
				Op:  conds.GetIDNumber().GetOp(),
				Val: conds.GetIDNumber().GetValue(),
			}
		}
		if conds.EntityType != nil {
			entType := conds.GetEntityType().GetValue()
			h.Conds.EntityType = &cruder.Cond{
				Op:  conds.GetEntityType().GetOp(),
				Val: basetypes.KycEntityType(entType),
			}
		}
		if conds.ReviewID != nil {
			id, err := uuid.Parse(conds.GetReviewID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ReviewID = &cruder.Cond{
				Op:  conds.GetReviewID().GetOp(),
				Val: id,
			}
		}
		if conds.State != nil {
			state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{
				Op:  conds.GetState().GetOp(),
				Val: basetypes.KycState(state),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
