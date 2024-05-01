package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ZyoGo/Backend-Challange/cmd/modules"
	"github.com/ZyoGo/Backend-Challange/config"
	"github.com/ZyoGo/Backend-Challange/pkg/db"
	"github.com/ZyoGo/Backend-Challange/pkg/http/logger"
	"github.com/ZyoGo/Backend-Challange/pkg/response"
	"github.com/gorilla/mux"
)

type healthResponse struct {
	Id           string `json:"id"`
	StatusHealth string `json:"status_health"`
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Id:           "12",
		StatusHealth: "Ok",
	}

	response.Encode(w, resp, http.StatusOK)
}

func main() {
	cfg := config.GetConfig()
	dbCon := db.NewDatabaseConnection(cfg)

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Use(logger.ZapLogger(logger.GetLogger()))
	r.HandleFunc("/health", health).Methods("GET")
	modules.RegisterModules(r, dbCon)

	srv := &http.Server{
		Addr: ":4001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
