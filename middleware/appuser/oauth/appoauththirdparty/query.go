package appoauththirdparty

import (
	"context"
	"encoding/hex"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/appuser/aes"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"
	entappoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/appoauththirdparty"
	entoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/oauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/appoauththirdparty"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppOAuthThirdPartySelect
	stmCount  *ent.AppOAuthThirdPartySelect
	infos     []*npool.OAuthThirdParty
	total     uint32
}

func (h *queryHandler) selectOAuthThirdParty(stm *ent.AppOAuthThirdPartyQuery) *ent.AppOAuthThirdPartySelect {
	return stm.Select(entappoauththirdparty.FieldID)
}

func (h *queryHandler) queryOAuthThirdParty(cli *ent.Client) {
	stm := cli.AppOAuthThirdParty.
		Query().
		Where(entappoauththirdparty.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappoauththirdparty.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappoauththirdparty.EntID(*h.EntID))
	}
	h.stmSelect = h.selectOAuthThirdParty(stm)
}

func (h *queryHandler) queryOAuthThirdParties(cli *ent.Client) (*ent.AppOAuthThirdPartySelect, error) {
	stm, err := appoauththirdpartycrud.SetQueryConds(cli.AppOAuthThirdParty.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectOAuthThirdParty(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappoauththirdparty.Table)
	s.AppendSelect(
		t.C(entappoauththirdparty.FieldEntID),
		t.C(entappoauththirdparty.FieldAppID),
		t.C(entappoauththirdparty.FieldThirdPartyID),
		t.C(entappoauththirdparty.FieldClientID),
		t.C(entappoauththirdparty.FieldClientSecret),
		t.C(entappoauththirdparty.FieldCallbackURL),
		t.C(entappoauththirdparty.FieldSalt),
		t.C(entappoauththirdparty.FieldCreatedAt),
		t.C(entappoauththirdparty.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinOAuthThirdParty(s *sql.Selector) error {
	t := sql.Table(entoauththirdparty.Table)
	s.Join(t).
		On(
			s.C(entappoauththirdparty.FieldThirdPartyID),
			t.C(entoauththirdparty.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entoauththirdparty.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.ClientName != nil {
		clientName, ok := h.Conds.ClientName.Val.(basetypes.SignMethod)
		if !ok {
			return fmt.Errorf("invalid oauth clientName")
		}
		s.Where(
			sql.EQ(t.C(entoauththirdparty.FieldClientName), clientName.String()),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entoauththirdparty.FieldClientName), "client_name"),
		sql.As(t.C(entoauththirdparty.FieldClientTag), "client_tag"),
		sql.As(t.C(entoauththirdparty.FieldClientLogoURL), "client_logo_url"),
		sql.As(t.C(entoauththirdparty.FieldClientOauthURL), "client_oauth_url"),
		sql.As(t.C(entoauththirdparty.FieldResponseType), "response_type"),
		sql.As(t.C(entoauththirdparty.FieldScope), "scope"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinOAuthThirdParty(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinOAuthThirdParty(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() error {
	isDecryptSecret := false
	if h.Conds != nil && h.Conds.DecryptSecret != nil {
		decryptSecret, ok := h.Conds.DecryptSecret.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid oauth decryptsecret")
		}
		isDecryptSecret = decryptSecret
	}

	for _, info := range h.infos {
		info.ClientName = basetypes.SignMethod(basetypes.SignMethod_value[info.ClientNameStr])
		if isDecryptSecret {
			ClientSecretBytes, err := hex.DecodeString(info.ClientSecret)
			if err != nil {
				return fmt.Errorf("secret err")
			}
			clientSecret, err := aes.AesDecrypt([]byte(info.Salt), ClientSecretBytes)
			if err != nil {
				return err
			}
			info.ClientSecret = string(clientSecret)
		}
		if info.Salt != "" {
			info.Salt = ""
		}
	}

	return nil
}

func (h *Handler) GetOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOAuthThirdParty(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	if err := handler.formalize(); err != nil {
		return nil, err
	}

	return handler.infos[0], nil
}

func (h *Handler) GetOAuthThirdParties(ctx context.Context) ([]*npool.OAuthThirdParty, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOAuthThirdParties(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryOAuthThirdParties(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entappoauththirdparty.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	if err := handler.formalize(); err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
