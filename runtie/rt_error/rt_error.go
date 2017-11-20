package rt_error

type ErrorType int

const (
	HTTPERROR ErrorType = iota // http  error
	FILERROR // file read/write error
	SQLERROR // sql error
	DEVICEERROR // device error
	PARAMERROR // param error
)

type Customerror struct {
	Code ErrorType
	Info string
}

func (err Customerror) Error() string {
	info := err.Info
	switch err.Code {
	case DEVICEERROR:
		info += ", device error"

	case FILERROR:
		info += ", file error"

	case HTTPERROR:
		info += ", http error"

	case SQLERROR:
		info += ", sql error"

	case PARAMERROR:
		info += ", param error"
	}

	return info
}
