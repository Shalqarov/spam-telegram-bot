package bot

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
	"spam-telegram-bot/internal/repository/models"
)

type SpamBot struct {
	Bot  *telebot.Bot
	User models.UserModel
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}
