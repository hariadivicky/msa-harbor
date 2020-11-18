package main

// Ticket defines ferry ticket.
type Ticket struct {
	IssuedBy *Ferry
	Price    float64
}

// TicketPrice defines ticket price info.
type TicketPrice struct {
	CarKind CarKind
	Price   float64
}
