package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	sql.DB
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

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) AddUser(user User) error {
	_, err := m.DB.Exec("insert into Users(telegram_id,username, firstName, lastName ) values(?,?,?,?)", user.Telegram_id, user.Username, user.FirstName, user.LastName)

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) FindOne(userFind User) (*User, error) {

	user := User{}

	err := m.DB.QueryRow("select * from Users where telegram_id=?", userFind.Telegram_id).Scan(&user.Id, &user.Telegram_id, &user.Username, &user.FirstName, &user.LastName)

	if err != nil {
		return nil, nil
	}

	return &user, err

}
func (m *UserModel) SelectAll() []User {
	allUsers := []User{}
	err, _ := m.DB.Query("select * from Users")

	for err.Next() {
		var user User
		if err := err.Scan(&user.Id, &user.Telegram_id, &user.Username, &user.FirstName, &user.LastName); err != nil {
			fmt.Print(err)
		}
		allUsers = append(allUsers, user)
	}

	return allUsers

}
