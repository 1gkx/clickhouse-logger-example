package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const LogFile = "./logs/logrus.log"

func main() {

	logFile, err := os.OpenFile(LogFile, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	//log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.SetOutput(logFile)

	for i := 1; i <= 10 ; i++ {
		log.WithFields(log.Fields{
			"animal": "walrus",
			"count":   i,
		}).Info("A group of walrus emerges from the ocean")
	}
}
