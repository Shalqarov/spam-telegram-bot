package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Shalqarov/spam-telegram-bot/configs"
)

func main() {
	configPath := flag.String("config", "local.toml", "Path to config file")
	flag.Parse()

	cfg, err := configs.NewConfig(*configPath)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}
	fmt.Println(cfg)
}
