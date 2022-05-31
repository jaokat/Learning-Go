package main

import "fmt"

func main() {
	printer("Hello")

	// this demonstrates pass by copy
	mainNumber := 0 // mainNumber has primitive type INT
	mainNumber = 2
	sizeChecker(mainNumber)
	// fmt.Printf("I am the original num:: %v \n", mainNumber)
	// counterReference(&mainNumber)
	// fmt.Printf("I am the original num after we called counterReference:: %v \n", mainNumber)

	fmt.Printf("I am the original num with a new value:: %v \n", mainNumber)
	// KAT WILL CALL COUNTERPASS
	counterPass(mainNumber)
	fmt.Printf("I am the original num after we called counterPass:: %v \n", mainNumber)

	mainNumber = counterPassWithReturn(mainNumber)
	fmt.Printf("I am the original num after we called counterPassWithReturn:: %v \n", mainNumber)
	sizeChecker(mainNumber)

	mainNumber = multiplier(mainNumber)
	fmt.Printf("I am the original num after we called multiplier:: %v \n", mainNumber)
	sizeChecker(mainNumber)

	// PUT YOUR NEW FUNCITON CALL HERE
	mainNumber = multiply(mainNumber, 8)
	fmt.Printf("I am the new multiply:: %v \n", mainNumber)
	sizeChecker(mainNumber)

	// try to write an if statement that prints "TADA" if the value of mainNumber is greater than 50
	if (mainNumber > 50) {
		fmt.Printf("TADA \n") 
	}

	// GO Over Slices, Structs, and For loops

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

	return num

}

func multiplier(num int) int {
	num = 4 * num

	return num
}

// Create another function that allows you to multiply any number with the given number
// and return the answer

func multiply(num, anyNumber int) int {
	num = anyNumber * num

	return num
}

// create a function that prints certain statements if the value of the number passed in is within the following range
// less than 10 print "tiny"
// less than 50 but greater than 10 print "medium"
// otherwise print "large"
// The name of the function can be "sizeChecker"

func sizeChecker(mainNumber int) {
	if mainNumber < 10 {
		fmt.Printf("Tiny \n")
	} else if mainNumber >= 10 && mainNumber < 50 {
		fmt.Printf("Medium \n")
	} else {
		fmt.Printf("Large \n")
	}
}

/*
[0][1][2][3]
x = 0
y = *x

*/
