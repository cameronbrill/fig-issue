package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
)

// TODO move this from /listener to /figma/listener
// TODO also move /model/figma to /figma/model
func Start(ctx context.Context, commentChan chan<- *figma.FileCommentResponse) *http.Server {
	r := chi.NewRouter()

	if os.Getenv("APP_ENV") == "dev" || os.Getenv("APP_ENV") == "prod" {
		r.Use(middleware.Logger)
	}

	validate := validator.New()

	wbhkSvc := &webhookSvc{commentChan, validate}

	r.Post("/figma", wbhkSvc.figmaHandler)

	svc := &http.Server{Addr: ":3000", Handler: r}

	go func() {
		<-ctx.Done()
		if err := svc.Shutdown(context.Background()); err != nil {
			log.Fatalf("shutting down listener server: %v", err)
		}
	}()

	return svc
}

type webhookSvc struct {
	commentChan chan<- *figma.FileCommentResponse
	validate    *validator.Validate
}

func (svc *webhookSvc) figmaHandler(w http.ResponseWriter, r *http.Request) {
	var res figma.FileCommentResponse
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error("reading response body")
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error("parsing response body")
		return
	}

	err = svc.validate.Struct(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), 400)
		log.Errorf("validating response body: %v", err)
		return
	}

	if res.Passcode != "secretpasscode" {
		http.Error(w, "Invalid passcode", http.StatusUnauthorized)
		log.Error("detected invalid passcode in response")
		return
	}

	svc.commentChan <- &res

	w.WriteHeader(http.StatusOK)
}
