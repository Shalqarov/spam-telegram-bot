package main

import (
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"spam-telegram-bot/configs"
	"spam-telegram-bot/internal/app"
)

var (
	cfg  *configs.Config
	addr *string
)

func init() {
	addr = flag.String("addr", ":5000", "Network address HTTP")
	configPath := flag.String("config", "local.toml", "Path to config file")
	flag.Parse()
	var err error
	cfg, err = configs.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}
	cfg.Addr = *addr
}

func main() {
	app.Run(cfg)
}
