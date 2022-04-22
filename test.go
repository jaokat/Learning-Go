package main

import "fmt"

func main() {
	printer("Hello")

	// this demonstrates pass by copy
	mainNumber := 0 // mainNumber has primitive type INT
	mainNumber = 2
	// fmt.Printf("I am the original num:: %v \n", mainNumber)
	// counterReference(&mainNumber)
	// fmt.Printf("I am the original num after we called counterReference:: %v \n", mainNumber)

	fmt.Printf("I am the original num with a new value:: %v \n", mainNumber)
	// KAT WILL CALL COUNTERPASS
	counterPass(mainNumber)
	fmt.Printf("I am the original num after we called counterPass:: %v \n", mainNumber)

	mainNumber = counterPassWithReturn(mainNumber)
	fmt.Printf("I am the original num after we called counterPassWithReturn:: %v \n", mainNumber)

	mainNumber = multiplier(mainNumber)
	fmt.Printf("I am the original num after we called counterPassWithReturn:: %v \n", mainNumber)
}

func printer(wordToPrint string) {
	fmt.Println(wordToPrint)
}

func counterReference(num *int) {
    fmt.Println(num)
	fmt.Printf("I am no longer a copy of mainNum, I am refrence to mainNum:: %v \n", *num)
	*num++ // the scope of the variable is local to this function
	fmt.Printf("I am a reference of mainNum that was increased:: %v \n", *num)
} 

func counterPass(num int) {
	fmt.Printf("I am a copy of mainNum:: %v \n", num)
	num++ // the scope of the variable is local to this function
	fmt.Printf("I am a copy of mainNum that was increased:: %v \n", num)
} 

func counterPassWithReturn(num int) int {
	fmt.Printf("I am a copy of mainNum:: %v \n", num)
	num++ // the scope of the variable is local to this function
	fmt.Printf("I am a copy of mainNum that was increased:: %v \n", num)
	
	return num;
	
} 

func multiplier(num int) int {
	num = 4*num

	return num
}


/*
[0][1][2][3]
x = 0
y = *x 

*/
