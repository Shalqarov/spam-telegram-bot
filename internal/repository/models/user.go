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

func (m *UserModel) Add(user User) error {
	_, err := m.DB.Exec(`insert into users(telegram_id, username,first_name, last_name) VALUES (?,?,?,?)`, user.TelegramId, user.Username, user.FirstName, user.LastName)
	return err
}

func (m *UserModel) Find(telegramId int64) (*User, error) {
	user := User{}
	err := m.DB.QueryRow(`select * from users where telegram_id=?`, telegramId).Scan(&user.Id, &user.TelegramId, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Delete(telegramId int64) error {
	_, err := m.DB.Exec(`DELETE FROM users WHERE telegram_id=?`, telegramId)
	return err
}

func (m *UserModel) All() ([]*User, error) {
	var allUsers []*User
	rows, err := m.DB.Query(`select * from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.Id, &user.TelegramId, &user.Username, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}
