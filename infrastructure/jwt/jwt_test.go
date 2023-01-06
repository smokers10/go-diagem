package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignKey(t *testing.T) {
	sKey := signKey()

	assert.NotEmpty(t, sKey)
}

func TestSignVerify(t *testing.T) {
	payload := Payload{
		ID:         1,
		Email:      "johndoe@gmail.com",
		Type:       "user",
		IsVerified: true,
		IsLogged:   true,
	}

	token := Sign(&payload)

	t.Run("sign token result should not empty", func(t *testing.T) {
		assert.NotEmpty(t, token)
	})

	t.Run("verify every payload", func(t *testing.T) {
		verifiedPayload := Verify(fmt.Sprintf("Bearer %s", token))

		assert.NotEmpty(t, verifiedPayload)
	})
}
