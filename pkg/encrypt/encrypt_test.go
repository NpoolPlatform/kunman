//nolint:nolintlint,gosec
package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	// 1. generate a salt.
	salt := Salt()

	// 2. input a password.
	truePassword := "lpz990627"

	// 3. encrypte password.
	enPass, err := EncryptWithSalt(truePassword, salt)
	assert.Nil(t, err)

	// 4. mock user input pasword
	inputPass := "lpz990627"
	err = VerifyWithSalt(inputPass, enPass, salt)
	assert.Nil(t, err)
}
