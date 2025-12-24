package main

import "fmt"

func addNum(a int, b int) int {
	return a+b
}

// also
func subNum(a, b int) int {
	return  a-b;
}

// returning multiple values
func getLanguages()(string, string, string, bool){
	return "goLang", "javascript", "c", true
}

/***************/
// In go functions are first calss citizen
// so can be assigned to varible, also can be
// passed as arguments

func processIt (fn func(a int) int){
  pro := fn(1)
  fmt.Println(pro)
}

func processIt2 () func(a int) int{
  return  func(a int) int {
	return  2
  }
}

func main(){
	res := addNum(12, 34)
	lang1, lang2, _, boolu:= getLanguages() // if we don't wnat to use something use _ at its place it wont throw error

	fmt.Println(res, lang1, lang2, boolu)

	fn := func(a int) int {
		return 2
	}

	processIt(fn)

	nf := processIt2()
	nf(2);
	
}
