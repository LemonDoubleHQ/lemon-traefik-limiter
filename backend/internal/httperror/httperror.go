package httperror

type HttpError struct {
	StatusCode int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}