package customerrors

type NotFound struct {
	error error
	code  int
}

func NewNotFoundError(err error) *NotFound {
	return &NotFound{error: err, code: 404}
}

func (e *NotFound) Error() string {
	return e.error.Error()
}

func (e *NotFound) GetCode() int {
	return e.code
}
