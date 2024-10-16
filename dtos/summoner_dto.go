package dtos

type SummonerDto struct {
	GameName              string
	TagLine               string
	SummonerLevel         int64
	SumonerProfileIconUrl string
	Socials               SocialsDto
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
