package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"github.com/sandrospengler/elotracker/database"
	"github.com/sandrospengler/elotracker/handler"
)

func main() {
	environment := os.Getenv("environment")

	if environment != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	connectionString := os.Getenv("DATABASE_URL")

	if connectionString == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	database.Connect(connectionString)

	e := echo.New()
	e.Use(middleware.Logger())

	homeHandler := handler.HomeHandler{}
	summonerOverviewHandler := handler.SummonerOverviewHandler{}

	e.GET("/", homeHandler.HandleHomeShow)
	e.GET("/summoner-overview", summonerOverviewHandler.HandleSummonerOverviewShow)
	e.Static("/assets", "assets")

	e.Start(":5555")
}
