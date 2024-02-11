package main

import (
	"alvintanoto/palworld-ds-monitor/internal/dto"
	"alvintanoto/palworld-ds-monitor/internal/palrcon"
	"alvintanoto/palworld-ds-monitor/internal/util"
	"alvintanoto/palworld-ds-monitor/view"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (app *Application) SetupRoutes() {

	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &dto.HomepageDTO{
			OnlinePlayers: []dto.PlayerDTO{},
		}

		palrcon := palrcon.NewPalRcon(app.Configurations.RconUrl, app.Configurations.RconPassword)

		infoResp, err := palrcon.Info()
		if err != nil {
			log.Println(err.Error())
		}

		infoRespShort := strings.ReplaceAll(infoResp, "Welcome to Pal Server", "")
		rgx := regexp.MustCompile(`\[(.*?)\]`)
		data.ServerVersion = strings.ReplaceAll(string(rgx.Find([]byte(infoRespShort))), "[", "")
		data.ServerVersion = strings.ReplaceAll(data.ServerVersion, "]", "")

		data.ServerName = strings.ReplaceAll(infoRespShort, string(rgx.Find([]byte(infoRespShort))), "")

		playersResp, err := palrcon.ShowPlayers()
		if err != nil {
			log.Println(err.Error())
		}

		if len(playersResp) > 0 {
			playersStringSlice := strings.Split(playersResp, string(rune(10)))

			playersStringSlice = util.RemoveFromSlice(playersStringSlice, 0)
			playersStringSlice = util.RemoveFromSlice(playersStringSlice, len(playersStringSlice)-1)

			for _, playersString := range playersStringSlice {
				// split it by ,
				playerSlice := strings.Split(playersString, ",")
				player := dto.PlayerDTO{
					Name:      playerSlice[0],
					PlayerUID: playerSlice[1],
					SteamID:   playerSlice[2],
				}
				data.OnlinePlayers = append(data.OnlinePlayers, player)
			}
		}

		view.Homepage(data).Render(r.Context(), w)
	})

	app.Router.HandleFunc("/send-broadcast", func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		message = strings.TrimSpace(message)
		if message == "" {
			return
		}

		palrcon := palrcon.NewPalRcon(app.Configurations.RconUrl, app.Configurations.RconPassword)
		_, err := palrcon.Broadcast(message)
		if err != nil {
			log.Println(err.Error())
			view.FailedToast("Failed to send broadcast").Render(r.Context(), w)

			return
		}

		view.SuccessToast("Broadcast sent.").Render(r.Context(), w)
	})

	app.Router.HandleFunc("/shut-down", func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		message = strings.TrimSpace(message)
		if message == "" {
			log.Println("shut down server message is required")
			view.FailedToast("Shut down server message is required.").Render(r.Context(), w)

			return
		}

		durationString := r.PostFormValue("duration")
		_, err := strconv.Atoi(durationString)
		if err != nil {
			log.Println("shut down duration must be a number value")
			view.FailedToast("Shut down server duration number bust be a number value.").Render(r.Context(), w)

			return
		}

		palrcon := palrcon.NewPalRcon(app.Configurations.RconUrl, app.Configurations.RconPassword)

		_, err = palrcon.Shutdown(durationString, message)
		if err != nil {
			log.Println(err.Error())
			view.FailedToast("Failed to shut down server").Render(r.Context(), w)

			return
		}

		view.SuccessToast("Server successfully shut down.").Render(r.Context(), w)
	})

	app.Router.HandleFunc("/force-shut-down", func(w http.ResponseWriter, r *http.Request) {
		palrcon := palrcon.NewPalRcon(app.Configurations.RconUrl, app.Configurations.RconPassword)

		_, err := palrcon.DoExit()
		if err != nil {
			log.Println(err.Error())
			view.FailedToast("Failed to shut down server").Render(r.Context(), w)

			return
		}

		view.SuccessToast("Server successfully shut down.").Render(r.Context(), w)
	})

	app.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./view/assets/"))))
}
