package entity

type Error struct {
	Code    int
	Message string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}
