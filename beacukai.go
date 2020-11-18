package main

import "fmt"

// BeaCukai defines a bea cukai for cargo car checkin.
type BeaCukai struct {
	Name         string
	Code         string
	enterChannel chan *Car
	Ships        []*Ferry
}

// NewBeaCukai creates bea cukai.
func NewBeaCukai(ships []*Ferry) *BeaCukai {
	return &BeaCukai{
		Name:         "Bea Cukai",
		Code:         "C",
		enterChannel: make(chan *Car),
		Ships:        ships,
	}
}

// Open starts gas station listening for car queue.
func (bea *BeaCukai) Open() {
	for {
		car := <-bea.enterChannel // car entering bea cukai.

		label := "tertutup"
		if car.IsOpenCargo {
			label = "terbuka"
		}
		fmt.Printf("[%s]: pemeriksaan isi kargo: kargo %s\n", bea, label)

		bea.sendToFerry(car)
	}
}

func (bea *BeaCukai) String() string {
	return fmt.Sprintf("%s - %s", bea.Code, bea.Name)
}

// GetEnterChannel returns enter channel.
func (bea *BeaCukai) GetEnterChannel() chan<- *Car {
	return bea.enterChannel
}

// sendToFerry moves car to ferry based on their ticket issuer.
func (bea *BeaCukai) sendToFerry(car *Car) {
	for _, ferry := range bea.Ships {
		if ferry.Code == car.GetTicket().IssuedBy.Code {
			car.MoveTo(ferry)
			break
		}
	}
}
