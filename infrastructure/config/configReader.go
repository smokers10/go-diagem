package config

import (
	"encoding/json"
	"io/ioutil"
)

type configurationScheme struct {
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
	APP_Secret                          string `json:"app_secret"`
	APP_PORT                            string `json:"app_port"`
	APP_Production_Mode                 string `json:"app_production_mode"`
	Midtrans_Server_key                 string `json:"midtrans_server_key"`
	Rajaongkir_API_Key                  string `json:"rajaongkir_api_key"`
}

func ReadConfig() *configurationScheme {
	result := configurationScheme{}

	rawConfig, err := ioutil.ReadFile("app.config.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(rawConfig, &result)

	return &result
}
