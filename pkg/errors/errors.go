package errors

func (err *Error) Error() string {
	return err.Title
}

func New(status int, title string) error {
	return &Error{
		Status: status,
		Title:  title,
	}
}
