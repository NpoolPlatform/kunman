package appoauththirdparty

import (
	"context"
	"encoding/hex"
	"fmt"

	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appoauththirdparty"
	"github.com/NpoolPlatform/kunman/pkg/aes"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/appoauththirdparty"
)

func (h *Handler) UpdateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	info, err := h.GetOAuthThirdParty(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		oauthThirdParty, err := tx.AppOAuthThirdParty.
			Query().
			Where(
				entappoauththirdparty.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if oauthThirdParty == nil {
			return fmt.Errorf("invalid oauththirdparty")
		}
		if h.ClientSecret != nil && *h.ClientSecret != "" {
			salt, err := aes.NewAesKey(aes.AES256)
			if err != nil {
				return fmt.Errorf("get salt failed")
			}
			h.Salt = &salt
			clientSecret, err := aes.AesEncrypt([]byte(salt), []byte(*h.ClientSecret))
			if err != nil {
				return fmt.Errorf("encrypt clientSecret failed")
			}
			clientSecretStr := hex.EncodeToString(clientSecret)
			h.ClientSecret = &clientSecretStr
			h.ClientSecret = &clientSecretStr
		}

		if _, err := appoauththirdpartycrud.UpdateSet(
			oauthThirdParty.Update(),
			&appoauththirdpartycrud.Req{
				ClientID:     h.ClientID,
				ClientSecret: h.ClientSecret,
				CallbackURL:  h.CallbackURL,
				Salt:         h.Salt,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
