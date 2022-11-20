package tests

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"os"
	"spam-telegram-bot/internal/repository/models"
	"testing"
)

func TestAddUser(t *testing.T) {
	require := require.New(t)

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		require.Equal(nil, err)
	}

	query, err := os.ReadFile("../internal/repository/migrations/migration.sql")

	if err != nil {
		require.Equal(nil, err)
	}

	if _, err := db.Exec(string(query)); err != nil {
		require.Equal(nil, err)
	}

	model := models.UserModel{DB: db}

	newUser := models.User{
		TelegramId: 8797,
		Username:   "AbobaTest",
	}

	err = model.AddUser(newUser)

	if err != nil {
		require.Equal(nil, err)
	}

	exist, err := model.FindOne(newUser)

	if err != nil {
		require.Equal(nil, err)
	}

	expected := fmt.Sprintf("telegram_id=%v, Username=%v", newUser.TelegramId, newUser.Username)
	got := fmt.Sprintf("telegram_id=%v, Username=%v", exist.TelegramId, exist.Username)
	require.Equal(expected, got)

}

func TestSelectAll(t *testing.T) {
	file, _ := os.Open("test.db")
	fmt.Print(file)
}
