package dto

type PlayerDTO struct {
	Name      string
	PlayerUID string
	SteamID   string
}

type HomepageDTO struct {
	Error                        string
	ServerVersion                string
	ServerName                   string
	NoPlayerAutoShutdownDuration int
	OnlinePlayers                []PlayerDTO
}
