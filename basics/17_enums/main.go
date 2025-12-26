package main

import "fmt"

/*
ENUMS IN GO
-----------
Go does NOT have enums like Java / C++.
Instead, we create:
1. A custom type
2. A set of constants of that type
This gives us type safety + readable code.
*/

// Custom type (not just an alias)
type OrderStatus int

// Enumerated constants using iota
const (
	Received OrderStatus = iota // 0
	Confirmed                   // 1
	Prepared                    // 2
	Delivered                   // 3
)

/*
String-based enum (useful for DB / APIs / logs)
*/
type OrderStatusString string

const (
	ReceivedS  OrderStatusString = "REC"
	ConfirmedS OrderStatusString = "CNF"
	PreparedS  OrderStatusString = "PREP"
	DeliveredS OrderStatusString = "DEL"
)

/*
Function that only accepts OrderStatus enum
This enforces correctness at compile time
*/
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)

	// This will NOT compile (type safety)
	// changeOrderStatus(2)
}
