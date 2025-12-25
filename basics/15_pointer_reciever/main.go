package main
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

func main(){

}
