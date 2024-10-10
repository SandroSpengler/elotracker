package handler

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/models"
	"github.com/sandrospengler/elotracker/views/home"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {

	ctx := context.Background()

	var summoners, err = models.Summoners().All(ctx, boil.GetContextDB())
	if err != nil {
		log.Fatal(err)
	}

	return render(c, home.Home(summoners))
}
