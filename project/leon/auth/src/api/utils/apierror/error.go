package apierror

type ApiError struct {
	Message string `json:"message"`
	Status  int  `json:"status"`
	Error   error  `json:"error"`
}
