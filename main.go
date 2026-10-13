package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/engine"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/gamestate"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/users"
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
	if err := db.AutoMigrate(&fighters.Fighter{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
	if err := db.AutoMigrate(&promotions.Promotion{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
	if err := db.AutoMigrate(&gamestate.GameState{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo, jwtSecret)
	userHandler := users.NewUserHandler(userService)
	fighterRepo := fighters.NewFighterRepository(db)
	fighterService := fighters.NewFighterService(fighterRepo)
	fh := fighters.NewFighterHandler(fighterService)
	gr := gamestate.NewRepository(db)
	gs := gamestate.NewService(gr)
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 1))
	eng := engine.New(gs, fighterService, rng)
	eh := engine.NewHandler(eng)
	pr := promotions.NewPromotionRepository(db)
	ps := promotions.NewPromotionService(pr, eng)
	ph := promotions.NewPromotionHandler(ps)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.Handle("GET /me", users.AuthMiddleware(jwtSecret, http.HandlerFunc(userHandler.Me)))
	mux.Handle("POST /fighters", users.AuthMiddleware(jwtSecret, http.HandlerFunc(fh.CreateFighter)))
	mux.Handle("GET /fighters/{id}", users.AuthMiddleware(jwtSecret, http.HandlerFunc(fh.GetFighter)))
	mux.Handle("DELETE /fighters/{id}", users.AuthMiddleware(jwtSecret, http.HandlerFunc(fh.DeleteFighter)))
	mux.Handle("POST /promotions", users.AuthMiddleware(jwtSecret, http.HandlerFunc(ph.CreatePromotion)))
	mux.Handle("GET /promotions/me", users.AuthMiddleware(jwtSecret, http.HandlerFunc(ph.GetMyPromotion)))
	mux.Handle("POST /advance", users.AuthMiddleware(jwtSecret, http.HandlerFunc(eh.Advance)))
	fmt.Println("Server Running..")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
