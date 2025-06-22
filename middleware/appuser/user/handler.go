//nolint:dupl
package user

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user"
	oauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/oauth/oauththirdparty"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                 *uint32
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	PhoneNO            *string
	CountryCode        *string
	EmailAddress       *string
	ImportFromAppID    *uuid.UUID
	Username           *string
	AddressFields      []string
	Gender             *string
	PostalCode         *string
	Age                *uint32
	Birthday           *uint32
	Avatar             *string
	Organization       *string
	FirstName          *string
	LastName           *string
	IDNumber           *string
	SigninVerifyType   *basetypes.SignMethod
	GoogleAuthVerified *bool
	PasswordHash       *string
	GoogleSecret       *string
	ThirdPartyID       *uuid.UUID
	ThirdPartyUserID   *string
	ThirdPartyUsername *string
	ThirdPartyAvatar   *string
	Banned             *bool
	BanMessage         *string
	RoleIDs            []uuid.UUID
	Kol                *bool
	KolConfirmed       *bool
	SelectedLangID     *uuid.UUID
	Account            *string
	AccountType        *basetypes.SignMethod
	Conds              *usercrud.Conds
	Offset             int32
	Limit              int32
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

// Here ID is UserID
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
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
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

func validateEmailAddress(emailAddress string) error {
	if _, err := mail.ParseAddress(emailAddress); err != nil {
		return err
	}
	return nil
}

func validatePhoneNO(phoneNO string) error {
	re := regexp.MustCompile(
		`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[` +
			`\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?)` +
			`{0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)` +
			`[\-\.\ \\\/]?(\d+))?$`,
	)
	if !re.MatchString(phoneNO) {
		return fmt.Errorf("invalid phone no")
	}

	return nil
}

func WithPhoneNO(phoneNO *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if phoneNO == nil {
			if must {
				return fmt.Errorf("invalid phoneno")
			}
			return nil
		}
		if err := validatePhoneNO(*phoneNO); err != nil {
			return err
		}
		accountType := basetypes.SignMethod_Mobile
		h.Account = phoneNO
		h.AccountType = &accountType
		h.PhoneNO = phoneNO
		return nil
	}
}

func WithCountryCode(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid countrycode")
			}
			return nil
		}
		h.CountryCode = s
		return nil
	}
}

func WithEmailAddress(emailAddress *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if emailAddress == nil {
			if must {
				return fmt.Errorf("invalid emailaddress")
			}
			return nil
		}
		if err := validateEmailAddress(*emailAddress); err != nil {
			return err
		}

		accountType := basetypes.SignMethod_Email
		h.Account = emailAddress
		h.AccountType = &accountType
		h.EmailAddress = emailAddress
		return nil
	}
}

func WithImportFromAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid importfromappid")
			}
			return nil
		}
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
		h.ImportFromAppID = &_id
		return nil
	}
}

func WithUsername(username *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if username == nil {
			if must {
				return fmt.Errorf("invalid username")
			}
			return nil
		}
		re := regexp.MustCompile("^[a-zA-Z0-9\u3040-\u31ff][[a-zA-Z0-9_\\-\\.\u3040-\u31ff]{3,32}$") //nolint
		if !re.MatchString(*username) {
			return fmt.Errorf("invalid username")
		}
		h.Username = username
		return nil
	}
}

func WithAddressFields(addressFields []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AddressFields = addressFields
		return nil
	}
}

func WithGender(gender *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gender == nil {
			if must {
				return fmt.Errorf("invalid gender")
			}
			return nil
		}
		if *gender == "" {
			return fmt.Errorf("invalid gender")
		}
		h.Gender = gender
		return nil
	}
}

func WithPostalCode(postalCode *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if postalCode == nil {
			if must {
				return fmt.Errorf("invalid postalcode")
			}
			return nil
		}
		if *postalCode == "" {
			return fmt.Errorf("invalid postalCode")
		}
		h.PostalCode = postalCode
		return nil
	}
}

func WithAge(age *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Age = age
		return nil
	}
}

func WithBirthday(birthday *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Birthday = birthday
		return nil
	}
}

func WithAvatar(avatar *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if avatar == nil {
			if must {
				return fmt.Errorf("invalid avatar")
			}
			return nil
		}
		if *avatar == "" {
			return fmt.Errorf("invalid avatar")
		}
		h.Avatar = avatar
		return nil
	}
}

func WithOrganization(organization *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if organization == nil {
			if must {
				return fmt.Errorf("invalid organization")
			}
			return nil
		}
		if *organization == "" {
			return fmt.Errorf("invalid organization")
		}
		h.Organization = organization
		return nil
	}
}

func WithFirstName(firstName *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if firstName == nil {
			if must {
				return fmt.Errorf("invalid firstname")
			}
			return nil
		}
		if *firstName == "" {
			return fmt.Errorf("invalid firstname")
		}
		h.FirstName = firstName
		return nil
	}
}

func WithLastName(lastName *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lastName == nil {
			if must {
				return fmt.Errorf("invalid lastname")
			}
			return nil
		}
		if *lastName == "" {
			return fmt.Errorf("invalid lastname")
		}
		h.LastName = lastName
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
		if *idNumber == "" {
			return fmt.Errorf("invalid idnumber")
		}
		h.IDNumber = idNumber
		return nil
	}
}

func WithSigninVerifyType(verifyType *basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if verifyType == nil {
			if must {
				return fmt.Errorf("invalid signinverifytype")
			}
			return nil
		}
		switch *verifyType {
		case basetypes.SignMethod_Email:
		case basetypes.SignMethod_Mobile:
		case basetypes.SignMethod_Google:
		case basetypes.SignMethod_Reset:
		default:
			return fmt.Errorf("invalid sign verify type")
		}
		h.SigninVerifyType = verifyType
		return nil
	}
}

func WithGoogleAuthVerified(verified *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoogleAuthVerified = verified
		return nil
	}
}

func WithPasswordHash(passwordHash *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if passwordHash == nil {
			if must {
				return fmt.Errorf("invalid passwordhash")
			}
			return nil
		}
		if *passwordHash == "" {
			return fmt.Errorf("invalid passwordhash")
		}
		h.PasswordHash = passwordHash
		return nil
	}
}

func WithGoogleSecret(secret *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if secret == nil {
			if must {
				return fmt.Errorf("invalid secret")
			}
			return nil
		}
		if *secret == "" {
			return fmt.Errorf("invalid google secret")
		}
		h.GoogleSecret = secret
		return nil
	}
}

func WithThirdPartyID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid thirdpartyid")
			}
			return nil
		}
		handler, err := oauththirdparty.NewHandler(
			ctx,
			oauththirdparty.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		thirdParty, err := handler.GetOAuthThirdParty(ctx)
		if err != nil {
			return err
		}
		if thirdParty == nil {
			return fmt.Errorf("invalid oauththirdparty")
		}
		accountType := basetypes.SignMethod(basetypes.SignMethod_value[thirdParty.ClientNameStr])
		h.AccountType = &accountType

		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		h.ThirdPartyID = &_id
		return nil
	}
}

func WithThirdPartyUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid thirdpartyuserid")
			}
			return nil
		}
		if *userID == "" {
			return fmt.Errorf("invalid thirdpartyuserid")
		}
		h.Account = userID
		h.ThirdPartyUserID = userID
		return nil
	}
}

func WithThirdPartyUsername(username *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if username == nil {
			if must {
				return fmt.Errorf("invalid thirdpartyusername")
			}
			return nil
		}
		if *username == "" {
			return fmt.Errorf("invalid thirdpartyusername")
		}
		h.ThirdPartyUsername = username
		return nil
	}
}

func WithThirdPartyAvatar(avatar *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if avatar == nil {
			if must {
				return fmt.Errorf("invalid thirdpartyavatar")
			}
			return nil
		}
		if *avatar == "" {
			return fmt.Errorf("invalid avatar")
		}
		h.ThirdPartyAvatar = avatar
		return nil
	}
}

func WithBanned(banned *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Banned = banned
		return nil
	}
}

func WithBanMessage(message *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			if must {
				return fmt.Errorf("invalid banmessage")
			}
			return nil
		}
		if *message == "" {
			return fmt.Errorf("invalid message")
		}
		h.BanMessage = message
		return nil
	}
}

func WithRoleIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(ids) == 0 {
			return nil
		}
		_ids := []uuid.UUID{}
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			_ids = append(_ids, _id)
		}
		h.RoleIDs = _ids
		return nil
	}
}

func WithKol(kol *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Kol = kol
		return nil
	}
}

func WithKolConfirmed(confirmed *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.KolConfirmed = confirmed
		return nil
	}
}

func WithSelectedLangID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid selectedlangid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.SelectedLangID = &_id
		return nil
	}
}

func WithAccount(account *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == nil {
			if must {
				return fmt.Errorf("invalid account")
			}
		}
		if *account == "" {
			return fmt.Errorf("invalid account")
		}

		var accountType basetypes.SignMethod
		if err := validatePhoneNO(*account); err == nil {
			h.PhoneNO = account
			accountType = basetypes.SignMethod_Mobile
		} else if err := validateEmailAddress(*account); err == nil {
			accountType = basetypes.SignMethod_Email
			h.EmailAddress = account
		} else {
			return fmt.Errorf("invalid account")
		}

		if h.AccountType != nil && accountType != *h.AccountType {
			return fmt.Errorf("invalid accounttype")
		}

		h.AccountType = &accountType
		h.Account = account
		return nil
	}
}

func WithAccountType(accountType *basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if accountType == nil {
			if must {
				return fmt.Errorf("invalid accounttype")
			}
			return nil
		}
		if h.AccountType != nil && *accountType != *h.AccountType {
			return fmt.Errorf("invalid accounttype")
		}
		switch *accountType {
		case basetypes.SignMethod_Mobile:
		case basetypes.SignMethod_Email:
		default:
			return fmt.Errorf("invalid accounttype")
		}
		h.AccountType = accountType
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &usercrud.Conds{}
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
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.PhoneNO != nil {
			h.Conds.PhoneNO = &cruder.Cond{
				Op:  conds.GetPhoneNO().GetOp(),
				Val: conds.GetPhoneNO().GetValue(),
			}
		}
		if conds.EmailAddress != nil {
			h.Conds.EmailAddress = &cruder.Cond{
				Op:  conds.GetEmailAddress().GetOp(),
				Val: conds.GetEmailAddress().GetValue(),
			}
		}
		if conds.ImportFromApp != nil {
			id, err := uuid.Parse(conds.GetImportFromApp().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ImportFromApp = &cruder.Cond{
				Op:  conds.GetImportFromApp().GetOp(),
				Val: id,
			}
		}
		if len(conds.GetEntIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if conds.ThirdPartyID != nil {
			id, err := uuid.Parse(conds.GetThirdPartyID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ThirdPartyID = &cruder.Cond{Op: conds.GetThirdPartyID().GetOp(), Val: id}
		}
		if conds.ThirdPartyUserID != nil {
			h.Conds.ThirdPartyUserID = &cruder.Cond{
				Op:  conds.GetThirdPartyUserID().GetOp(),
				Val: conds.GetThirdPartyUserID().GetValue(),
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
