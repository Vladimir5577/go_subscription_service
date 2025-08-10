package helper

type ErrorResponse struct {
	Message    string
	StatusCode int
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) ErrorStatusCode() int {
	return e.StatusCode
}
