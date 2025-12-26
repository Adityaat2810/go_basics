package main

import "fmt"

/***
-- An interface is a contract.
***/


// razor-pay payemnt gateway
type razorpay struct {}


func (r razorpay) pay( amount float32){
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// stripe-payment gateway
type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe", amount)
}

// for testing
type fakePayment struct{}

func (f fakePayment) pay(amount float32){
  fmt.Println("making payement using fake gateway for testing", amount)
}

// payment

/*******************
type Payment struct {}

func (p Payment) makePayment(amount float32){
//   razorpayPaymentGateway := razorpay {}
//   razorpayPaymentGateway.pay(amount)
											 // we need to modify actaul code(method)
											 // break open-close princile ( one of SOLID principle)
											 // methods should be open for extension but closed for modification
  stripePaymentGateway := stripe{}
  stripePaymentGateway.pay(amount)
}

*******************/

// bit improved
type Payment struct{
  gateway stripe                   // problem is still not solved
                                   // changeing gatway still require modification of struct
}

func (p Payment) makePayment(amount float32){
	p.gateway.pay(amount) //
}



func main(){
  stripePaymentGateway := stripe{}
  // fakePaymentGateway := fakePayment{}
  newPayment := Payment{
    gateway: stripePaymentGateway,
    // gateway: fakePaymentGateway // cannot pass fake payment gateway here as in struct payment - payemnt method is literally hardcoded 
  }
  newPayment.makePayment(100)
}
