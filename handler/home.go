package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/dtos"
	"github.com/sandrospengler/elotracker/models"
	"github.com/sandrospengler/elotracker/views/home"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {

	ctx := context.Background()

	var summonerDtos []dtos.SummonerDto
	var summoners, err = models.Summoners().All(ctx, boil.GetContextDB())
	if err != nil {
		log.Fatal(err)
	}

	var summonerLevel = summoners[0].SummonerLevel
	s := strconv.FormatInt(summonerLevel.Int64, 10)

	log.Println(s)

	for _, summoner := range summoners {

		summonerDto := dtos.SummonerDto{}

		summonerDto.GameName = summoner.GameName
		summonerDto.TagLine = summoner.TagLine
		summonerDto.SummonerLevel = 0
		summonerDto.SumonerProfileIconUrl =
			"https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon0.jpg?image=q_auto,f_webp,w_auto&v=1710914129937"

		if summoner.SummonerLevel.Valid {
			summonerDto.SummonerLevel = summoner.SummonerLevel.Int64
		}

		if summoner.ProfileIconID.Valid {
			profileIconId := summoner.ProfileIconID.Int

			summonerDto.SumonerProfileIconUrl =
				fmt.Sprintf("https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon%d.jpg?image=q_auto,f_webp,w_auto&v=1710914129937", profileIconId)
		}

		summonerDtos = append(summonerDtos, summonerDto)
	}

	return render(c, home.Home(summonerDtos))
}
