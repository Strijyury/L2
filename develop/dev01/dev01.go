package main

import (
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	// STDERR
	logger := log.New(os.Stderr, "", 0)

	// Querying the current time
	curTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}
	log.Printf("The current time is %s", curTime)

	// Querying time metadata
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}
	metaTime := time.Now().Add(response.ClockOffset)
	log.Printf("The time metadata is %s", metaTime)
}
