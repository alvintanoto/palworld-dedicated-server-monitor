package main

import (
	"log"
	"os"
	"strconv"
)

type Configurations struct {
	AppPort int

	AutoShutDownDuration int

	RconUrl      string
	RconPassword string
}

func (c *Configurations) ReadConfig() error {

	AppPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("app port must be specified")
	}

	c.AppPort = AppPort

	RconUrl := os.Getenv("RCON_URL")
	if RconUrl == "" {
		log.Fatal("rcon_url must be specified")
	}
	c.RconUrl = RconUrl

	RconPassword := os.Getenv("RCON_PASSWORD")
	c.RconPassword = RconPassword

	c.AutoShutDownDuration = 0
	return nil
}
