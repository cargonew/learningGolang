package main 

import (
	"fmt"
	"slices"
)


func slicing (){ 

	var s []string

	fmt.Println("uninit: ", s, s ==nil , len(s)==0)

	s = make([]string, 3)

	fmt.Println("empl:", s , " len:", len(s), " cap:", cap(s))

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"

	fmt.Println("set:", s)
    fmt.Println("get:", s[2])

	fmt.Println("len: ", len(s))
	
	s = append(s, "d")
	s = append(s, "e", "f", "g")

	fmt.Println("new len: ", len(s))
	fmt.Println("new cap: ", cap(s))

	fmt.Println(cap(s)-len(s))

	c := make([]string , len(s))
	copy(c,s)
	fmt.Println("This is how c looks like now: ", c)

	l := s[2:5] //This one takes from 2 to before 5 ... yes including 2 but excluding 5.

	fmt.Println("This is the 1st 'slice': ", l )

	l = s[:5] //takes all to before 5 ... yes excluding 5

	fmt.Println("This is the teh 2nd slice: ", l)

	l = s[2:]

	fmt.Println("This is teh last slice: ", l) // thisfrom takes all after and including 2

	t :=[]string{"1 ", "2 ", "3 "}

	fmt.Println("Look i initiated and declared  a slice and called it =>", t)

	t2 := make([]string , len(t))

	copy(t2, t)

	if slices.Equal(t, t2) { 

		fmt.Println("Hey look these guys are equal: ")
	}


	//We now look at the almighty 2D slices

	twoD := make([][] int , 3)
	
	for i := range 3 { 
		innerlen := i+1
		twoD[i] = make([]int,innerlen )
		
		for j := range innerlen{

			twoD[i][j] = i + j
		}
	}

	fmt.Println("Hey look i made a 2D , how cool is that => " , twoD)

}
