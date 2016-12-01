package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/publish"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"net/http"
	"os"
)

var port string
var esUrl string

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
		cli.StringFlag{
			Name:        "elasticurl, e",
			Value:       "http://172.17.0.2:9200",
			Usage:       "elastic search URL",
			EnvVar:      "ELASTICSEARCH_URL",
			Destination: &esUrl,
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
	log.Info("esUrl: ", esUrl)
	router := mux.NewRouter().StrictSlash(true)
	publishHandler := publish.NewPublishHandler(esUrl)
	router.HandleFunc("/results", publishHandler.Handle).Methods("POST").HeadersRegexp("Content-Type", "application/(gzip|json)")
	router.HandleFunc("/events", publishHandler.HandleEvent).Methods("POST").HeadersRegexp("Content-Type", "application/json")
	log.Infof("Tugbot elasticsearch result service listening on port %s", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
