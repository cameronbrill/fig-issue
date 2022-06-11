package listener

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
)

func Start(ctx context.Context, commentChan chan<- *figma.FigmaFileCommentResponse) *http.Server {
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

	svc := &http.Server{Addr: ":3000", Handler: r}

	go func() {
		<-ctx.Done()
		svc.Shutdown(context.Background())
	}()

	return svc
}
