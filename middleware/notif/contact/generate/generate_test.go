package generate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	contactmw "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	contact "github.com/NpoolPlatform/kunman/middleware/notif/contact"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/notif/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = contactmw.Contact{
		AppID:          uuid.NewString(),
		Account:        "vagrant1@163.com",
		AccountType:    basetypes.SignMethod_Email,
		AccountTypeStr: basetypes.SignMethod_Email.String(),
		UsedFor:        basetypes.UsedFor_Contact,
		UsedForStr:     basetypes.UsedFor_Contact.String(),
		Sender:         "vagrant3@163.com",
	}

	subject    = uuid.NewString()
	body       = uuid.NewString()
	senderName = uuid.NewString()
)

func setupGenerateContact(t *testing.T) func(*testing.T) {
	handler, err := NewHandler(
		context.Background(),
		contact.WithAppID(&ret.AppID, true),
		contact.WithAccount(&ret.Account, true),
		contact.WithAccountType(&ret.AccountType, true),
		contact.WithUsedFor(&ret.UsedFor, true),
		contact.WithSender(&ret.Sender, true),
	)
	assert.Nil(t, err)

	contact1, err := handler.CreateContact(context.Background())
	assert.Nil(t, err)

	handler.ID = &contact1.ID

	return func(*testing.T) {
		_, _ = handler.DeleteContact(context.Background())
	}
}

func generateContact(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		contact.WithAppID(&ret.AppID, true),
		contact.WithUsedFor(&ret.UsedFor, true),
		contact.WithSender(&ret.Sender, true),
		WithSubject(&subject, true),
		WithBody(&body, true),
		WithSenderName(&senderName, true),
	)
	assert.Nil(t, err)

	info, err := handler.GenerateContact(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestGenerateContact(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupGenerateContact(t)
	defer teardown(t)

	t.Run("generateContact", generateContact)
}
