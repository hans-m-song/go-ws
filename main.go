package main

import (
	"go-ws/client"
	"go-ws/server"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Debug(os.Args)
	if len(os.Args) < 3 {
		log.Fatal("Not enough args, usage: go run . (server | client) port")
	}

	mode := os.Args[1]
	if mode != "server" && mode != "client" {
		log.WithFields(logrus.Fields{"mode": mode}).Fatal("invalid mode, should be \"server\" or \"client\"")
	}

	port := os.Args[2]

	log.WithFields(logrus.Fields{"mode": mode}).Info("starting")

	switch mode {
	case "client":
		client.Start(port)
	case "server":
		server.Start(port)
	}
}
