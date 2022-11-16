package bot

import (
	"net/http"
)

func SetRoutes(r *http.ServeMux, bot *SpamBot) {
	r.HandleFunc("/api/send", bot.SendHandler)
}
