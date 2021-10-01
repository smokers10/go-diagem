package domain

type Response struct {
	Message string
	Data    interface{}
	Success bool
	Status  int
	Token   string
}
