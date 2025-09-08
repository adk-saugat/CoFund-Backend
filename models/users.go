package models

import (
	"errors"
	"time"

	"github.com/adk-saugat/cofund/db"
	"github.com/adk-saugat/cofund/utils"
)

type User struct{
	ID 			int64 		`json:"id"`
	FirstName	string 		`json:"firstName"`
	LastName 	string 		`json:"lastName"`
	Email 		string 		`json:"email" binding:"required"`
	Password 	string 		`json:"password" binding:"required"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

func GetProfileById(userId int64) (*User, error){
	query :=  `
		SELECT id, firstName, lastName, email FROM users where id = ?
	`
	row := db.DB.QueryRow(query, userId)

	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil{
		return nil, err
	}
	
	return &user, nil
}

func (user *User) ValidateCredentials() error{
	query := `
		SELECT id, password FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil{
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func (user *User) Save() error{
	query := `
		INSERT INTO users (firstName, lastName, email, password, createdAt)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//hashing password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result , err := stmt.Exec(user.FirstName, user.LastName, user.Email, hashedPassword, time.Now())
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	user.CreatedAt = time.Now()
	user.Password = hashedPassword

	return err
}