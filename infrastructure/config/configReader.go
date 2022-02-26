package config

import (
	"encoding/json"
	"io/ioutil"
)

type configurationScheme struct {
	Application *APP      `json:"application"`
	Database    *MYSQL    `json:"database"`
	SSHDatabase *SSHMYSQL `json:"database_ssh"`
	SMTP        *SMTP     `json:"smtp"`
	ETC         *ETC      `json:"etc"`
}

type APP struct {
	APP_Secret          string `json:"app_secret"`
	APP_PORT            string `json:"app_port"`
	APP_Production_Mode string `json:"app_production_mode"`
	APP_Base_URL        string `json:"app_base_url"`
}

type MYSQL struct {
	MYSQL_Host                          string `json:"mysql_host"`
	MYSQL_Port                          int    `json:"mysql_port"`
	MYSQL_Username                      string `json:"mysql_username"`
	MYSQL_Database                      string `json:"mysql_database"`
	MYSQL_Password                      string `json:"mysql_password"`
	MYSQL_SessionTable                  string `json:"mysql_session_table"`
	MYSQL_URI                           string `json:"mysql_uri"`
	MYSQL_Max_Connection                int    `json:"mysql_max_connection"`
	MYSQL_Max_Connection_Life_Time      int    `json:"mysql_max_connection_life_time"`
	MYSQL_Max_Idle_Connection           int    `json:"mysql_max_idle_connection"`
	MYSQL_Max_Idle_Connection_Life_Time int    `json:"mysql_max_idle_connection_life_time"`
}

type SSHMYSQL struct {
	UseSSH         bool   `json:"use_ssh"`
	SSH_Host       string `json:"ssh_host"`
	SSH_Port       int    `json:"ssh_port"`
	SSH_Password   string `json:"ssh_password"`
	MYSQL_Username string `json:"db_username"`
	MYSQL_Host     string `json:"db_host"`
	MYSQL_Name     string `json:"db_name"`
	MYSQL_Password string `json:"db_password"`
}

type SMTP struct {
	SMTP_Host          string `json:"SMTP_host"`
	SMTP_Port          int    `json:"SMTP_port"`
	SMTP_Sender_Name   string `json:"SMTP_sender_name"`
	SMTP_Auth_Username string `json:"SMTP_auth_username"`
	SMTP_Auth_Password string `json:"SMTP_auth_password"`
}

type ETC struct {
	Midtrans_Server_key string `json:"midtrans_server_key"`
	Rajaongkir_API_Key  string `json:"rajaongkir_api_key"`
	RechaptaServerKey   string `json:"rechapta_server_key"`
	RechaptaSiteKey     string `json:"rechapta_site_key"`
}

func ReadConfig() *configurationScheme {
	result := configurationScheme{}

	r, err := ioutil.ReadFile("app.config.json")
	if err != nil {
		r, _ := ioutil.ReadFile("../../app.config.json")
		json.Unmarshal(r, &result)
		return &result
	}

	json.Unmarshal(r, &result)
	return &result
}
