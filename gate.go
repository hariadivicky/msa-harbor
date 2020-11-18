package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// ErrNoFerryFound .
var ErrNoFerryFound = errors.New("no ferry that can accept this car")

// GateDestination defines destination list of gates.
type GateDestination struct {
	GasStation     *GasStation
	BeaCukai       *BeaCukai
	BatteryStation *BatteryStation
}

// Gate defines a starting point to harbor.
type Gate struct {
	Name         string
	Code         string
	Ships        []*Ferry
	Staffs       []*Staff
	enterChannel chan *Car

	Destination GateDestination
}

// NewGate creates a gate.
func NewGate(ships []*Ferry, staffs []*Staff, destination GateDestination) *Gate {
	return &Gate{
		Name:         "Gerbang Masuk",
		Code:         "A",
		Ships:        ships,
		Staffs:       staffs,
		Destination:  destination,
		enterChannel: make(chan *Car),
	}
}

// Open starts gate listening for car queue.
func (gate *Gate) Open() {
	for {
		car := <-gate.enterChannel // car entering gate.
		fmt.Print("tekan CTRL + C untuk menampilkan laporan\n")
		fmt.Printf("[%s] mobil masuk: \n%s\n", gate, car)

		if err := gate.findFerryFor(car); err != nil {
			log.Println(err)
			continue
		}

		staff := gate.getParkingStaff()
		staff.Park(car)

		// the car must be fueled if current fuel is less than 10%.
		if car.FuelPercentage() < 10 {
			switch car.FuelType {
			case FuelPetrol:
				car.MoveTo(gate.Destination.GasStation)
			case FuelBattery:
				car.MoveTo(gate.Destination.BatteryStation)
			}

			continue
		}

		// all cargo cars must be send to bea cukai.
		if car.PayloadType == PayloadFreight {
			car.MoveTo(gate.Destination.BeaCukai)
			continue
		}

		gate.sendToFerry(car)
	}
}

func (gate *Gate) String() string {
	return fmt.Sprintf("%s - %s", gate.Code, gate.Name)
}

// IsStartingPoint returns true because gate was starting point.
func (gate *Gate) IsStartingPoint() bool {
	return true
}

// GetEnterChannel returns enter channel.
func (gate *Gate) GetEnterChannel() chan<- *Car {
	return gate.enterChannel
}

// findFerryFor finds a fit ferry for this car.
// and then issue a new ticket for this car.
func (gate *Gate) findFerryFor(car *Car) error {
	for _, ferry := range gate.Ships {

		if ferry.CanAccept(car) {
			ferry.IssueTicketFor(car)
			return nil
		}
	}

	return ErrNoFerryFound
}

// getParkingStaff returns random staff.
func (gate *Gate) getParkingStaff() *Staff {
	i := rand.Intn(len(gate.Staffs))

	return gate.Staffs[i]
}

// sendToFerry moves car to ferry based on their ticket issuer.
func (gate *Gate) sendToFerry(car *Car) {
	for _, ferry := range gate.Ships {
		if ferry.Code == car.GetTicket().IssuedBy.Code {
			log.Println(ferry.Name)
			car.MoveTo(ferry)
			return
		}
	}
}
