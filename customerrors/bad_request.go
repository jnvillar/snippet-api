package customerrors

type BadRequest struct {
	error error
	code  int
}

func NewBadRequestError(err error) *BadRequest {
	return &BadRequest{error: err, code: 400}
}

func (e *BadRequest) Error() string {
	return e.error.Error()
}

func (e *BadRequest) GetCode() int {
	return e.code
}
