//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"strings"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"

	"github.com/google/uuid"
)

func CreateSuperiorProcedure(ctx context.Context) error {
	const procedure = `
		DROP PROCEDURE IF EXISTS get_superiores;
		SET SESSION GROUP_CONCAT_MAX_LEN = 102400;
		CREATE PROCEDURE get_superiores (IN invitees TEXT)
		BEGIN
		  DECLARE superiores TEXT;
		  DECLARE my_invitees TEXT;
		  SET superiores = 'N/A';
		  SET my_invitees = invitees;
		  WHILE my_invitees is not null DO
		    if superiores = 'N/A' THEN
			  SET superiores = my_invitees;
			else
			  SET superiores = CONCAT(superiores, ',', my_invitees);
			END if;
		    SELECT GROUP_CONCAT(DISTINCT inviter_id) INTO my_invitees FROM registrations WHERE FIND_IN_SET(invitee_id, my_invitees) AND deleted_at=0;
		  END WHILE;
		  SELECT superiores;
		END
	`
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.ExecContext(_ctx, procedure)
		return err
	})
}

func (h *queryHandler) getInviteeIDsWithClient(ctx context.Context, cli *ent.Client) error {
	if h.Conds.InviteeIDs == nil {
		return wlog.Errorf("invalid inviteeids")
	}

	inviteeIDs, ok := h.Conds.InviteeIDs.Val.([]uuid.UUID)
	if !ok {
		return wlog.Errorf("invalid inviteeids")
	}
	_inviteeIDs := ""
	for _, id := range inviteeIDs {
		if _inviteeIDs != "" {
			_inviteeIDs = fmt.Sprintf("%v,", _inviteeIDs)
		}
		_inviteeIDs = fmt.Sprintf("%v%v", _inviteeIDs, id)
	}

	rows, err := cli.QueryContext(
		ctx,
		fmt.Sprintf("CALL get_superiores(\"%v\")", _inviteeIDs),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	defer rows.Close()

	subordinates := ""
	for rows.Next() {
		if err := rows.Scan(&subordinates); err != nil {
			return wlog.WrapError(err)
		}
	}

	__inviteeIDs := strings.Split(subordinates, ",") //nolint
	for _, id := range __inviteeIDs {
		_id, err := uuid.Parse(id)
		if err != nil {
			return wlog.WrapError(err)
		}
		inviteeIDs = append(inviteeIDs, _id)
	}

	h.Conds.InviteeIDs.Val = inviteeIDs
	return nil
}

func (h *queryHandler) getInviteeIDs(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h.getInviteeIDsWithClient(_ctx, cli)
	})
}

func (h *Handler) GetSuperiores(ctx context.Context) ([]*npool.Registration, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	if err := handler.getInviteeIDs(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	infos, total, err := h.GetRegistrations(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return infos, total, nil
}

func (h *Handler) GetSuperioresWithClient(ctx context.Context, cli *ent.Client) ([]*npool.Registration, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	if err := handler.getInviteeIDsWithClient(ctx, cli); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	infos, total, err := h.GetRegistrationsWithClient(ctx, cli)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return infos, total, nil
}
