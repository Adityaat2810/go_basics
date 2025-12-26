package main

import "fmt"

/*
INTERFACE CONCEPT
-----------------
An interface defines a CONTRACT.
Any type that implements ALL methods of the interface
automatically satisfies it (no explicit keyword like `implements`).
*/

/*
Concrete implementation: Stripe payment gateway
*/
type stripe struct{}

// stripe implements the `paymenter` interface
func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe:", amount)
}

/*
Interface definition
Naming convention: end with -er
*/
type paymenter interface {
	pay(amount float32)
}

/*
payment struct depends on an INTERFACE, not a concrete type
This enables dependency injection and loose coupling
*/
type payment struct {
	gateway paymenter
}

// Method that uses the interface
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func main() {
	// concrete implementation
	stripePaymentGateway := stripe{}

	// inject dependency
	newPayment := payment{
		gateway: stripePaymentGateway,
	}

	// make payment
	newPayment.makePayment(199)
}
