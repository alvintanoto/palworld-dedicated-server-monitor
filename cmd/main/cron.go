package main

import (
	"alvintanoto/palworld-ds-monitor/internal/dto"
	"alvintanoto/palworld-ds-monitor/internal/palrcon"
	"alvintanoto/palworld-ds-monitor/internal/util"
	"log"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func (a *Application) NoOnlinePlayersShutDownCron() error {
	if a.Configurations.AutoShutDownDuration == 0 {
		if a.Scheduler != nil {
			a.Scheduler = nil
		}
		return nil
	}

	if a.Scheduler == nil {
		sch, err := gocron.NewScheduler()
		if err != nil {
			return err
		}

		a.Scheduler = sch
	}

	log.Println("create a job")
	job, err := a.Scheduler.NewJob(
		gocron.DurationJob(time.Duration(a.Configurations.AutoShutDownDuration)*time.Second),
		gocron.NewTask(
			func() {
				log.Println("running cronjob: ", time.Now())
				// shut down server
				palrcWrapper := palrcon.NewPalRcon(a.Configurations.RconUrl, a.Configurations.RconPassword)

				playersResp, err := palrcWrapper.ShowPlayers()
				if err != nil {
					log.Println(err.Error())
					a.Scheduler.Shutdown()
					return
				}

				var onlinePlayers []dto.PlayerDTO
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
						onlinePlayers = append(onlinePlayers, player)
					}
				}

				if len(onlinePlayers) > 0 {
					return
				}

				palrcWrapper.DoExit()
				a.Scheduler.Shutdown()
			},
		),
	)
	if err != nil {
		log.Println("failed to run a job:", err.Error())
		return err
	}

	log.Println("job created: ", job.ID())

	a.Scheduler.Start()
	log.Println("scheduler started")
	return nil
}
