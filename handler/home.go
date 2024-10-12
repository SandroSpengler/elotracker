package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/dtos"
	"github.com/sandrospengler/elotracker/views/home"

	"github.com/sandrospengler/elotracker/models/elotracker/public/model"
	"github.com/sandrospengler/elotracker/models/elotracker/public/table"
	"github.com/sandrospengler/elotracker/database"

	. "github.com/go-jet/jet/v2/postgres"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {

	var summonerDtos []dtos.SummonerDto

	stmt := SELECT(table.Summoner.AllColumns).FROM(
		table.Summoner.
			INNER_JOIN(table.Socials, table.Socials.PlayerName.EQ(table.Summoner.PlayerName)),
	)

	var dest []struct {
		model.Summoner
		model.Socials
	}

	err := stmt.Query(database.DB, &dest)
	if err != nil {
		log.Fatal(err)
	}

	for _, summoner := range dest {

		summonerDto := dtos.SummonerDto{}

		summonerDto.GameName = summoner.GameName
		summonerDto.TagLine = summoner.TagLine
		summonerDto.SummonerLevel = 0
		summonerDto.SumonerProfileIconUrl =
			"https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon0.jpg?image=q_auto,f_webp,w_auto&v=1710914129937"

		if summoner.SummonerLevel != nil {
			summonerDto.SummonerLevel = int64(*summoner.SummonerLevel)
		}

		if summoner.ProfileIconID != nil {
			profileIconId := int64(*summoner.ProfileIconID)

			summonerDto.SumonerProfileIconUrl =
				fmt.Sprintf("https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon%d.jpg?image=q_auto,f_webp,w_auto&v=1710914129937", profileIconId)
		}

		summonerDtos = append(summonerDtos, summonerDto)
	}

	return render(c, home.Home(summonerDtos))
}
