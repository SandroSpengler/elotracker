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

var Socials = newSocialsTable("public", "socials", "")

type socialsTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnInteger
	PlayerName    postgres.ColumnString
	DiscordLink   postgres.ColumnString
	InstagramLink postgres.ColumnString
	TiktokLink    postgres.ColumnString
	TwitterLink   postgres.ColumnString
	TwitchLink    postgres.ColumnString
	YoutubeLink   postgres.ColumnString
	IconName      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SocialsTable struct {
	socialsTable

	EXCLUDED socialsTable
}

// AS creates new SocialsTable with assigned alias
func (a SocialsTable) AS(alias string) *SocialsTable {
	return newSocialsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SocialsTable with assigned schema name
func (a SocialsTable) FromSchema(schemaName string) *SocialsTable {
	return newSocialsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SocialsTable with assigned table prefix
func (a SocialsTable) WithPrefix(prefix string) *SocialsTable {
	return newSocialsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SocialsTable with assigned table suffix
func (a SocialsTable) WithSuffix(suffix string) *SocialsTable {
	return newSocialsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSocialsTable(schemaName, tableName, alias string) *SocialsTable {
	return &SocialsTable{
		socialsTable: newSocialsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newSocialsTableImpl("", "excluded", ""),
	}
}

func newSocialsTableImpl(schemaName, tableName, alias string) socialsTable {
	var (
		IDColumn            = postgres.IntegerColumn("id")
		PlayerNameColumn    = postgres.StringColumn("player_name")
		DiscordLinkColumn   = postgres.StringColumn("discord_link")
		InstagramLinkColumn = postgres.StringColumn("instagram_link")
		TiktokLinkColumn    = postgres.StringColumn("tiktok_link")
		TwitterLinkColumn   = postgres.StringColumn("twitter_link")
		TwitchLinkColumn    = postgres.StringColumn("twitch_link")
		YoutubeLinkColumn   = postgres.StringColumn("youtube_link")
		IconNameColumn      = postgres.StringColumn("icon_name")
		allColumns          = postgres.ColumnList{IDColumn, PlayerNameColumn, DiscordLinkColumn, InstagramLinkColumn, TiktokLinkColumn, TwitterLinkColumn, TwitchLinkColumn, YoutubeLinkColumn, IconNameColumn}
		mutableColumns      = postgres.ColumnList{PlayerNameColumn, DiscordLinkColumn, InstagramLinkColumn, TiktokLinkColumn, TwitterLinkColumn, TwitchLinkColumn, YoutubeLinkColumn, IconNameColumn}
	)

	return socialsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		PlayerName:    PlayerNameColumn,
		DiscordLink:   DiscordLinkColumn,
		InstagramLink: InstagramLinkColumn,
		TiktokLink:    TiktokLinkColumn,
		TwitterLink:   TwitterLinkColumn,
		TwitchLink:    TwitchLinkColumn,
		YoutubeLink:   YoutubeLinkColumn,
		IconName:      IconNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
