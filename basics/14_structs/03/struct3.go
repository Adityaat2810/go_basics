package main
// Struct with Pointers (VERY IMPORTANT)

import "fmt"

type User struct {
	id    int
	name  string
	email string
	age   int
}

func main() {
	u1 := User{1, "aditya", "aditya@example.com", 25}
	// fmt.Println(u1)

	// Pointer to struct
	var u2 *User = &u1
	fmt.Println(u2)

	// Accessing fields using pointer
	fmt.Println(u2.name)

	// Updating fields using pointer
	u2.age = 30
	fmt.Println(u1.age) // Reflects in original struct

	// Creating struct using new keyword
	u3 := new(User)
	u3.id = 3
	u3.name = "new user"
	u3.email = "newuser@example.com"
	u3.age = 28
	fmt.Println(u3)

	// Struct as Function Parameter
	printUserInfo(u1)
	printUserInfo(*u2)
	printUserInfo(*u3)

}

func printUserInfo(u User) {
	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", u.id, u.name, u.email, u.age)
}


/****
[IMPORTANT]

func update(u User) {
    u.Name = "X"   // modifies only the copy
}


func update(u *User) {
    u.Name = "X"   // modifies the original struct via pointer
}


*****/