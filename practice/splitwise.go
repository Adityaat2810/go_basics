package main

type User struct{
  Id int
  Name int
}

// balances[debtor][creditor] = amount
// → debtor owes creditor
type BalanceSheet struct{
  balances map[int]map[int]int
  // debtor, creditor, amount
}

type Group struct{
  users []User
  balanceSheet *BalanceSheet
}