package main

import "fmt"

func sum( nums ...int) int {
  total := 0
  for _, num := range nums {
	total += num
  }

  return  total
}

func reciveAnything(_ ...interface{}) interface{}{
	return "jai ho"
}

func main(){
	fmt.Println(1, 2, 3, 4, 5, "hello", false) // can pass any (n) no of arguments
	res := sum(1, 2, 3,4 ,6 ,7 ,8 ) // so it is variadic function
	fmt.Println(res)

	res2 := reciveAnything("ahd", true, false, 1)

	nums := [] int {1, 2, 33, 44, 12, 34}
	fmt.Println(res2, sum(nums...))

}
