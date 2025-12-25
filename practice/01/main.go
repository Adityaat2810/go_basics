package main

import "fmt"

/*****
Q. Create a Student struct with:
Name
RollNo
Marks
Print details if marks ≥ 40 else print “Fail”.
*****/

type Student struct {
	Name   string
	RollNo int
	Marks  int
}

func createStudent(name string, rollNo int, marks int) Student {
	return Student{
		Name:   name,
		RollNo: rollNo,
		Marks:  marks,
	}
}

func isPass(s Student) {
	if s.Marks >= 40 {
		fmt.Printf("Pass: %+v\n", s)
	} else {
		fmt.Printf("Fail: Name=%s Marks=%d\n", s.Name, s.Marks)
	}
}

func main() {
	students := []Student{
		createStudent("Aditya", 1, 55),
		createStudent("Ravi", 2, 33),
		createStudent("Ravi", 3, 23),
		createStudent("Ravi", 4, 88),
		createStudent("Ravi", 5, 12),
		createStudent("Ravi", 6, 96),
	}

	for _, s := range students {
		isPass(s)
	}
}
