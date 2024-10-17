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

type leagueRow struct {
	model.League
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

	leagueStmt := SELECT(
		table.League.AllColumns,
	).FROM(
		table.League,
	).WHERE(
		table.League.QueueType.EQ(String("RANKED_SOLO_5x5")).
			AND(table.League.LastLeagueUpdate.GT(Timestamp(2024, 9, 25, 8, 0, 0))),
	).ORDER_BY(
		table.League.LastLeagueUpdate.DESC(),
	)

	summonerRows := []summonerRow{}
	leagueRows := []leagueRow{}

	err := summonerStmt.Query(database.DB, &summonerRows)
	if err != nil {
		log.Fatal(err)
	}

	err = leagueStmt.Query(database.DB, &leagueRows)
	if err != nil {
		log.Fatal(err)
	}

	summonerRows = *sortByMostRecentMatch(&summonerRows)

	for _, row := range summonerRows {
		summonerDto := dtos.SummonerDto{}
		socialsDto := dtos.SocialsDto{}

		summonerLeagues := []dtos.LeagueDto{}

		summonerDto.GameName = row.GameName
		summonerDto.TagLine = row.TagLine
		summonerDto.HasRankedSolo5x5 = false
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

		for _, league := range leagueRows {
			if league.SummonerID == *row.Summoner.ID && league.QueueType == "RANKED_SOLO_5x5" {
				leagueDto := dtos.LeagueDto{}

				leagueDto.Tier = league.Tier
				leagueDto.Rank = league.Rank
				leagueDto.LeaguePoints = league.LeaguePoints
				leagueDto.Wins = league.Wins
				leagueDto.Losses = league.Losses
				leagueDto.LastLeagueUpdate = league.LastLeagueUpdate

				summonerLeagues = append(summonerLeagues, leagueDto)

			}
		}

		if len(summonerLeagues) > 0 {
			summonerDto.HasRankedSolo5x5 = true
			summonerDto.League = summonerLeagues[0]
			summonerDto.Winrate = calculateWinrate(summonerDto.League.Wins, summonerDto.League.Losses)
		}

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

func calculateWinrate(wins int32, losses int32) float32 {

	if wins == 0 && losses == 0 {
		return 0.0
	}

	if wins == 0 && losses != 0 {
		return 0.0
	}

	if wins != 0 && losses == 0 {
		return 0.0
	}

	return float32(wins) / float32(wins+losses)
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
