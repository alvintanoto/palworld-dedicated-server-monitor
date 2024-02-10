package dto

type PlayerDTO struct {
	Name      string
	PlayerUID string
	SteamID   string
}

type HomepageDTO struct {
	ServerVersion string
	ServerName    string
	OnlinePlayers []PlayerDTO
}
