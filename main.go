package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var interval int
	flag.IntVar(&interval, "interval", 5, "interval car enterance in second, default 5")
	flag.Parse()

	harbor := NewHarbor()
	gate, err := harbor.Open()
	if err != nil {
		log.Fatal(err)
	}

	go simulateTraffic(gate, time.Duration(interval)*time.Second)

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown // print report on shutdown
	harbor.GenerateReport()
}

func simulateTraffic(startingPoint StartingPoint, interval time.Duration) {
	for {
		NewRandomCar().MoveTo(startingPoint)
		<-time.After(interval)
	}
}
