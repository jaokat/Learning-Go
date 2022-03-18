package main

import "fmt"

func main() {
	printer("Hello")

	mainNumber := 0
	fmt.Println(mainNumber)
	counter(mainNumber)
	fmt.Println(mainNumber)
}

func printer(wordToPrint string) {
	fmt.Println(wordToPrint)
}

func counter(num int) {
	num++
	fmt.Println(num)
} 
