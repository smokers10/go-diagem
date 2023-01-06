package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	plaintext := "password123"
	hashed := Hash(plaintext)

	t.Run("check hashed", func(t *testing.T) {
		assert.NotEmpty(t, hashed)
		assert.NotEqual(t, hashed, plaintext)
	})

	t.Run("check compare", func(t *testing.T) {
		isMatch := Compare(hashed, plaintext)

		t.Run("correct plaintext", func(t *testing.T) {
			assert.Equal(t, true, isMatch)
		})

		t.Run("wrong plaintext", func(t *testing.T) {
			isMatch := Compare(hashed, "ajun!ajin")
			assert.Equal(t, false, isMatch)
		})
	})
}
