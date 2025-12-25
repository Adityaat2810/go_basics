package main

import "fmt"

/****
Create a slice of User structs.
Add 3 users
Print only users whose age > 18

****/

type User struct{
	Id int
	Age int
	Name string
}

/****
Here pointer not required , in go slices are passed throgh refrence
****/
func addUser(users *[]User, id int, age int , name string){
  *users = append(*users, User{
	Id: id,
	Age: age,
	Name: name,
  })
}

func main(){
	var users []User
	p := &users

	addUser(p, 1, 22, "aditya");
	addUser(p, 2, 13, "raju");
	addUser(p, 3, 14, "nobita");

	for _, user := range users {
		if user.Age > 18{
			fmt.Println(user)
		}
	}

}


/*****
[IMP]
--------- PASS BY VALUE -------------
- int
- float
- bool
- struct
- array

---------- PASS BY REFRENCE ----------
- slice
- map
- channel
- function
- interface 

******/
