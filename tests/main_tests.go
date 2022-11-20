package tests

import (
	"fmt"
	"github.com/Shalqarov/spam-telegram-bot/internal/repository/migrations"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	db, err := migrations.SqliteMigration("spambot.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(db)
}
