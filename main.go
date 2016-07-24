package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"github.com/gaia-docker/tugbot-result-service-es/publisher"
)

var port string

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port, p",
			Value:       "8081",
			Usage:       "http service port",
			Destination: &port,
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug mode with verbose logging",
		},
	}
	app.Name = "tugbot-result-service-es"
	app.Usage = "Publish test results to elasticsearch."
	app.Before = before
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Errorf("Exit (%+v)", err)
	}
}

func before(c *cli.Context) error {

	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	return nil
}

func start(c *cli.Context) error {

	log.Info("Starting tugbot elasticsearch result service...")
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/results", publisher.Handle).Methods("POST").HeadersRegexp("Content-Type", "application/(gzip|json)")
	log.Infof("Tugbot elasticsearch result service listening on port %s", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
