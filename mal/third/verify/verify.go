package verify

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/mal/third/google"
)

func (h *Handler) VerifyGoogleRecaptchaV3(ctx context.Context) error {
	if h.RecaptchaToken == "" {
		return fmt.Errorf("recaptchaToken is empty")
	}
	return google.VerifyGoogleRecaptchaV3(h.RecaptchaToken)
}
