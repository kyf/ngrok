package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
	"ngrok/client"
)

type Application struct {
	Name        string
	DisplayName string
}

func (a Application) Start(s service.Service) error {
	go a.run(s)
	return nil
}

func (a Application) run(s service.Service) {
	client.Main()
}

func (a Application) Stop(s service.Service) error {
	return nil
}

func main() {
	app := Application{"ngrok", "ngrok service"}
	s, err := service.New(app, &service.Config{
		Name:        app.Name,
		DisplayName: app.DisplayName,
	})
	if err != nil {
		log.Printf("create [%s] failure, error is %v", app.DisplayName, err)
		return
	}

	control := ""
	if len(os.Args) > 1 {
		control = os.Args[1]
	}

	if len(control) > 0 {
		err := service.Control(s, control)
		if err != nil {
			log.Printf("[%s] [%s] failure, error is %v", app.DisplayName, control, err)
		}
		return
	}

	logger, _ := s.Logger(nil)
	err = s.Run()
	if err != nil {
		logger.Errorf("[%s] Run failure, error is %v", app.DisplayName, err)
	}
}
