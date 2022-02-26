package database

import (
	"testing"

	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"github.com/stretchr/testify/assert"
)

func TestViaTCP(t *testing.T) {
	config := config.ReadConfig()
	mysql := config.Database
	db, err := viaTCP(mysql)

	t.Run("check connection", func(t *testing.T) {
		if err != nil {
			assert.Empty(t, err.Error())
		}
	})

	t.Run("ping connection", func(t *testing.T) {
		if err := db.Ping(); err != nil {
			assert.Empty(t, err.Error())
		}
	})
}
