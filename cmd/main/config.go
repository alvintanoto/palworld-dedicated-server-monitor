package main

import (
	"log"
	"os"
	"strconv"
)

type Configurations struct {
	AppPort int

	RconUrl      string
	RconPassword string
}

// PORT=3000
// RCON_URL=rcon://192.168.1.148:25575
// RCON_PASSWORD=password

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

	return nil
}
