package models

import db "goapi.com/event/DB"

type User struct {
	ID       int64  `binding: "required"`
	FullName string `binding: "required"`
	Email    string `binding: "required"`
	PhoneNo  string `binding: "required"`
	Password string `binding: "required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (full_name, email, phone_no, password)
	          VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.FullName, u.Email, u.PhoneNo, u.Password)
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
