package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	userName  string = "root"
	password  string = "abc"
	ipAddress string = "localhost"
	port      int    = 3306
	dbName    string = "hello"
	charset   string = "utf8"
)

type user struct {
	id int
	username string
	password string
	age int
}

func dao() (user, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	fmt.Println("connect database success")
	//查询
	var u user
	str := "select id,username,password,age from user where id=?"
	rowObj := Db.QueryRow(str, 2)
	err = rowObj.Scan(&u.id, &u.username, &u.password, &u.age)
	if err != nil {
		return user{}, errors.Wrap(err, "数据不存在")
	}
	return u, nil
}

func service() (user, error) {
	u, err := dao()
	if err != nil {
		return user{}, err
	}
	return u, nil
}

func main() {
	u, err := service()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(u)
	}
}
