package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"spam-telegram-bot/configs"
	"spam-telegram-bot/internal/bot"
	"spam-telegram-bot/internal/repository/migrations"
	"spam-telegram-bot/internal/repository/models"
	"spam-telegram-bot/internal/web"
)

func Run(cfg *configs.Config) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := migrations.SqliteMigration(cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	router := http.NewServeMux()
	handler := &web.Handler{
		Bot:      bot.InitBot(cfg.BotToken),
		User:     models.UserModel{DB: db},
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
	web.SetRoutes(router, handler)

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			handler.ErrorLog.Fatalln(err)
		}
	}()
	handler.InfoLog.Println("Server is listening...")
	handler.Bot.Start()
}
