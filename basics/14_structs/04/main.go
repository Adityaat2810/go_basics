package main

import "fmt"

type User struct{
	id int
	name string
	email string
}

func main(){
  // Structs with array
  var users []User
  users = append(users, User{1, "aditya", "adityasinics10@gmail.com"})

  users2 := []User {
	{1, "aditya", "adi@gmail.com"},
	{2, "rahul", "rahul@gmail.com"},
	{3, "ravi", "ravi@gmail.com"},
  }
  fmt.Println(users)
  fmt.Println(users2)

  // Structs with map
  var users3  = make(map[int]User)
  users3[1] = User{1, "aditya", "adi@gmail.com"}

  userMap2 := map[int] User{
	1: {1, "aditya", "adi@gmail.com"},
	2: {2, "rahul", "rahul@gmail.com"},
  }


  for idx, user := range users3{
	fmt.Println(user, idx)
  }

  fmt.Println(userMap2)
}