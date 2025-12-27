package main

import (
	"errors"
	"fmt"
)

//
// ===== ENUMS =====
//

type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	switch v {
	case Bike:
		return "Bike"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	default:
		return "Unknown"
	}
}

type SlotSize int

const (
	Small SlotSize = iota
	Medium
	Large
)

func (s SlotSize) String() string {
	switch s {
	case Small:
		return "Small"
	case Medium:
		return "Medium"
	case Large:
		return "Large"
	default:
		return "Unknown"
	}
}

//
// ===== VEHICLE INTERFACE =====
//

type Vehicle interface {
	ID() int
	Type() VehicleType
}

type BaseVehicle struct {
	id          int
	vehicleType VehicleType
}

func (v BaseVehicle) ID() int {
	return v.id
}

func (v BaseVehicle) Type() VehicleType {
	return v.vehicleType
}

//
// ===== PARKING SLOT =====
//

type ParkingSlot struct {
	id       int
	size     SlotSize
	occupied bool
	vehicle  Vehicle
}

func (s *ParkingSlot) CanFit(v Vehicle) bool {
	switch s.size {
	case Small:
		return v.Type() == Bike
	case Medium:
		return v.Type() == Bike || v.Type() == Car
	case Large:
		return true
	default:
		return false
	}
}

//
// ===== PARKING LOT =====
//

type ParkingLot struct {
	id    int
	slots []ParkingSlot
}

func NewParkingLot(id int, slots []ParkingSlot) *ParkingLot {
	return &ParkingLot{
		id:    id,
		slots: slots,
	}
}

func (p *ParkingLot) Park(v Vehicle) error {
	for i := range p.slots {
		slot := &p.slots[i]

		if !slot.occupied && slot.CanFit(v) {
			slot.occupied = true
			slot.vehicle = v
			fmt.Printf("✅ %s parked in slot %d (%s)\n",
				v.Type(), slot.id, slot.size)
			return nil
		}
	}

	return errors.New("❌ no available slot")
}

func (p *ParkingLot) Unpark(vehicleID int) error {
	for i := range p.slots {
		slot := &p.slots[i]

		if slot.occupied && slot.vehicle.ID() == vehicleID {
			fmt.Printf("🚗 Vehicle %d left slot %d\n",
				vehicleID, slot.id)
			slot.occupied = false
			slot.vehicle = nil
			return nil
		}
	}

	return errors.New("❌ vehicle not found")
}

func (p *ParkingLot) AvailableSlots() {
	count := map[SlotSize]int{
		Small:  0,
		Medium: 0,
		Large:  0,
	}

	for _, slot := range p.slots {
		if !slot.occupied {
			count[slot.size]++
		}
	}

	fmt.Println("\n📊 Available Slots:")
	for size, c := range count {
		fmt.Printf("- %s: %d\n", size, c)
	}
}

//
// ===== MAIN =====
//

func main() {
	parkingLot := NewParkingLot(1, []ParkingSlot{
		{id: 1, size: Small},
		{id: 2, size: Small},
		{id: 3, size: Medium},
		{id: 4, size: Medium},
		{id: 5, size: Large},
		{id: 6, size: Large},
	})

	car1 := BaseVehicle{id: 1, vehicleType: Car}
	car2 := BaseVehicle{id: 2, vehicleType: Car}
	bike1 := BaseVehicle{id: 3, vehicleType: Bike}
	truck1 := BaseVehicle{id: 4, vehicleType: Truck}

	parkingLot.Park(bike1)
	parkingLot.Park(car1)
	parkingLot.Park(car2)
	parkingLot.Park(truck1)

	parkingLot.AvailableSlots()

	parkingLot.Unpark(2)
	parkingLot.AvailableSlots()
}
