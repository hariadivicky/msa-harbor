package main

import (
	"errors"
	"fmt"
)

// ErrStartingPointNotFound thrown when user doesn't provide a starting point.
var ErrStartingPointNotFound = errors.New("no starting point found, please see StartingPoint contracts")

// Harbor defines a harbor.
type Harbor struct {
	Ships  []*Ferry
	People []*Staff
	Places []Place
}

func newAsset() ([]*Ferry, []*Staff) {
	ships := []*Ferry{
		NewSmallFerry(),
		NewLargeFerry(),
		NewEcoFerry(),
	}

	people := []*Staff{NewStaff("hariadi", 10), NewStaff("vicky", 11)}

	return ships, people
}

// NewHarbor creates a new harbor.
func NewHarbor() *Harbor {
	ships, people := newAsset()

	beaCukai := NewBeaCukai(ships)

	gasDestination := GasDestination{BeaCukai: beaCukai}
	gasStation := NewGasStation(ships, gasDestination)

	batteryStation := NewBatteryStation(ships)

	gateDestination := GateDestination{
		BeaCukai:       beaCukai,
		GasStation:     gasStation,
		BatteryStation: batteryStation,
	}
	gate := NewGate(ships, people, gateDestination)
	places := []Place{gate, gasStation, batteryStation, beaCukai}

	return &Harbor{ships, people, places}
}

// Open will open all places and return starting point for car enterance.
func (h *Harbor) Open() (StartingPoint, error) {
	// makes places accept car queue.
	for _, place := range h.Places {
		go place.Open()
	}

	// makes ferry accept car queue.
	for _, ferry := range h.Ships {
		go ferry.Open()
	}

	startingPoint, err := h.findStartingPoint()
	if err != nil {
		return nil, err
	}

	return startingPoint, nil
}

func (h *Harbor) findStartingPoint() (StartingPoint, error) {
	for _, place := range h.Places {
		if startable, ok := place.(StartingPoint); ok && startable.IsStartingPoint() {
			return startable, nil
		}
	}

	return nil, ErrStartingPointNotFound
}

// GenerateReport prints staff & ferry report.
func (h *Harbor) GenerateReport() {
	var totalRevenue float64

	fmt.Print("\n[Laporan Pendapatan Tiket]\n")
	for _, ferry := range h.Ships {
		totalRevenue = totalRevenue + ferry.GetTotalRevenue()
		ferry.GenerateReport()
	}

	fmt.Printf("\nTOTAL SELURUH PENDAPATAN KAPAL=Rp. %.2f\n", totalRevenue)
}
