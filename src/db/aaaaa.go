package db

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

var (
	db *sql.DB
	err error
)

func init(){
	db, err = sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
}

func Insert() int64{

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-10")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}