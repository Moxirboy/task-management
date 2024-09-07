package dto

type BaseResponse struct {
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	InvalidParams interface{} `json:"invalidParams,omitempty"`
	Data          interface{} `json:"data"`
}
