package server

const (
	ERROR_NO_COMMAND_FOUND         = "no command found"
	ERROR_STATUS_COMMAND_NOT_FOUND = "command not found"
)

// IResponse define an interface to repond to a Request
type IResponse interface {
	Payload() string
	Error() error
	ErrorType() string
}

// Response represent a reponse for thoth commands
// It implements a IResponse
type Response struct {
	Pl   string
	Err  error
	ErrT string
}

func (resp Response) Payload() string {
	return resp.Pl
}

func (resp Response) Error() error {
	return resp.Err
}

func (resp Response) ErrorType() string {
	return resp.ErrT
}
