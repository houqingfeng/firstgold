package rt_error
import "fmt"
type ErrorType int

const (
	HTTPERROR ErrorType = iota
	FILERROR
	SQLERROR
	DEVICEERROR
)

type Customerror struct {
	code ErrorType
	err error
}

func (err Customerror)Info() string {
	info := fmt.Sprintf("%s", )
}