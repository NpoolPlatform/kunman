package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	registrationcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/registration"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type sortHandler struct {
	*Handler
	inviters   []*npool.Registration
	inviterIDs []string
}

func (h *sortHandler) getSuperioresWithClient(ctx context.Context, cli *ent.Client) error {
	h.Limit = constant.DefaultRowLimit
	h.Offset = 0
	h.Conds = &registrationcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		InviteeIDs: &cruder.Cond{Op: cruder.IN, Val: []uuid.UUID{*h.InviteeID}},
	}

	inviters := []*npool.Registration{}
	for {
		_inviters, _, err := h.GetSuperioresWithClient(ctx, cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(_inviters) == 0 {
			break
		}
		inviters = append(inviters, _inviters...)
		h.Offset += h.Limit
	}

	h.inviters = inviters

	return nil
}

func (h *sortHandler) getSuperiores(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h.getSuperioresWithClient(_ctx, cli)
	})
}

func (h *sortHandler) sortInviters() error {
	inviters := h.inviters

	inviteeMap := map[string]struct{}{}
	for _, inviter := range inviters {
		inviteeMap[inviter.InviteeID] = struct{}{}
	}

	inviterCount := len(inviters)
	_inviters := []*npool.Registration{}

	for i, inviter := range inviters {
		_, ok := inviteeMap[inviter.InviterID]
		if !ok {
			_inviters = append(_inviters, inviter)
			inviters = append(inviters[0:i], inviters[i+1:]...)
			break
		}
	}

	if inviterCount == 0 {
		_inviters = append(_inviters, &npool.Registration{
			AppID:     h.AppID.String(),
			InviterID: uuid.Nil.String(),
			InviteeID: h.InviteeID.String(),
		})
	}

	if len(_inviters) == 0 {
		return wlog.Errorf("invalid top inviter")
	}

	for {
		if inviterCount == 0 || len(inviters) == 0 {
			break
		}

		if len(inviters) == 1 {
			if _inviters[len(_inviters)-1].InviteeID != inviters[0].InviterID {
				return wlog.Errorf("mismatch registration")
			}
			_inviters = append(_inviters, inviters[0])
			break
		}

		for i, inviter := range inviters {
			if _inviters[len(_inviters)-1].InviteeID == inviter.InviterID {
				_inviters = append(_inviters, inviter)
				inviters = append(inviters[0:i], inviters[i+1:]...)
				break
			}
		}
	}

	inviterIDs := []string{h.InviteeID.String()}
	if inviterCount > 0 {
		inviterIDs = []string{_inviters[0].InviterID}
		for _, inviter := range _inviters {
			inviterIDs = append(inviterIDs, inviter.InviteeID)
		}
	}

	h.inviters = _inviters
	h.inviterIDs = inviterIDs

	return nil
}

func (h *Handler) GetSortedInviters(ctx context.Context) ([]*npool.Registration, []string, error) {
	handler := &sortHandler{
		Handler: h,
	}

	if err := handler.getSuperiores(ctx); err != nil {
		return nil, nil, wlog.WrapError(err)
	}
	if err := handler.sortInviters(); err != nil {
		return nil, nil, wlog.WrapError(err)
	}

	return handler.inviters, handler.inviterIDs, nil
}

func (h *Handler) GetSortedInvitersWithClient(ctx context.Context, cli *ent.Client) ([]*npool.Registration, []string, error) {
	handler := &sortHandler{
		Handler: h,
	}

	if err := handler.getSuperioresWithClient(ctx, cli); err != nil {
		return nil, nil, wlog.WrapError(err)
	}
	if err := handler.sortInviters(); err != nil {
		return nil, nil, wlog.WrapError(err)
	}

	return handler.inviters, handler.inviterIDs, nil
}
