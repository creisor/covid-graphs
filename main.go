package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/creisor/covid-graphs/internal/data"
	"github.com/creisor/covid-graphs/internal/graphs"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	app := &cli.App{
		Name:  "covid-graphs",
		Usage: "Displays graphs using Johns Hopkins COVID-19 data",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Sets log level to DEBUG",
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Set listening port",
				Value:   "8081",
			},
			&cli.StringFlag{
				Name:    "data-directory",
				Aliases: []string{"d"},
				Usage:   "Set data directory",
				Value:   "./data",
			},
			&cli.BoolFlag{
				Name:    "no-pull",
				Aliases: []string{"n"},
				Usage:   "Does not pull the latest from the repo",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("verbose") {
				log.SetLevel(log.DebugLevel)
			}

			d := data.Data{c.String("data-directory")}
			if !exists(d.Directory) {
				d.Clone()
			} else {
				if !c.Bool("no-pull") {
					d.Pull()
				}
			}

			graphsRouter := mux.NewRouter()
			graphs := graphs.NewServer(graphsRouter)

			graphs.Routes()

			n := negroni.Classic()
			n.UseHandler(graphsRouter)

			n.Run(fmt.Sprintf(":%s", c.String("port")))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
