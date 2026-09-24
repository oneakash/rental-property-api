package errors

type APIError struct {
	StatusCode int
	Field string
	Message string
}

func (e *APIError) Error() string{
	return e.Field + ": "+e.Message
}