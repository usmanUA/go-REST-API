package models

import (
	"errors"
	"fmt"

	db "example.com/REST-API/DB"
	"example.com/REST-API/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid Credentials.")
	}
	fmt.Println("\033[31mhere\033[0m")
	validPassword := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !validPassword {
		return errors.New("Invalid Credentials.")
	}
	return nil
}
