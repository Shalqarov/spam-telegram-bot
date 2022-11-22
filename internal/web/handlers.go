package web

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gopkg.in/telebot.v3"
	"spam-telegram-bot/internal/repository/models"
)

type Handler struct {
	Bot      *telebot.Bot
	User     models.UserModel
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

type Message struct {
	Message string `json:"message"`
}

func SetRoutes(r *http.ServeMux, h *Handler) {
	r.HandleFunc("/api/send", h.Send)
	h.Bot.Handle("/start", h.Start)
	h.Bot.Handle("/delete", h.Delete)
}

func (h *Handler) Start(ctx telebot.Context) error {
	newUser := models.User{
		TelegramId: ctx.Chat().ID,
		Username:   ctx.Sender().Username,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
	}

	existUser, err := h.User.Find(newUser.TelegramId)
	if err != nil {
		if err == sql.ErrNoRows {
			h.InfoLog.Printf("Пользователь не найден, запись в БД %s...", newUser.Username)
		} else {
			h.ErrorLog.Printf("Ошибка поиска юзера: %v", err.Error())
			return nil
		}
	}
	if existUser == nil {
		err := h.User.Add(newUser)
		if err != nil {
			h.ErrorLog.Println("ошибка поиска юзера", err)
		}
	}
	return ctx.Send("Привет " + ctx.Sender().FirstName)
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(statusMessage(http.StatusMethodNotAllowed))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.ErrorLog.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(statusMessage(http.StatusBadRequest))
		return
	}

	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		h.InfoLog.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(statusMessage(http.StatusBadRequest))
		return
	}

	users, err := h.User.All()
	if err != nil {
		if err == sql.ErrNoRows {
			h.InfoLog.Println(err.Error())
			w.WriteHeader(http.StatusOK)
			w.Write(statusMessage(http.StatusOK))
			return
		}
		h.ErrorLog.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(statusMessage(http.StatusInternalServerError))
		return
	}

	for _, user := range users {
		u := &telebot.User{
			ID:        user.TelegramId,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
		}
		_, err := h.Bot.Send(u, msg.Message)
		if err != nil {
			h.ErrorLog.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(statusMessage(http.StatusInternalServerError))
			return
		}
	}
	w.Write(statusMessage(200))
}

func (h *Handler) Delete(context telebot.Context) error {
	err := h.User.Delete(context.Sender().ID)
	if err != nil {
		h.ErrorLog.Println("ошибка удаления юзера", err)
	}
	return context.Send("Вы были удалены")
}
