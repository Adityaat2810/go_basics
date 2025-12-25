package main

import "fmt"

/******
Q3->
Write a function:
func updateEmail(u *User, newEmail string)
Update email using pointer.

Q4
Create a map:
map[int]User
Loop through it and print:
ID: 1 Name: X


********/


type Users struct{
  Id int
  Email string
  Name string
  Age int

}


func updateEmail(u *Users, newMail string) {
  u.Email = newMail
}


func main(){
  u := Users{
    Id: 1,
    Email: "adi@example.com",
    Name: "aditya",
    Age: 22,
  }

  p := &u

  fmt.Println(u.Email)
  updateEmail(p, "r@example.com")
  fmt.Println(u.Email)

  var usersMap = make(map[int] Users)
  usersMap[1] = Users{
    1, "aditya@gmail.com", "aditya", 12,
  }

  usersMap[2] = Users{
    2, "aditya@gmail.com", "aditya", 12,
  }

  usersMap[3] = Users{
    3, "aditya@gmail.com", "aditya", 12,
  }

  usersMap[4] = Users{
    4, "aditya@gmail.com", "aditya", 12,
  }

  for _, user := range usersMap{
    fmt.Println("Id is ",user.Id, "Name ", user.Name)
  }
}
