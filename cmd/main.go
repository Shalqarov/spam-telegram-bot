package main

import (
	"flag"
	"github.com/Shalqarov/spam-telegram-bot/configs"
	"github.com/Shalqarov/spam-telegram-bot/internal/bot"
	"github.com/Shalqarov/spam-telegram-bot/internal/repository/migrations"
	"github.com/Shalqarov/spam-telegram-bot/internal/repository/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	configPath := flag.String("config", "local.toml", "Path to config file")
	flag.Parse()

	cfg, err := configs.NewConfig(*configPath)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := migrations.SqliteMigration(cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	spambot := bot.SpamBot{
		Bot:  bot.InitBot(cfg.BotToken),
		User: models.UserModel{db},
	}

	spambot.Bot.Handle("/start", spambot.StartHandler)

	spambot.Bot.Start()

}
