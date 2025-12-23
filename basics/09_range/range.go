package main

import "fmt"

// iterating over data structure

func main(){
  /***************** iterating array or slice using raneg ********/
  nums := []int {6, 7, 8}

  // iterating using simple for loop
  for i :=0 ; i<len(nums); i++{
	//fmt.Println(nums[i])
  }

  // iterating using range
  for _, num := range nums{
	fmt.Println(num)
  }

  for idx, num := range nums{
    fmt.Println("on index ", idx, " number is ", num)
  }

  /************** iterating maps using range ***********************/
  m := map[string]string {"fname":"aditya", "lname":"saini", }

  for key, value := range m{
	fmt.Println("key is ", key, "value is ", value)
  }

  for k := range m {  // if only single is given than it will hold key
	fmt.Println("for key ", k, "value is ",m[k])
  }

  /**************** iterating strings *****************************/
  for sbr , character := range "aditya" {
	// [ IMP!! ]charcter will hold the unicode { code point rune }
	// sbr hold the starting point of rune
	fmt.Println("sbr is starting point of rune  ", sbr, "char is ", character)

	// to actually get the char
	fmt.Println("char here is ", string(character))
  }




}