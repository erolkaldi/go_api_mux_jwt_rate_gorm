package models

type Response struct {
	IsSuccessfull bool   `json:"is_successfull"`
	Message       string `json:"message"`
}
