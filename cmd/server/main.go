package main

import (
	"encoding/json"
	"fmt"
	"neolist-backend/internal/app"
	"neolist-backend/internal/config"
	"neolist-backend/internal/db"
	"neolist-backend/internal/handlers"
	"neolist-backend/internal/repositories"
	"neolist-backend/internal/services"
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

	// user routes
	userRepo := repositories.NewUserRepository(newDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	router.Handle("POST /users/create", userHandler.RegisterHandler)
	router.Handle("GET /users", userHandler.ListHandler)
	router.Handle("GET /users/{id}", userHandler.FindByIDHandler)
	router.Handle("PUT /users/update", userHandler.UpdateHandler)
	router.Handle("PATCH /users/soft-delete/{id}", userHandler.SoftDeleteHandler)
	router.Handle("DELETE /users/force-delete/{id}", userHandler.ForceDeleteHandler)

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
