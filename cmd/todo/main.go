package main

import (
	"fmt"
	"log"
	"todoapp/internal/config"
	"todoapp/internal/database"
	"todoapp/internal/handler"
	"todoapp/internal/route"
	"todoapp/internal/server"
	"todoapp/internal/service"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		fmt.Print(err)
		return 
	}
	service := service.NewService(db)

	handler := handler.NewHandler(service)

	handlers := route.New(handler)

	if err := db.RunMigrations(); err != nil {
		fmt.Printf("run migrations: %v", err)
		return
	}
	fmt.Println("migrations applied, database is ready")
	
	srv := server.New(cfg.ServerPort, handlers)
	if err := srv.StartServer(); err != nil {
    	log.Fatalf("start server: %v", err)
	}
}