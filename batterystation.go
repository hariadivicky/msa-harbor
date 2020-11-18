package main

import (
	"fmt"
)

// BatteryStation defines battery station.
type BatteryStation struct {
	Name         string
	Code         string
	enterChannel chan *Car

	Ships []*Ferry
}

// NewBatteryStation creates a new battery station.
func NewBatteryStation(ships []*Ferry) *BatteryStation {
	return &BatteryStation{
		Name:         "Station pengisian baterei",
		Code:         "B",
		enterChannel: make(chan *Car),
		Ships:        ships,
	}
}

// GetEnterChannel returns enterance channel.
func (battery *BatteryStation) GetEnterChannel() chan<- *Car {
	return battery.enterChannel
}

// String implements stringers.
func (battery *BatteryStation) String() string {
	return fmt.Sprintf("%s - %s", battery.Code, battery.Name)
}

// Open starts battery station listening for car queue.
func (battery *BatteryStation) Open() {
	for {
		car := <-battery.enterChannel // car entering battery station.

		fmt.Printf("\n[%s]: bahan bakar saat ini: %s\n", battery, car.FuelSummary())

		battery.sendToFerry(car)
	}
}

// sendToFerry moves car to ferry based on their ticket issuer.
func (battery *BatteryStation) sendToFerry(car *Car) {
	for _, ferry := range battery.Ships {
		if ferry.Code == car.GetTicket().IssuedBy.Code {
			car.MoveTo(ferry)
			break
		}
	}
}
