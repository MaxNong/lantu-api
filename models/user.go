package models

import (
	"fmt"
	"lantu/dao"
)

type User struct {
	id       int
	username string
	password string
}

func QueryUserById(id int) (user User, err error) {
	s := "select * from user_tbl where id = ?"
	var u User
	queryErr := dao.DB.QueryRow(s, id).Scan(&u.id, &u.username, &u.password)

	if queryErr != nil {
		return User{}, queryErr
	} else {
		return u, nil
	}
}

func QueryUserByUsernameAndPassword(username string, password string) (user User, err error) {
	s := "select * from user_tbl where username = ? and password = ?"
	var u User
	queryErr := dao.DB.QueryRow(s, username, password).Scan(&u.id, &u.username, &u.password)

	if queryErr != nil {
		fmt.Printf("queryErr: %v\n", queryErr)
		return User{}, queryErr
	} else {
		return u, nil
	}
}
