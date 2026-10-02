package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Azizbenhammouda/SoundStream/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("POSTGRES_CONFIG")
	jwtSecret := os.Getenv("JWT_SECRET")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto`).Error; err != nil {
		log.Fatal("failed to enable pgcrypto:", err)
	}
	if err := db.AutoMigrate(&users.User{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo, jwtSecret)
	userHandler := users.NewUserHandler(userService)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.Handle("GET /me", users.AuthMiddleware(jwtSecret, http.HandlerFunc(userHandler.Me)))
	fmt.Println("Server Running..")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
