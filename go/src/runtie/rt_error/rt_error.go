package rt_error

type ErrorType int

const (
	HTTPERROR ErrorType = iota
	FILERROR
	SQLERROR
	DEVICEERROR
	PARAMERROR
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
