package dtos

import "time"

type SummonerDto struct {
	GameName              string
	TagLine               string
	SummonerLevel         int64
	SumonerProfileIconUrl string
	HasRankedSolo5x5      bool
	Winrate               float32
	Socials               SocialsDto
	League                LeagueDto
}

type SocialsDto struct {
	IconName      string
	PlayerName    string
	DiscordLink   string
	InstagramLink string
	TiktokLink    string
	TwitterLink   string
	TwitchLink    string
	YoutubeLink   string
}

type LeagueDto struct {
	Tier             string
	Rank             string
	LeaguePoints     int32
	Wins             int32
	Losses           int32
	LastLeagueUpdate time.Time
}
