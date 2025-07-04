package main

import (
	"fmt"
	"maps"
)



func mapping(){ 

	m := make(map[string]int)

	m["k1"] = 1 
	m["k2"] = 2
	m["k3"] = 3 


	fmt.Println("map: ", m)

	v1 := m["k1"]
	
	v3 := m["k3"]

	fmt.Println("This is value one: ", v1 , " and this is value three: ", v3)

	fmt.Println(len(m))

	delete(m , "k2")

	fmt.Println(len(m))

	clear(m)

	fmt.Println(len(m))


	n := map[string]int{"foo":1, "bar": 2}

	n2 :=  map[string]int{"foo":1, "bar": 2}


	if maps.Equal(n ,n2){ 

		fmt.Println("These bad boys are Equal indeed.")
	}


}

