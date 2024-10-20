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

var Summoner = newSummonerTable("public", "summoner", "")

type summonerTable struct {
	postgres.Table

	// Columns
	Puuid              postgres.ColumnString
	GameName           postgres.ColumnString
	TagLine            postgres.ColumnString
	LastMatchUpdate    postgres.ColumnTimestamp
	LastSummonerUpdate postgres.ColumnTimestamp
	ID                 postgres.ColumnString
	AccountID          postgres.ColumnString
	ProfileIconID      postgres.ColumnInteger
	SummonerLevel      postgres.ColumnInteger
	RevisionDate       postgres.ColumnTimestamp
	PlayerName         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SummonerTable struct {
	summonerTable

	EXCLUDED summonerTable
}

// AS creates new SummonerTable with assigned alias
func (a SummonerTable) AS(alias string) *SummonerTable {
	return newSummonerTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SummonerTable with assigned schema name
func (a SummonerTable) FromSchema(schemaName string) *SummonerTable {
	return newSummonerTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SummonerTable with assigned table prefix
func (a SummonerTable) WithPrefix(prefix string) *SummonerTable {
	return newSummonerTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SummonerTable with assigned table suffix
func (a SummonerTable) WithSuffix(suffix string) *SummonerTable {
	return newSummonerTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSummonerTable(schemaName, tableName, alias string) *SummonerTable {
	return &SummonerTable{
		summonerTable: newSummonerTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newSummonerTableImpl("", "excluded", ""),
	}
}

func newSummonerTableImpl(schemaName, tableName, alias string) summonerTable {
	var (
		PuuidColumn              = postgres.StringColumn("puuid")
		GameNameColumn           = postgres.StringColumn("game_name")
		TagLineColumn            = postgres.StringColumn("tag_line")
		LastMatchUpdateColumn    = postgres.TimestampColumn("last_match_update")
		LastSummonerUpdateColumn = postgres.TimestampColumn("last_summoner_update")
		IDColumn                 = postgres.StringColumn("id")
		AccountIDColumn          = postgres.StringColumn("account_id")
		ProfileIconIDColumn      = postgres.IntegerColumn("profile_icon_id")
		SummonerLevelColumn      = postgres.IntegerColumn("summoner_level")
		RevisionDateColumn       = postgres.TimestampColumn("revision_date")
		PlayerNameColumn         = postgres.StringColumn("player_name")
		allColumns               = postgres.ColumnList{PuuidColumn, GameNameColumn, TagLineColumn, LastMatchUpdateColumn, LastSummonerUpdateColumn, IDColumn, AccountIDColumn, ProfileIconIDColumn, SummonerLevelColumn, RevisionDateColumn, PlayerNameColumn}
		mutableColumns           = postgres.ColumnList{GameNameColumn, TagLineColumn, LastMatchUpdateColumn, LastSummonerUpdateColumn, IDColumn, AccountIDColumn, ProfileIconIDColumn, SummonerLevelColumn, RevisionDateColumn, PlayerNameColumn}
	)

	return summonerTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Puuid:              PuuidColumn,
		GameName:           GameNameColumn,
		TagLine:            TagLineColumn,
		LastMatchUpdate:    LastMatchUpdateColumn,
		LastSummonerUpdate: LastSummonerUpdateColumn,
		ID:                 IDColumn,
		AccountID:          AccountIDColumn,
		ProfileIconID:      ProfileIconIDColumn,
		SummonerLevel:      SummonerLevelColumn,
		RevisionDate:       RevisionDateColumn,
		PlayerName:         PlayerNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
