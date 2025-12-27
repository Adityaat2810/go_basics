package main

import "fmt"

// parking lot

type ParkingLot struct{
  Id int
  Slots []ParkingSlot
}

type ParkingLoter interface{
	park(vehicle Vehicle)
	unPark(vehicle Vehicle)
}

func helpParking(lot *ParkingLot, minSize slotSize, v Vehicle){
  idxBest := -1
  diff := slotSize(1<<10)
  for idx, slot := range lot.Slots{
	currDiff := slot.Size-minSize
	if !slot.IsOccupied && currDiff >= 0 {
        if currDiff < diff {
			diff = currDiff
            idxBest = idx
		}

	}
  }

  if idxBest == -1{
	fmt.Println("Cannot park lot is full")
  }else{
	lot.Slots[idxBest].IsOccupied = true
	lot.Slots[idxBest].Vehicle = v
	fmt.Println("vehicle ", GetVehicleType(v.vehicleType), " is Parked to ", GetSlotType(lot.Slots[idxBest].Size)," slot")
  }
}

func (lot *ParkingLot) park(v Vehicle){
	switch v.vehicleType{
	  case Bike:{
		// can be parked to small, medium and large( priority to large)
		helpParking(lot, Small, v)
	  }
	  case Car:{
		// can be parked to medium & large( priority to medium )
		helpParking(lot, Medium, v)

	  }

	  case Truck:{
		// can be parked to large
		helpParking(lot, Large, v)
	  }

	}

}

func (lot *ParkingLot) unPark(v Vehicle) {
	for i := range lot.Slots {
		slot := &lot.Slots[i]
		if slot.IsOccupied && slot.Vehicle.Id == v.Id {
			slot.IsOccupied = false
			slot.Vehicle = Vehicle{}
			return
		}
	}
}


// parking slot
type ParkingSlot struct{
  Id int
  Size slotSize
  IsOccupied bool
  Vehicle Vehicle

}

type slotSize int

const (
	Small slotSize = iota
	Medium
	Large
)

func GetSlotType(s slotSize) string{
	switch s{
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
// vehicle
type Vehicle struct{
	Id int
	vehicleType vehicleType
}

type vehicleType int

const (
  Bike vehicleType = iota
  Car
  Truck
)

func GetVehicleType(v vehicleType) string{
	switch v{
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


func main(){
	lot := ParkingLot{
		Id: 1,
		Slots: []ParkingSlot{
			{Id: 1, Size: Small},
			{Id: 2, Size: Small},
			{Id: 3, Size: Small},
			{Id: 4, Size: Medium},
			{Id: 5, Size: Medium},
			{Id: 6, Size: Medium},
			{Id: 7, Size: Large},
			{Id: 8, Size: Large},
			{Id: 9, Size: Large},
			{Id: 10, Size: Small},
		},
	}

	car1 := Vehicle{Id: 1, vehicleType: Car}
	bike1 := Vehicle{Id: 2, vehicleType: Bike}
	truck1 := Vehicle{Id: 3, vehicleType: Truck}

	car2 := Vehicle{Id: 4, vehicleType: Car}
	bike2 := Vehicle{Id: 5, vehicleType: Bike}
	truck2 := Vehicle{Id: 6, vehicleType: Truck}



	car3 := Vehicle{Id: 7, vehicleType: Car}
	bike3 := Vehicle{Id: 8, vehicleType: Bike}
	truck3 := Vehicle{Id: 9, vehicleType: Truck}




	lot.park(car1)
	lot.park(bike1)
	lot.park(truck1)

	lot.park(car2)
	lot.park(bike2)
	lot.park(truck2)

	lot.park(car3)
	lot.park(bike3)
	lot.park(truck3)

	fmt.Println(lot)
}