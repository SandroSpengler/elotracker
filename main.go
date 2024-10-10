package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq"

	"github.com/sandrospengler/elotracker/handler"
)

var DB *sql.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")

	if connectionString == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db := connectDB(connectionString)

	boil.SetDB(db)

	e := echo.New()
	e.Use(middleware.Logger())

	userHandler := handler.HomeHandler{}

	e.GET("/", userHandler.HandleHomeShow)
	e.Static("/assets", "assets")

	e.Start(":5555")
}

func connectDB(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
