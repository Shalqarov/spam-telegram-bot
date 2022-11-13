package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id          int64
	Telegram_id int64  `json:"telegram_id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type Service interface {
	CreateUser(user *User) (int64, error)
	GetUserById(id int64) (*User, error)
}

type Repo interface {
	CreateUser(user *User) (int64, error)
	GetUserById(id int64) (*User, error)
}

func AddUser(user User, db *sql.DB) error {
	_, err := db.Exec("insert into Users(telegram_id,username, firstName, lastName ) values(?,?,?,?)", user.Telegram_id, user.Username, user.FirstName, user.LastName)

	if err != nil {
		return err
	}

	return nil
}

func FindOne(userFind User, Db *sql.DB) (*User, error) {

	user := User{}

	err := Db.QueryRow("select * from Users where telegram_id=?", userFind.Telegram_id).Scan(&user.Id, &user.Telegram_id, &user.Username, &user.FirstName, &user.LastName)

	if err != nil {
		return nil, nil
	}

	return &user, err

}
func SelectAll(db *sql.DB) []User {
	allUsers := []User{}
	err, _ := db.Query("select * from Users")

	for err.Next() {
		var user User
		if err := err.Scan(&user.Id, &user.Telegram_id, &user.Username, &user.FirstName, &user.LastName); err != nil {
			fmt.Print(err)
		}
		allUsers = append(allUsers, user)
	}

	return allUsers

}
