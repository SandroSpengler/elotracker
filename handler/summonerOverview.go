package handler

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/dtos"
	"github.com/sandrospengler/elotracker/views/overview"

	"github.com/sandrospengler/elotracker/database"
	"github.com/sandrospengler/elotracker/models/elotracker/public/model"
	"github.com/sandrospengler/elotracker/models/elotracker/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type SummonerOverviewHandler struct{}

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

type seasonRow struct {
	model.RankedSeason
	model.SummonerMatches
}

type rankedSeasonRow struct {
	model.RankedSeason
}

func (h SummonerOverviewHandler) HandleSummonerOverviewShow(c echo.Context) error {
	selectedSummonersQuery := c.QueryParam("selectedSummoners")
	selectedRankSeasonQuery, err := strconv.ParseInt(c.QueryParam("selectedRankedSeason"), 10, 32)
	if err != nil {
		selectedRankSeasonQuery = 0
	}

	rankedSeasonRows := []rankedSeasonRow{}

	var rankedSummonerDtos []dtos.SummonerDto
	var unrankedSummonerDtos []dtos.SummonerDto
	var playerNameDtos []dtos.PlayerNameDto

	selectedRankSeasonQuery = prepareRankedSeasonRows(rankedSeasonRows, selectedRankSeasonQuery)

	prepareSummonerDtos(
		&rankedSummonerDtos,
		&unrankedSummonerDtos,
		&playerNameDtos,
		selectedSummonersQuery,
		selectedRankSeasonQuery,
	)

	return render(c, overview.SummonerOverview(rankedSummonerDtos, unrankedSummonerDtos))
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

func prepareSummonerDtos(
	rankedSummonerDtos *[]dtos.SummonerDto,
	unrankedSummonerDtos *[]dtos.SummonerDto,
	playerNameDtos *[]dtos.PlayerNameDto,
	selectedSummonersQuery string,
	selectedRankSeasonQuery int64,
) {
	summonerRows := []summonerRow{}
	summonerRowsSelected := []summonerRow{}
	leagueRows := []leagueRow{}

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
		table.RankedSeason.AllColumns,
	).FROM(
		table.League.
			CROSS_JOIN(table.RankedSeason),
	).WHERE(
		table.League.QueueType.EQ(String("RANKED_SOLO_5x5")).AND(
			table.RankedSeason.Rid.EQ(Int(selectedRankSeasonQuery)),
		).AND(
			table.League.LastLeagueUpdate.BETWEEN(
				table.RankedSeason.StartDate, table.RankedSeason.EndDate,
			),
		),
	).ORDER_BY(
		table.League.LastLeagueUpdate.DESC(),
	)

	err := summonerStmt.Query(database.DB, &summonerRows)
	if err != nil {
		log.Fatal(err)
	}

	err = leagueStmt.Query(database.DB, &leagueRows)
	if err != nil {
		log.Fatal(err)
	}

	err = summonerStmt.Query(database.DB, &summonerRowsSelected)
	if err != nil {
		log.Fatal(err)
	}

	// filter by selected summoners
	if selectedSummonersQuery != "" {
		filteredRow := make([]summonerRow, 0)

		for i, row := range summonerRows {
			if strings.Contains(selectedSummonersQuery, row.Socials.PlayerName) {
				filteredRow = append(filteredRow, summonerRows[i])
			}
		}

		summonerRowsSelected = filteredRow
	}

	summonerRowsSelected = *sortByMostRecentMatch(&summonerRowsSelected)

	for _, row := range summonerRowsSelected {
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

			*rankedSummonerDtos = append(*rankedSummonerDtos, summonerDto)
		} else {
			*unrankedSummonerDtos = append(*unrankedSummonerDtos, summonerDto)
		}

		log.Println(len(*rankedSummonerDtos))
	}

	uniquePlayerNames := make(map[string]dtos.PlayerNameDto)

	for _, row := range summonerRows {
		if _, exist := uniquePlayerNames[row.Socials.PlayerName]; !exist {
			playerNameDto := dtos.PlayerNameDto{}

			playerNameDto.PlayerName = row.Socials.PlayerName
			playerNameDto.Selected = false

			if strings.Contains(selectedSummonersQuery, row.Socials.PlayerName) {
				playerNameDto.Selected = true
			}

			uniquePlayerNames[row.Socials.PlayerName] = playerNameDto
			*playerNameDtos = append(*playerNameDtos, playerNameDto)
		}
	}

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

func prepareRankedSeasonRows(
	rankedSeasonRows []rankedSeasonRow,
	selectedRankSeasonQuery int64,
) int64 {

	rankedSeasonStmt := SELECT(
		table.RankedSeason.Rid,
	).FROM(
		table.RankedSeason,
	).ORDER_BY(
		table.RankedSeason.Rid.DESC(),
	).LIMIT(1)

	err := rankedSeasonStmt.Query(database.DB, &rankedSeasonRows)
	if err != nil {
		log.Fatal(err)
	}

	if selectedRankSeasonQuery == 0 || selectedRankSeasonQuery > int64(rankedSeasonRows[0].Rid) {
		selectedRankSeasonQuery = int64(rankedSeasonRows[0].Rid)
	}

	return selectedRankSeasonQuery
}

func createRankedSeasonString(seasonId int32, splitId int32) string {
	return fmt.Sprintf("Season %d-%d", seasonId, splitId)
}
