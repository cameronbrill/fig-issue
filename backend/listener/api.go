package listener

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
)

func Start(ctx context.Context, commentChan chan<- *figma.FigmaFileCommentResponse) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/figma", func(w http.ResponseWriter, r *http.Request) {
		var res *figma.FigmaFileCommentResponse
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res.Passcode != "secretpasscode" {
			http.Error(w, "Invalid passcode", http.StatusUnauthorized)
		}

		commentChan <- res

		w.WriteHeader(http.StatusOK)
	})
	log.Info("starting figma listener on port :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
