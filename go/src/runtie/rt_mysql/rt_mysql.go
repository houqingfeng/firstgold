package rt_mysql

import (
	"runtie/rt_error"
)

type SQL struct {
}

func (sql SQL)Select(table string, condition string) (error, map[string]string) {
	err := rt_error.Customerror{}
	if len(table) == 0 || len(condition) == 0 {
		err.Code = rt_error.PARAMERROR
		err.Info = "Select"
		err.Info += condition
	}

	query := make(map[string]string)

	return err, query
}
