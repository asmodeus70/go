package main

import (
	"fmt"
)

func main() {
	// Ask user for the bill amount
	fmt.Println("How much is the bill?")
	var bill float32
	fmt.Scanln(&bill)
	// Ask the user how much of a tip is required
	fmt.Println("How big is the tip gong to be?")
	var tip float32
	fmt.Scanln(&tip)

	// Now we calculate the tip
	fmt.Println("The tip is")
	var totaltip float32 = (bill * tip / 100)
	var billtotal float32 = (totaltip + bill)
	fmt.Printf("%.2f \n", +totaltip)
	fmt.Println("The total bill is")
	fmt.Printf("%.2f \n", +billtotal)
}
