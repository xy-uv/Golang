package main

import (
	"fmt"
	"maps"
)

// Map is the hash,object,dict,ect.
func main() {
	// creating map

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns zero value

	mp := make(map[string]int)
	mp["age"] = 30
	mp["price"] = 50
	fmt.Println(mp["phone"])
	fmt.Println(len(mp))

	// delete(mp, "price") // deleting single fields in a map
	// clear(mp) // clearing the full of map

	fmt.Println(mp)
	fmt.Println(mp)

	ms:= map[string]int{"model":97,"price":999,"lunch":2025,"market":01}
	v,ok := ms["price"]
	fmt.Println(v)
	if ok{
		fmt.Println("All is ok")
	}else{
		fmt.Println("Something went wrong!")
	}
	ms1 := map[string]int{"price": 40, "phones": 3}
	ms2 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(maps.Equal(ms1, ms2))

}