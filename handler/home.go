package handler

import (
	"log"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sandrospengler/elotracker/database"
	"github.com/sandrospengler/elotracker/dtos"
	"github.com/sandrospengler/elotracker/models/elotracker/public/table"
	"github.com/sandrospengler/elotracker/views/home"

	. "github.com/go-jet/jet/v2/postgres"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	selectedSummonersQuery := c.QueryParam("selectedSummoners")
	selectedRankSeasonQuery, err := strconv.ParseInt(c.QueryParam("selectedRankedSeason"), 10, 32)
	if err != nil {
		selectedRankSeasonQuery = 0
	}

	rankedSeasonRows := []rankedSeasonRow{}
	seasonRows := []seasonRow{}

	selectedRankSeasonQuery = prepareRankedSeasonRows(rankedSeasonRows, selectedRankSeasonQuery)

	var rankedSummonerDtos []dtos.SummonerDto
	var unrankedSummonerDtos []dtos.SummonerDto
	var playerNameDtos []dtos.PlayerNameDto
	var seasonDtos []dtos.SeasonDto

	prepareSummonerDtos(
		&rankedSummonerDtos,
		&unrankedSummonerDtos,
		&playerNameDtos,
		selectedSummonersQuery,
		selectedRankSeasonQuery,
	)

	seasonStmt := SELECT(
		table.RankedSeason.AllColumns,
		table.SummonerMatches.AllColumns,
	).FROM(
		table.RankedSeason.
			LEFT_JOIN(table.SummonerMatches, table.SummonerMatches.GameEndTime.
				BETWEEN(table.RankedSeason.StartDate, table.RankedSeason.EndDate)),
	).WHERE(
		table.SummonerMatches.MatchID.IS_NOT_NULL(),
	)

	err = seasonStmt.Query(database.DB, &seasonRows)
	if err != nil {
		log.Fatal(err)
	}

	uniqueSeasons := make(map[int32]dtos.SeasonDto)

	for _, row := range seasonRows {
		if _, exist := uniqueSeasons[row.Rid]; !exist {
			seasonDto := dtos.SeasonDto{}

			seasonDto.Rid = row.Rid
			seasonDto.SeasonId = row.RankedSeasonID
			seasonDto.SplitId = row.SplitID
			seasonDto.Selected = false
			seasonDto.RankedSeasonString = createRankedSeasonString(seasonDto.SeasonId, seasonDto.SplitId)

			if int32(selectedRankSeasonQuery) == seasonDto.Rid {
				seasonDto.Selected = true
			}

			uniqueSeasons[row.Rid] = seasonDto
			seasonDtos = append(seasonDtos, seasonDto)
		}
	}

	sort.Slice(seasonDtos, func(i, j int) bool {
		return seasonDtos[i].Rid > seasonDtos[j].Rid
	})

	return render(c, home.Home(rankedSummonerDtos, unrankedSummonerDtos, playerNameDtos, seasonDtos))
}
