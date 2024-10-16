//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var League = newLeagueTable("public", "league", "")

type leagueTable struct {
	postgres.Table

	// Columns
	LeagueID         postgres.ColumnString
	QueueType        postgres.ColumnString
	Tier             postgres.ColumnString
	Rank             postgres.ColumnString
	SummonerID       postgres.ColumnString
	LeaguePoints     postgres.ColumnInteger
	Wins             postgres.ColumnInteger
	Losses           postgres.ColumnInteger
	Veteran          postgres.ColumnBool
	Inactive         postgres.ColumnBool
	Freshblood       postgres.ColumnBool
	Hotstreak        postgres.ColumnBool
	LastLeagueUpdate postgres.ColumnTimestamp
	ID               postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LeagueTable struct {
	leagueTable

	EXCLUDED leagueTable
}

// AS creates new LeagueTable with assigned alias
func (a LeagueTable) AS(alias string) *LeagueTable {
	return newLeagueTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LeagueTable with assigned schema name
func (a LeagueTable) FromSchema(schemaName string) *LeagueTable {
	return newLeagueTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LeagueTable with assigned table prefix
func (a LeagueTable) WithPrefix(prefix string) *LeagueTable {
	return newLeagueTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LeagueTable with assigned table suffix
func (a LeagueTable) WithSuffix(suffix string) *LeagueTable {
	return newLeagueTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLeagueTable(schemaName, tableName, alias string) *LeagueTable {
	return &LeagueTable{
		leagueTable: newLeagueTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newLeagueTableImpl("", "excluded", ""),
	}
}

func newLeagueTableImpl(schemaName, tableName, alias string) leagueTable {
	var (
		LeagueIDColumn         = postgres.StringColumn("league_id")
		QueueTypeColumn        = postgres.StringColumn("queue_type")
		TierColumn             = postgres.StringColumn("tier")
		RankColumn             = postgres.StringColumn("rank")
		SummonerIDColumn       = postgres.StringColumn("summoner_id")
		LeaguePointsColumn     = postgres.IntegerColumn("league_points")
		WinsColumn             = postgres.IntegerColumn("wins")
		LossesColumn           = postgres.IntegerColumn("losses")
		VeteranColumn          = postgres.BoolColumn("veteran")
		InactiveColumn         = postgres.BoolColumn("inactive")
		FreshbloodColumn       = postgres.BoolColumn("freshblood")
		HotstreakColumn        = postgres.BoolColumn("hotstreak")
		LastLeagueUpdateColumn = postgres.TimestampColumn("last_league_update")
		IDColumn               = postgres.IntegerColumn("id")
		allColumns             = postgres.ColumnList{LeagueIDColumn, QueueTypeColumn, TierColumn, RankColumn, SummonerIDColumn, LeaguePointsColumn, WinsColumn, LossesColumn, VeteranColumn, InactiveColumn, FreshbloodColumn, HotstreakColumn, LastLeagueUpdateColumn, IDColumn}
		mutableColumns         = postgres.ColumnList{LeagueIDColumn, QueueTypeColumn, TierColumn, RankColumn, SummonerIDColumn, LeaguePointsColumn, WinsColumn, LossesColumn, VeteranColumn, InactiveColumn, FreshbloodColumn, HotstreakColumn, LastLeagueUpdateColumn}
	)

	return leagueTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		LeagueID:         LeagueIDColumn,
		QueueType:        QueueTypeColumn,
		Tier:             TierColumn,
		Rank:             RankColumn,
		SummonerID:       SummonerIDColumn,
		LeaguePoints:     LeaguePointsColumn,
		Wins:             WinsColumn,
		Losses:           LossesColumn,
		Veteran:          VeteranColumn,
		Inactive:         InactiveColumn,
		Freshblood:       FreshbloodColumn,
		Hotstreak:        HotstreakColumn,
		LastLeagueUpdate: LastLeagueUpdateColumn,
		ID:               IDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
