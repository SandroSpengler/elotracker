package handler

import (
	"fmt"
	"log"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/dtos"
	"github.com/sandrospengler/elotracker/views/home"

	"github.com/sandrospengler/elotracker/database"
	"github.com/sandrospengler/elotracker/models/elotracker/public/model"
	"github.com/sandrospengler/elotracker/models/elotracker/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type HomeHandler struct{}

type summonerRow struct {
	model.Summoner
	model.Socials

	Matches []struct {
		model.SummonerMatches
	}
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	var summonerDtos []dtos.SummonerDto

	summonerStmt := SELECT(
		table.Summoner.AllColumns,
		table.Socials.AllColumns,
		table.SummonerMatches.AllColumns,
	).FROM(
		table.Summoner.
			LEFT_JOIN(table.Socials, table.Socials.PlayerName.EQ(table.Summoner.PlayerName)).
			LEFT_JOIN(table.SummonerMatches, table.SummonerMatches.Puuid.EQ(table.Summoner.Puuid)),
	).WHERE(
		table.SummonerMatches.QueueID.EQ(String("420")),
	)

	dbRows := []summonerRow{}

	err := summonerStmt.Query(database.DB, &dbRows)
	if err != nil {
		log.Fatal(err)
	}

	dbRows = *sortByMostRecentMatch(&dbRows)

	for _, row := range dbRows {
		summonerDto := dtos.SummonerDto{}
		socialsDto := dtos.SocialsDto{}

		summonerDto.GameName = row.GameName
		summonerDto.TagLine = row.TagLine
		summonerDto.SummonerLevel = 0
		summonerDto.SumonerProfileIconUrl =
			"https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon0.jpg?image=q_auto,f_webp,w_auto&v=1710914129937"

		if row.SummonerLevel != nil {
			summonerDto.SummonerLevel = int64(*row.SummonerLevel)
		}

		if row.ProfileIconID != nil {
			profileIconId := int64(*row.ProfileIconID)

			summonerDto.SumonerProfileIconUrl =
				fmt.Sprintf("https://opgg-static.akamaized.net/meta/images/profile_icons/profileIcon%d.jpg?image=q_auto,f_webp,w_auto&v=1710914129937",
					profileIconId)
		}

		socialsDto.PlayerName = row.Socials.PlayerName
		socialsDto.IconName = DerefString(row.Socials.IconName)
		socialsDto.DiscordLink = DerefString(row.Socials.DiscordLink)
		socialsDto.InstagramLink = DerefString(row.Socials.InstagramLink)
		socialsDto.TiktokLink = DerefString(row.Socials.TiktokLink)
		socialsDto.TwitterLink = DerefString(row.Socials.TwitterLink)
		socialsDto.TwitchLink = DerefString(row.Socials.TwitchLink)
		socialsDto.YoutubeLink = DerefString(row.Socials.YoutubeLink)

		summonerDto.Socials = socialsDto

		summonerDtos = append(summonerDtos, summonerDto)
	}

	return render(c, home.Home(summonerDtos))
}

func sortByMostRecentMatch(rows *[]summonerRow) *[]summonerRow {
	// Sort matches by most recent match
	for _, row := range *rows {
		sort.Slice(row.Matches, func(i, j int) bool {
			gameEndTimeI := row.Matches[i].GameEndTime
			gameEndTimeJ := row.Matches[j].GameEndTime

			if gameEndTimeI == nil {
				return false
			}

			if gameEndTimeJ == nil {
				return false
			}

			return gameEndTimeI.Unix() > gameEndTimeJ.Unix()
		})

	}

	// Sort summoners by most recent match
	sort.Slice(*rows, func(i, j int) bool {
		gameEndTimeI := (*rows)[i].Matches[0].GameEndTime
		gameEndTimeJ := (*rows)[j].Matches[0].GameEndTime

		if gameEndTimeI == nil {
			return false
		}

		if gameEndTimeJ == nil {
			return false
		}

		return gameEndTimeI.Unix() > gameEndTimeJ.Unix()
	})

	return rows
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
