package main

import "fmt"

/******
Q9 Create a struct BankAccount:
ID
Balance

Functions:
deposit(amount int)
withdraw(amount int) → use conditionals

Use pointer receiver

*******/

/***

Q10.
a) Why slice doesn’t need pointer
👉 Slice already behaves like a reference

Real reason
- A slice is not the actual data.
- A slice is a descriptor that contains:
   (pointer to array, length, capacity)

b) Why map value needs pointer
👉 Map elements are NOT addressable
- Map internally can rehash / move memory
- Go does NOT allow direct modification of stored values


c) Why struct should use pointer receiver
- If method changes state → pointer required

***/
type Account struct{
  Name string
  Balance int
} 

func createAccount(users map[int] *Account, name string) *Account{
  key := len(users) + 1
  users[key] = &Account {
	Name: name,
	Balance: 0,
  }
  return  users[key]

}

func (acc *Account) deposite(amount int){
  acc.Balance += amount
}

func (acc *Account) withdrow(amount int){
  if acc.Balance < amount {
	fmt.Println("Insufficent balance !")
  }else{
    acc.Balance -= amount
  }
}

func (acc *Account) checkBalance(){
	fmt.Println("Balanc is ", acc.Balance)
}

func main(){
	users := make(map[int]* Account)
	u1 := createAccount(users, "aditya")
	createAccount(users, "ravi")
	createAccount(users, "cp")
	createAccount(users, "devansh")

	u1.deposite(500)
	u1.checkBalance()
	u1.withdrow(200)
	u1.checkBalance()
	u1.withdrow(400)
	u1.checkBalance()


}