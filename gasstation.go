package main

import "fmt"

// GasDestination defines destination list of gas station.
type GasDestination struct {
	BeaCukai *BeaCukai
}

// GasStation defines gas station.
type GasStation struct {
	Name         string
	Code         string
	enterChannel chan *Car
	Destination  GasDestination

	Ships []*Ferry
}

// NewGasStation creates a new gas station.
func NewGasStation(ships []*Ferry, destination GasDestination) *GasStation {
	return &GasStation{
		Name:         "Pom Bensin",
		Code:         "P",
		Destination:  destination,
		enterChannel: make(chan *Car),
		Ships:        ships,
	}
}

// GetEnterChannel returns enterance channel.
func (gas *GasStation) GetEnterChannel() chan<- *Car {
	return gas.enterChannel
}

// String implements stringers.
func (gas *GasStation) String() string {
	return fmt.Sprintf("%s - %s", gas.Code, gas.Name)
}

// Open starts gas station listening for car queue.
func (gas *GasStation) Open() {
	for {
		car := <-gas.enterChannel // car entering gas station.

		fmt.Printf("\n[%s]: bahan bakar saat ini: %s\n", gas, car.FuelSummary())

		if car.PayloadType == PayloadFreight {
			car.MoveTo(gas.Destination.BeaCukai)
			continue
		}

		gas.sendToFerry(car)
	}
}

// sendToFerry moves car to ferry based on their ticket issuer.
func (gas *GasStation) sendToFerry(car *Car) {
	for _, ferry := range gas.Ships {
		if ferry.Code == car.GetTicket().IssuedBy.Code {
			car.MoveTo(ferry)
			break
		}
	}
}
