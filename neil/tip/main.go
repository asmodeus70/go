package main

import (
	"fmt"
	"github.com/fatih/color"
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
	fmt.Println("\nThe tip is")
	color.Set(color.FgGreen)
	var totaltip float32 = (bill * tip / 100)
	var billtotal float32 = (totaltip + bill)
	fmt.Printf("£%.2f \n", +totaltip)
	color.Set(color.FgWhite)
	fmt.Println("\nThe total bill is")
	color.Set(color.FgRed)
	fmt.Printf("£%.2f \n", +billtotal)
	color.Set(color.FgWhite)
	fmt.Println("\n")
}
