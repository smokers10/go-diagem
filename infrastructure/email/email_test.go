package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := SMTP().NativeFire([]string{"nadzarmutaqin4@gmail.com"}, "good morning! why 2023", "hey good morning budy! let's fuck")
	assert.Empty(t, err)
}
