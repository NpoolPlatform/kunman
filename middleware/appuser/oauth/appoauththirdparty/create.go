package appoauththirdparty

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/appuser/aes"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"

	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/appoauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	redis2 "github.com/NpoolPlatform/kunman/framework/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppOAuthThirdParty, h.AppID, h.ThirdPartyID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	oauthHandler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			ThirdPartyID: &basetypes.StringVal{Op: cruder.EQ, Value: h.ThirdPartyID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := oauthHandler.ExistOAuthThirdPartyConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("oauththirdparty exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	salt, err := aes.NewAesKey(aes.AES256)
	if err != nil {
		return nil, fmt.Errorf("get salt failed")
	}
	clientSecret, err := aes.AesEncrypt([]byte(salt), []byte(*h.ClientSecret))
	if err != nil {
		return nil, fmt.Errorf("encrypt clientSecret failed")
	}
	clientSecretStr := hex.EncodeToString(clientSecret)
	h.ClientSecret = &clientSecretStr

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := appoauththirdpartycrud.CreateSet(
			tx.AppOAuthThirdParty.Create(),
			&appoauththirdpartycrud.Req{
				EntID:        h.EntID,
				AppID:        &h.AppID,
				ThirdPartyID: h.ThirdPartyID,
				ClientID:     h.ClientID,
				ClientSecret: h.ClientSecret,
				CallbackURL:  h.CallbackURL,
				Salt:         &salt,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
