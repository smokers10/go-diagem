package domain

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success,omitempty"`
	Status  int         `json:"status,omitempty"`
	Token   string      `json:"token,omitempty"`
}
