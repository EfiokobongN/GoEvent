package models

import (
	"errors"

	db "goapi.com/event/DB"
	util "goapi.com/event/Util"
)

type User struct {
	ID       int64  `json:"id"`
	FullName string `binding: "required"`
	Email    string `binding: "required"`
	PhoneNo  string `binding: "required"`
	Password string `binding: "required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (fullName, email, phoneNo, password)
	          VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.FullName, u.Email, u.PhoneNo, hashPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (login *User) ValidateLogin() error {
	query := "SELECT id, password FROM users WHERE email= ?"
	row := db.DB.QueryRow(query, login.Email)

	var retrievedPassword string
	err := row.Scan(&login.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid login details")
	}

	passwordValid := util.CheckPassword(login.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("invalid login details")
	}
	return nil
}
