package utils

type AppError struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
	Code string `json:"code"`
}
