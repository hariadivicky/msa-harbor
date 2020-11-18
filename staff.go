package main

import "fmt"

// Staff defines a parking staff.
type Staff struct {
	Name         string
	Interest     float64
	TotalRevenue float64
	parkedCars   []*Car
}

// NewStaff creates a new parking staff.
func NewStaff(name string, interest float64) *Staff {
	return &Staff{
		Name:         name,
		Interest:     interest,
		TotalRevenue: 0,
		parkedCars:   make([]*Car, 0),
	}
}

// Park simulates a staff parking current car.
func (staff *Staff) Park(car *Car) {
	collected := car.GetTicket().Price * staff.Interest / 100
	staff.TotalRevenue = staff.TotalRevenue + collected

	fmt.Printf(
		"\n[%s]: menerima jasa parkir \ntiket mobil diparkir=%.2f \nditerima=%.2f (%.2f%%) \ntotal penghasilan=%.2f\n",
		staff.Name, car.GetTicket().Price, collected, staff.Interest, staff.TotalRevenue,
	)

	staff.parkedCars = append(staff.parkedCars, car)
}
