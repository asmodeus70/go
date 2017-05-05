package main

import "fmt"

// Declare constant
const Title = "My Personal Details"

// Declare package variable
var Country = "UK"

func main() {
	fname, lname := "Neil", "Radley"
	age := 48
	// Print constant variable
	fmt.Println(Title)
	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)
	// Print package variable
	fmt.Println("Country:", Country)
}
