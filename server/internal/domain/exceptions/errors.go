package exceptions

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
