package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jasontalon/ncov-tracker/api/handlers"
	log "github.com/sirupsen/logrus"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (*App) StartServer() {
	setHandlers()
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	log.Infof("listens to port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}

func GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(404)
			fmt.Fprint(w, "Not found")
			return
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Infof("request=%s;method=%s;elapsed: %s ", r.URL, r.Method, time.Since(start))
	})
}

func setHandlers() {
	http.Handle("/hospital", GET(http.HandlerFunc(handlers.Hospital)))
	http.Handle("/individual", GET(http.HandlerFunc(handlers.Individual)))
	http.Handle("/numbers", GET(http.HandlerFunc(handlers.Numbers)))
	http.Handle("/residence", GET(http.HandlerFunc(handlers.Residence)))
}
