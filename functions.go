package main

import "fmt"

func add(a, b int) int {

	return a + b
}

func subAdd(a, b, c int) int{ //Yes go can take many params and assign them all at ones .. hwo cool is that.

	return (a - b) + c
}
//Another cool thing is the multiple returned values ....right? ..how cool is that. 

func vals () ( int , int){ 

	return  3, 4
}

//Varaidic functions => ANotehr cool go feature .. functions that can take many arguments..comes soooo handy at times.
//Here is one >>> 

func sum (nums...int){ 

	fmt.Println("nums => " , nums)

	total := 0 
	for _,num := range nums { //Yes you can iterate over the nums... why coz its basically like an []int ...
		total += num 
	}
	fmt.Println("Look this is the total of nums => ", total)

}
//CLOSURES ARE EVEN WAAAAY COOLER BUDDY!!!..OKAY ? 


func intSeq() func() int {  // This function returns a function that returns an int!!...see how cool that is ?? 
	i := 0 
	return func() int {
		i++
		return i 
	}
}


