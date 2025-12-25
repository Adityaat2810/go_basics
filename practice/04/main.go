package main

import "fmt"

/***
QUES 6
  1. Update name of a user using ID
  2. Explain why pointer is needed

  { we need pointer so that chaneg reflect in map not in some copy of it (pass by refrence) }
***/
type User struct {
	Name string
	Age int
	Email string
}

func updateUserAge(users *map[int]User, newAge int, id int){
   u, ok := (*users)[id]
   if !ok{
    fmt.Println("no user with id ", id)
   }

   u.Age = newAge
   (*users)[id] = u
}

func createUser(users *map[int]User, name string , age int , email string){
  (*users)[len(*users)+1] = User{
	Name: name,
	Age: age,
	Email: email,
  }
}

func updateNameById(users * map[int] User, id int, newName string){
  u, ok := (*users)[id]
  if !ok {
    fmt.Println("no user with id", id)
    return
  }
  u.Name = newName
  (*users)[id] = u
  fmt.Println("name updated to ", (*users)[id].Name)
}

// Q7 { Return a slice of adult users. }
func findAdultUsers(users []User) []User{
	var adultUsers []User
	for _, user := range users{
		if user.Age >= 18 {
			adultUsers = append(adultUsers, user)
		}
	}

	return  adultUsers
}

func main(){
  var users = make(map[int]User)

  createUser(&users, "aditya", 22, "adi@example.com")
  createUser(&users, "ravi", 35, "ravi@example.com")
  createUser(&users, "cp", 33, "cp@example.com")
  createUser(&users, "devansh", 21, "dev@example.com")

  updateNameById(&users, 2, "ravindra")

  for id, user := range users{
	fmt.Println("ID:", id, " Name:", user.Name)
  }

}
