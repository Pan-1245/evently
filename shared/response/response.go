package response

type Wrapper[T any] struct {
	StatusCode int            `json:"status_code"`
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
	Data       DataWrapper[T] `json:"data"`
}

type DataWrapper[T any] struct {
	Data  *T  `json:"Data,omitempty"`
	Error any `json:"Error,omitempty"`
}
