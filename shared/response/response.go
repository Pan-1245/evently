package response

type SuccessWrapper[T any] struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       *T     `json:"data"`
}

type ErrorWrapper struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}
