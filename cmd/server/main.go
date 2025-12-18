package main

import (
	"encoding/json"
	"fmt"
	"neolist-backend/internal/app"
	"neolist-backend/internal/config"
	"neolist-backend/internal/db"
	"neolist-backend/internal/handlers"
	"neolist-backend/internal/repositories/mysql"
	"neolist-backend/internal/service/user"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	cfg := config.Load()
	router := app.NewRouter()
	newDB := db.NewDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	router.Handle("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("message: Server is running ok")
	})

	userRepo := mysql.NewUserRepository(newDB)
	userService := user.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	router.Handle("/user/create", userHandler.RegisterHandler)

	fmt.Printf("Server running on %v\n", cfg.Port)
	go http.ListenAndServe(cfg.Port, router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	err := newDB.Close()
	if err != nil {
		println(err)
	}

	println("Database Closed")

	println("Shut down server")
}
