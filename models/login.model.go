package models

import (
	"database/sql"
	"fmt"

	"github.com/mafr017/golang-rest-echo/db"
	"github.com/mafr017/golang-rest-echo/helpers"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("Err: Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Err: Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Err: Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}