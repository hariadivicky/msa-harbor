package main

import (
	"fmt"
	"math/rand"
)

type (
	// CarSize .
	CarSize int
	// CarType .
	CarType int
	// CarKind .
	CarKind int
	// PayloadType .
	PayloadType int
	// FuelType .
	FuelType int
)

var (
	carTypes     = []string{"kargo", "penumpang"}
	payloadTypes = []string{"penumpang", "barang"}
	carSizes     = []string{"kecil", "besar"}
	carKinds     = []string{"car", "pickup", "bus", "truck", "eco"}
	fuelTypes    = []string{"minyak", "listrik"}
)

const (
	// SmallCar .
	SmallCar CarSize = iota
	// LargeCar .
	LargeCar
)

const (
	// CargoType .
	CargoType CarType = iota
	// NonCargoType .
	NonCargoType
)

const (
	// PayloadHuman .
	PayloadHuman PayloadType = iota
	// PayloadFreight .
	PayloadFreight
)

const (
	// KindCar .
	KindCar CarKind = iota
	// KindPickup .
	KindPickup
	// KindBus .
	KindBus
	// KindTruck .
	KindTruck
	// KindEco .
	KindEco
)

const (
	// FuelPetrol .
	FuelPetrol FuelType = iota
	// FuelBattery .
	FuelBattery
)

func (enum CarType) String() string {
	return carTypes[enum]
}

func (enum CarSize) String() string {
	return carSizes[enum]
}

func (f PayloadType) String() string {
	return payloadTypes[f]
}

func (enum CarKind) String() string {
	return carKinds[enum]
}

func (enum FuelType) String() string {
	return fuelTypes[enum]
}

// Car defines a car.
type Car struct {
	Kind         CarKind
	Size         CarSize
	PayloadType  PayloadType
	IsOpenCargo  bool
	FuelType     FuelType
	FuelCapacity float32
	CurrentFuel  float32
	ticket       Ticket
	takenRoutes  []string
}

// NewCar creates a new car.
func NewCar(
	kind CarKind,
	size CarSize,
	payloadType PayloadType,
	isOpenCargo bool,
	fuelType FuelType,
	currentFuel,
	fuelCapacity float32,
) *Car {
	return &Car{
		Kind:         kind,
		Size:         size,
		PayloadType:  payloadType,
		IsOpenCargo:  isOpenCargo,
		CurrentFuel:  currentFuel,
		FuelType:     fuelType,
		FuelCapacity: fuelCapacity,
		takenRoutes:  make([]string, 0),
	}
}

// NewSmallCar create a new small/city/family car.
func NewSmallCar(currentFuel float32) *Car {
	return NewCar(KindCar, SmallCar, PayloadHuman, false, FuelPetrol, currentFuel, 80.0)
}

// NewPickup create a new pickup.
func NewPickup(currentFuel float32) *Car {
	return NewCar(KindPickup, SmallCar, PayloadFreight, true, FuelPetrol, currentFuel, 90.0)
}

// NewBus create a new bus.
func NewBus(currentFuel float32) *Car {
	return NewCar(KindBus, LargeCar, PayloadHuman, false, FuelPetrol, currentFuel, 100.0)
}

// NewTruck create a new truck.
func NewTruck(currentFuel float32) *Car {
	return NewCar(KindTruck, LargeCar, PayloadFreight, false, FuelPetrol, currentFuel, 110.0)
}

// NewEcoCar create a new truck.
func NewEcoCar(currentFuel float32) *Car {
	return NewCar(KindEco, SmallCar, PayloadHuman, false, FuelBattery, currentFuel, 80.0)
}

var availableRandomCars = []string{"car", "truck", "pickup", "bus", "eco"}

// NewRandomCar returns random car type.
func NewRandomCar() *Car {
	randCarType := rand.Intn(len(availableRandomCars))
	randFuel := float32(rand.Intn(20))

	switch availableRandomCars[randCarType] {
	case "truck":
		return NewTruck(randFuel)

	case "bus":
		return NewBus(randFuel)

	case "pickup":
		return NewPickup(randFuel)

	case "eco":
		return NewEcoCar(randFuel)

	default:
		return NewSmallCar(randFuel)
	}
}

// SetTicket sets enterance ticket.
func (car *Car) SetTicket(ticket Ticket) {
	car.ticket = ticket
}

// GetTicket returns ticket.
func (car *Car) GetTicket() Ticket {
	return car.ticket
}

func (car *Car) String() string {
	return fmt.Sprintf(
		"jenis=%s \nukuran=%s \nmuatan=%s \nsisa_bahan_bakar=%.1f%% \njenis_bahan_bakar=%s",
		car.Kind, car.Size, car.PayloadType, car.FuelPercentage(), car.FuelType,
	)
}

// FuelPercentage returns percentage of current fuel.
func (car *Car) FuelPercentage() float32 {
	return car.CurrentFuel / car.FuelCapacity * 100
}

// MoveTo moves car to target destination.
func (car *Car) MoveTo(destination Visitable) {
	destination.GetEnterChannel() <- car
	car.takenRoutes = append(car.takenRoutes, destination.String())
}

// FuelSummary returns fuel summary in string.
func (car *Car) FuelSummary() string {
	return fmt.Sprintf(
		"%.1fL dari %.1fL (%.1f%%)",
		car.CurrentFuel, car.FuelCapacity, car.FuelPercentage(),
	)
}

// PrintRoutes displays taken route.
func (car *Car) PrintRoutes() {
	for _, route := range car.takenRoutes {
		fmt.Printf("|\n%s\n", route)
	}
	fmt.Print("--------------------------------------------------------\n\n")
}
