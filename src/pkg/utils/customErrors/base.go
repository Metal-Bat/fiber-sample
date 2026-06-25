package customErrors

type InternalError struct {
	HTTPStatus int
	MessageKey string
	Err        error
}

func (e *InternalError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.MessageKey
}

func (e *InternalError) Unwrap() error {
	return e.Err
}
