package main

import (
	"fmt"
	"github.com/JohnKucharsky/fiber_pgx_jwt/db"
	"github.com/JohnKucharsky/fiber_pgx_jwt/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}
	postgresURI := os.Getenv("POSTGRES_URI")
	redisURI := os.Getenv("REDIS_URI")

	f := fiber.New()
	f.Use(logger.New())
	f.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
				AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
			},
		),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = f.Shutdown()
	}()
	d := db.New(postgresURI)
	redis := db.NewRedis(redisURI)

	router.Register(f, d, redis)
	f.Static("/api/docs", "./docs")

	err := f.Listen(":8080")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Println("Running cleanup tasks...")
}
