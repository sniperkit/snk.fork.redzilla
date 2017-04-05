package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/muka/redzilla/docker"
	"github.com/muka/redzilla/model"
	"github.com/muka/redzilla/service"
)

func main() {

	cfg := model.NewDefaultConfig()

	log.SetLevel(log.DebugLevel)

	err := service.Start(cfg)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		panic(err)
	}

}

func startContainer(cfg *model.Config) {

	name := "red3"

	err := docker.StartContainer(name, cfg)
	if err != nil {
		panic(err)
	}

	// log.Info("Ok, stopping now")
	// err = docker.StopContainer(name)
	// if err != nil {
	// 	panic(err)
	// }

	log.Info("Done")

}
