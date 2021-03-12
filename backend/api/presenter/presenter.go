package presenter

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func NewResponse(success bool, data interface{}, errMessage string) *Response {
	return &Response{
		Success: success,
		Data:    data,
		Error:   errMessage,
	}
}
