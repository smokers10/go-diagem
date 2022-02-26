package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	c := ReadConfig()

	t.Run("check application config", func(t *testing.T) {
		appc := c.Application
		assert.NotNil(t, appc)
		t.Run("check application config properties", func(t *testing.T) {
			assert.NotEmpty(t, appc.APP_PORT)
			assert.NotEmpty(t, appc.APP_Secret)
			assert.NotEmpty(t, appc.APP_Production_Mode)
			assert.NotEmpty(t, appc.APP_Base_URL)
		})
	})

	t.Run("check database config", func(t *testing.T) {
		dbc := c.Database
		assert.NotNil(t, c.Database)
		t.Run("check database config properties", func(t *testing.T) {
			assert.NotEmpty(t, dbc.MYSQL_Host)
			assert.NotEmpty(t, dbc.MYSQL_Port)
			assert.NotEmpty(t, dbc.MYSQL_Username)
			assert.NotEmpty(t, dbc.MYSQL_Database)
			assert.NotEmpty(t, dbc.MYSQL_Password)
			assert.NotEmpty(t, dbc.MYSQL_SessionTable)

			assert.NotEmpty(t, dbc.MYSQL_URI)
			assert.NotEmpty(t, dbc.MYSQL_Max_Connection)
			assert.NotEmpty(t, dbc.MYSQL_Max_Connection_Life_Time)
			assert.NotEmpty(t, dbc.MYSQL_Max_Idle_Connection)
			assert.NotEmpty(t, dbc.MYSQL_Max_Idle_Connection_Life_Time)
		})
	})

	t.Run("check ssh database config", func(t *testing.T) {
		sshc := c.SSHDatabase
		assert.NotNil(t, sshc)
		t.Run("check ssh database config properties", func(t *testing.T) {
			assert.NotEmpty(t, sshc.UseSSH)
			assert.NotEmpty(t, sshc.SSH_Host)
			assert.NotEmpty(t, sshc.SSH_Port)
			assert.NotEmpty(t, sshc.SSH_Password)

			assert.NotEmpty(t, sshc.MYSQL_Username)
			assert.NotEmpty(t, sshc.MYSQL_Host)
			assert.NotEmpty(t, sshc.MYSQL_Name)
			assert.NotEmpty(t, sshc.MYSQL_Password)
		})
	})

	t.Run("check smtp config", func(t *testing.T) {
		smtpc := c.SMTP
		assert.NotNil(t, smtpc)
		t.Run("check smtp config properties", func(t *testing.T) {
			assert.NotEmpty(t, smtpc.SMTP_Host)
			assert.NotEmpty(t, smtpc.SMTP_Port)
			assert.NotEmpty(t, smtpc.SMTP_Sender_Name)
			assert.NotEmpty(t, smtpc.SMTP_Auth_Username)
			assert.NotEmpty(t, smtpc.SMTP_Auth_Password)
		})
	})

	t.Run("check etc config", func(t *testing.T) {
		etcc := c.ETC
		assert.NotNil(t, etcc)
		t.Run("check etc config properties", func(t *testing.T) {
			assert.NotEmpty(t, etcc.Midtrans_Server_key)
			assert.NotEmpty(t, etcc.Rajaongkir_API_Key)
			assert.NotEmpty(t, etcc.RechaptaServerKey)
			assert.NotEmpty(t, etcc.RechaptaSiteKey)
		})
	})
}
