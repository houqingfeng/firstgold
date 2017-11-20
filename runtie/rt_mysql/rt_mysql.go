package rt_mysql

import (
	"runtie/rt_error"
	"database/sql"
	_ "github.com/mysql"
	"fmt"
)

type SQL struct {
	db *sql.DB
}

func (rt_sql SQL)Open(database string,
					  user string,
					  pwd string,
	               	  ip string,
	               	  port string) error {
	commad := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, database)
	db, err := sql.Open("mysql", commad)
	if err != nil {
		rt_err := rt_error.Customerror{}
		rt_err.Code = rt_error.SQLERROR
		rt_err.Info = "Open database "
		rt_err.Info += commad
		return rt_err
	}

	rt_sql.db = db
	return nil
}

func (rt_sql SQL)Create(table string, values map[string]interface){

}

func (rt_sql SQL)Select(table string, condition string) (error, map[string]string) {
	err := rt_error.Customerror{}
	if len(table) == 0 || len(condition) == 0 {
		err.Code = rt_error.PARAMERROR
		err.Info = "Select"
		err.Info += condition
	}

	query := make(map[string]string)
	return err, query
}
