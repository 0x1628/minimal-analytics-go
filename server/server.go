package server

import (
	"context"
	"github.com/fragment0/minimal-analytics-go/util"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func Start() {
	r := chi.NewRouter()

	r.Post("/event", CollectEvent)
	r.Post("/register", CollectRegister)
	r.Post("/crash", CollectCrash)

	r.Get("/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(204)
	})

	start(r)
}

func start(handler http.Handler) {
	config := util.GetConfig()
	s := strings.Join([]string{config.Server.Host, ":", strconv.Itoa(config.Server.Port)}, "")

	server := &http.Server{Addr: s, Handler: handler}

	go func() {
		log.Println("Server Start: " + s)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<- stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Panicln(err)
	}
}
