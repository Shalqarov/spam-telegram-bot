package models

import (
	"database/sql"
)

type User struct {
	Id         int64
	TelegramId int64  `json:"telegram_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) AddUser(user User) error {
	_, err := m.DB.Exec(`insert into "users" ("telegram_id","username", "first_name", "last_name") values(?,?,?,?)`, user.TelegramId, user.Username, user.FirstName, user.LastName)
	return err
}

func (m *UserModel) FindOne(userFind User) (*User, error) {

	user := User{}

	err := m.DB.QueryRow(`select * from "users" where telegram_id=?`, userFind.TelegramId).Scan(&user.Id, &user.TelegramId, &user.Username, &user.FirstName, &user.LastName)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
func (m *UserModel) SelectAll() ([]*User, error) {
	var allUsers []*User
	err, _ := m.DB.Query(`select * from "users"`)

	for err.Next() {
		user := &User{}
		if err := err.Scan(&user.Id, &user.TelegramId, &user.Username, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}
