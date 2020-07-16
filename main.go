package main

import (
	"flag"
	"os"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/app"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/config/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var port string

func init() {
	flag.StringVar(&port, "port", "6000", "Assigning the port that the server should listen on.")

	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting Server.")

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}
func main() {

	s := app.NewServer()

	s.Init(port)
	s.StartApplication()
}
