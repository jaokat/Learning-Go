package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Hello, 世界")

	//Slices
	// I've created an array that has 5 spaces to be filled
	
	/*
	0 1 2 3 4
	{[obj][][][obj][]}

	obj
	obj

	*/

	// hey get me the object in position 0
	// hey get me the object in position 3

	var veggies = []string{
		"brocolli", 
		"mushroom", 
		"potato",
	}

	spew.Dump(veggies)

	fmt.Println(veggies[1])

	/*

	0 1 2
	{[brocolli][mushroom][potato]}

	*/


	spices := []string{
		"pepper",
		"salt",
		"turmeric",
		"pepper",
	}

	spew.Dump(spices)

	/*
	for SOME STARTING NUMBER, SOME ENDING NUMBER, increment 

		do something
	}
	*/

	fmt.Println("---------------------------------------")

	for i := 0; i < 3; i++ {
		fmt.Println(veggies[i])
	}


	//range has two vars the index number and the value at that index
	for _, kat := range spices {

		fmt.Println(kat)
		if kat  == "pepper" {
			fmt.Println("I found it!")
		}

		if kat == "salt" {
			fmt.Println("I found it!")
		}

	}


}





