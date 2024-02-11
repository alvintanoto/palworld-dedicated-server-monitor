package palrcon

import (
	"log"
	"strings"

	"github.com/gorcon/rcon"
)

const (
	PARAM_STEAM_ID     = "<SteamID>"
	PARAM_MESSAGE_TEXT = "<MessageText>"
	PARAM_SECONDS      = "<Seconds>"

	COMMAND_SHUT_DOWN          = "Shutdown <Seconds> <MessageText>"
	COMMAND_DO_EXIT            = "DoExit"
	COMMAND_BROADCAST          = "Broadcast<MessageText>"
	COMMAND_KICK_PLAYER        = "KickPlayer <SteamID>"
	COMMAND_BAN_PLAYER         = "BanPlayer <SteamID>"
	COMMAND_TELEPORT_TO_PLAYER = "TeleportToPlayer <SteamID>"
	COMMAND_TELEPORT_TO_ME     = "TeleportToMe <SteamID>"
	COMMAND_SHOW_PLAYERS       = "ShowPlayers"
	COMMAND_INFO               = "Info"
	COMMAND_SAVE               = "Save"
)

type prWrapper struct {
	Address  string
	Password string
}

type PalrconWrapper interface {
	Shutdown(seconds string, message string) (string, error)
	DoExit() (string, error)
	Broadcast(message string) (string, error)
	KickPlayer(steamID string) (string, error)
	BanPlayer(steamID string) (string, error)
	TeleportToPlayer(steamID string) (string, error)
	TeleportToMe(steamID string) (string, error)
	ShowPlayers() (string, error)
	Info() (string, error)
	SaveServer() (string, error)
}

func NewPalRcon(addr, password string) PalrconWrapper {
	return &prWrapper{
		Address:  addr,
		Password: password,
	}
}

func (w prWrapper) Shutdown(seconds string, message string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_SHUT_DOWN, PARAM_SECONDS, seconds)
	command = strings.ReplaceAll(command, PARAM_MESSAGE_TEXT, message)

	return conn.Execute(command)
}

func (w prWrapper) DoExit() (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	return conn.Execute(COMMAND_DO_EXIT)
}

func (w prWrapper) Broadcast(message string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_BROADCAST, PARAM_MESSAGE_TEXT, message)

	return conn.Execute(command)
}

func (w prWrapper) KickPlayer(steamID string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_KICK_PLAYER, PARAM_STEAM_ID, steamID)

	return conn.Execute(command)
}

func (w prWrapper) BanPlayer(steamID string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_BAN_PLAYER, PARAM_STEAM_ID, steamID)

	return conn.Execute(command)
}

func (w prWrapper) TeleportToPlayer(steamID string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_TELEPORT_TO_PLAYER, PARAM_STEAM_ID, steamID)

	return conn.Execute(command)
}

func (w prWrapper) TeleportToMe(steamID string) (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	command := strings.ReplaceAll(COMMAND_TELEPORT_TO_ME, PARAM_STEAM_ID, steamID)

	return conn.Execute(command)
}

func (w prWrapper) ShowPlayers() (resp string, err error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		return "", err
	}

	resp, err = conn.Execute(COMMAND_SHOW_PLAYERS)
	if err != nil {
		return "", err
	}
	conn.Close()
	return resp, err
}

func (w prWrapper) Info() (resp string, err error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		return "", err
	}

	resp, err = conn.Execute(COMMAND_INFO)
	if err != nil {
		return "", err
	}

	conn.Close()
	return resp, err
}

func (w prWrapper) SaveServer() (string, error) {
	conn, err := rcon.Dial(w.Address, w.Password)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	return conn.Execute(COMMAND_SAVE)
}
