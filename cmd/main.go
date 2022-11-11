package main

import (
	"flag"
	"fmt"
	"github.com/Shalqarov/spam-telegram-bot/configs"
	"github.com/Shalqarov/spam-telegram-bot/internal/repository"
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

	_, err = repository.SqliteMigration(cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
