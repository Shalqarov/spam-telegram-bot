package bot

import (
	"fmt"
	"github.com/Shalqarov/spam-telegram-bot/internal/repository/models"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type SpamBot struct {
	Bot  *telebot.Bot
	User models.UserModel
}

func (bot *SpamBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Telegram_id: ctx.Chat().ID,
		Username:    ctx.Sender().Username,
		FirstName:   ctx.Sender().FirstName,
		LastName:    ctx.Sender().LastName,
	}

	existUser, err := bot.User.FindOne(newUser)

	if err != nil {
		fmt.Errorf("Ошибка поиска юзера %v", err)
	}

	if existUser == nil {
		err := bot.User.AddUser(newUser)

		if err != nil {
			fmt.Errorf("Ошибка создания юзера %v", err)
		}
	}

	return ctx.Send("Привет " + ctx.Sender().FirstName)
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
