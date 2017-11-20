package main

import (
	"fmt"
	"runtie/rt_mysql"
)

func main() {
	i := 10
	fmt.Println(i)
	fmt.Println("润铁科技有限责任公司")

	sql := rt_mysql.SQL{}
	err := sql.Open("runtie",
		"root",
		"hou1415926",
		"127.0.0.1",
		"3306")
	if err != nil {
		fmt.Println("Open mysql error")
	} else {
		fmt.Println("Open mysql success")
	}
}


