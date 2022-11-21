package models

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestAddUser(t *testing.T) {
	req := require.New(t)

	db, err := sql.Open("sqlite3", "test.db")
	req.Equal(nil, err)
	defer db.Close()

	query, err := os.ReadFile("./../migrations/migration.sql")
	req.Equal(nil, err)
	if _, err := db.Exec(string(query)); err != nil {
		req.Equal(nil, err)
	}

	model := UserModel{DB: db}

	newUser := User{
		TelegramId: 8797,
		Username:   "AbobaTest",
	}

	err = model.AddUser(newUser)

	req.Equal(nil, err)

	exist, err := model.FindOne(newUser)

	req.Equal(nil, err)

	expected := fmt.Sprintf("telegram_id=%v, Username=%v", newUser.TelegramId, newUser.Username)
	got := fmt.Sprintf("telegram_id=%v, Username=%v", exist.TelegramId, exist.Username)
	req.Equal(expected, got)

}

func TestSelectAll(t *testing.T) {
	req := require.New(t)

	db, err := sql.Open("sqlite3", "test.db")

	req.Equal(nil, err)

	defer db.Close()

	model := UserModel{DB: db}

	testUsers := []User{
		User{
			TelegramId: 8797,
			Username:   "AbobaTest",
		},
	}

	selectAllUsers, err := model.SelectAll()

	req.Equal(nil, err)
	req.Equal(testUsers[0].TelegramId, selectAllUsers[0].TelegramId)
	req.Equal(testUsers[0].Username, selectAllUsers[0].Username)

	_ = os.Remove("test.db")
}
