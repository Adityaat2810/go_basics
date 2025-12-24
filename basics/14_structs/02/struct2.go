package main

// public vs private field
/*
* Capital letter → PUBLIC (Exported)
* Small letter → PRIVATE (Unexported)
 */

// Public -> Accessible from other packages

/*********
package main

import "project/models"

func main() {
    u := models.User{Name: "Aditya"}
    fmt.Println(u.Name)
}


**********/
// Private -> Accessible ONLY inside the same package

/**********

package models

type User struct {
    name string
}

u := models.User{name: "Aditya"} // ❌ compile error
fmt.Println(u.name)              // ❌

***********/


/*
* Private does NOT mean “only inside struct”
* Private means “only inside package”
*/


func main() {
    // u := models.UserPublic{Name: "Aditya"}

}
