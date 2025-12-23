package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dictionary
func main(){
	/******* creating maps *****************/
	// make (map[typeof(key)] typeof(value))
	m := make( map[string]string )

	// setting an element
	m["name"] = "aditya"
	m["area"] = "backend"
	m["phone"] = "8181717190"

	// get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["age"])
	// for non-existing key it will print zero value ( string -> "", int -> 0, boolean -> false)

	// len of map
	fmt.Println("length of map:", len(m))

	// delete an element
	delete(m, "area")
	fmt.Println(m)

	// clering a map
	clear(m)
	fmt.Println("after clearing:", m)

	/*************declaring map*****************/
	m2 := map[string] int{ "price": 40, "phones": 3} // if elements available from starting

	// check if elemnt in map
	_, ok := m2["price"]
	if(ok) {
		fmt.Println("ALL ok")
	}else {
		fmt.Println("Not ok")
	}

	k, ok2 := m2["phones"]
	if(ok2){
		fmt.Println(k) // k hold the value for key{ if key not present it will give 0 }
	}
	fmt.Println(m2)
	m3 := map[string] int{ "price": 40, "phones": 3}
	// to check if 2 maps are equal
	fmt.Println(maps.Equal(m2, m3))
}
