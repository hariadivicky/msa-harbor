package main

import "fmt"

// Ferry defines a ferry ship.
type Ferry struct {
	Name         string
	Code         string
	Capacity     int
	loadedCars   []*Car
	acceptedCars []CarKind
	ticketPrices []TicketPrice
	totalRevenue float64

	enterChannel chan *Car
}

// NewFerry creates a new ferry.
func NewFerry(name, code string, capacity int, acceptedCars []CarKind, ticketPrices []TicketPrice) *Ferry {

	return &Ferry{
		Name:         name,
		Code:         code,
		Capacity:     capacity,
		acceptedCars: acceptedCars,
		ticketPrices: ticketPrices,
		enterChannel: make(chan *Car),
	}
}

// NewSmallFerry creates a small ferry.
func NewSmallFerry() *Ferry {
	ticketPrices := []TicketPrice{
		{CarKind: KindCar, Price: 30000},
		{CarKind: KindPickup, Price: 40000},
	}
	acceptedCars := []CarKind{KindCar, KindPickup}

	return NewFerry("Ferry Kecil", "S", 8, acceptedCars, ticketPrices)
}

// NewLargeFerry creates a large ferry.
func NewLargeFerry() *Ferry {
	ticketPrices := []TicketPrice{
		{CarKind: KindTruck, Price: 60000},
		{CarKind: KindBus, Price: 50000},
	}
	acceptedCars := []CarKind{KindTruck, KindBus}

	return NewFerry("Ferry Besar", "L", 6, acceptedCars, ticketPrices)
}

// CanAccept checks if ferry can load current kind of car.
func (ferry *Ferry) CanAccept(car *Car) bool {
	for _, accepted := range ferry.acceptedCars {
		// check if current type is accepted and has free space for this car.
		if car.Kind == accepted && len(ferry.loadedCars) < ferry.Capacity {
			return true
		}
	}

	return false
}

// IssueTicketFor .
func (ferry *Ferry) IssueTicketFor(car *Car) {
	for _, ticketPrice := range ferry.ticketPrices {

		// check ticket's price for current kind of car.
		if ticketPrice.CarKind == car.Kind {
			car.SetTicket(Ticket{
				Price:    ticketPrice.Price,
				IssuedBy: ferry,
			})

			break
		}
	}
}

// Open starts ferry listening for car queue.
func (ferry *Ferry) Open() {
	for {
		car := <-ferry.enterChannel

		ferry.loadedCars = append(ferry.loadedCars, car)
		ferry.totalRevenue = ferry.totalRevenue + car.GetTicket().Price

		fmt.Print("\nRute yang telah dilewati:\n")
		car.PrintRoutes()
	}
}

// GetTotalRevenue returns total revenue.
func (ferry *Ferry) GetTotalRevenue() float64 {
	return ferry.totalRevenue
}

// GetEnterChannel returns enterance channel.
func (ferry *Ferry) GetEnterChannel() chan<- *Car {
	return ferry.enterChannel
}

func (ferry *Ferry) String() string {
	return fmt.Sprintf("%s - %s", ferry.Code, ferry.Name)
}

// GenerateReport prints issued ticket.
func (ferry *Ferry) GenerateReport() {
	type detail struct {
		Qty          int
		TotalRevenue float64
	}

	issuedTickets := make(map[CarKind]*detail, 0)

	// fill default kind of car and value.
	for _, kind := range ferry.acceptedCars {
		issuedTickets[kind] = &detail{
			Qty:          0,
			TotalRevenue: 0.0,
		}
	}

	// calc tiket.
	for _, car := range ferry.loadedCars {
		kind := issuedTickets[car.Kind]
		kind.Qty = kind.Qty + 1
		kind.TotalRevenue = kind.TotalRevenue + car.GetTicket().Price
	}

	fmt.Printf("\nferry=[%s]\n", ferry)
	// prints.
	for kind, detail := range issuedTickets {
		fmt.Printf(
			"\nkind=%s \nqty=%d \ntotal_pendapatan=%.2f\n",
			kind, detail.Qty, detail.TotalRevenue,
		)
	}
	fmt.Print("\n--------------------------------------------------------\n")
}
