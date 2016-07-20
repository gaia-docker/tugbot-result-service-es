package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var port string

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port, p",
			Value:       "8081",
			Usage:       "http service port",
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "loglevel, l",
			Value:       log.DebugLevel,
			Usage:       "log level",
			Destination: &log.DebugLevel,
		},
	}
	app.Name = "tugbot-result-service-es"
	app.Usage = "Publish test results to elasticsearch."
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Errorf("Exit (%+v)", err)
	}
}

func start(c *cli.Context) error {
	return nil
}
