package main

import "fmt"

/*
* In Go, methods are functions attached to a type.

===========
   func (receiver Type) methodName() {}
==========

* There are two kinds of receivers:
 1. Value Receiver  // func (u User) method() {}
    Works on a copy
    Changes do NOT affect original data

 2. Pointer Receiver  //func (u *User) method() {}
    Works on the original memory
    Changes DO affect original data

* Pointer receiver lets a method modify the original struct.

******************
Methods are functions attached to a type, and they receive that type
(either by value or by pointer) so they can do things with it — often
to modify the struct.
******************

“Methods operate on structs.
Receivers decide whether the method sees a copy or the real thing.”

*/

type Account struct{
  Name string
  Balance int
}

// if use value reciever here it modify copy of it and orignally remained unchanged
func (acc *Account) Deposite(sum int){
  acc.Balance = acc.Balance + sum
}

// use value reciever here { we are doing read operation only }
func (acc Account) GetBalance() int {
	return acc.Balance
}

// [IMP!!] If method modifies the struct → use pointer receiver

func main(){
  acc := Account{
	Name: "aditya",
	Balance: 10000,
  }

  acc.Deposite(100)

  fmt.Println(acc.GetBalance())
}
