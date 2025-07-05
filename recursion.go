package main

import (
	"fmt"
)


func fib(n int) int{ 

	if n ==0 { 
	return 1
	}
	return n * fib(n-1)
}



func main(){ 

	fmt.Println(fib(5))

	var fib func(n int)int 

	fib = func(n int) int {
		
		if n < 2 { 
			return n 
		}
		return fib(n-1) + fib(n-2)	
	}	
	fmt.Println(fib(5))
}
