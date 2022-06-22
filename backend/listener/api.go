package listener

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
)

func Start(ctx context.Context, commentChan chan<- *figma.FileCommentResponse) *http.Server {
	r := chi.NewRouter()

	if os.Getenv("APP_ENV") == "dev" || os.Getenv("APP_ENV") == "prod" {
		r.Use(middleware.Logger)
	}

	r.Post("/figma", func(w http.ResponseWriter, r *http.Request) {
		var res figma.FileCommentResponse
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res.Passcode != "secretpasscode" {
			http.Error(w, "Invalid passcode", http.StatusUnauthorized)
			return
		}

		commentChan <- &res

		w.WriteHeader(http.StatusOK)
	})

	svc := &http.Server{Addr: ":3000", Handler: r}

	go func() {
		<-ctx.Done()
		if err := svc.Shutdown(context.Background()); err != nil {
			log.Fatalf("shutting down listener server: %v", err)
		}
	}()

	return svc
}
