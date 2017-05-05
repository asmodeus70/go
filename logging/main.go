package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello " + user.Name)
	fmt.Println("====")
	fmt.Println("ID: " + user.Uid)
	fmt.Println("Username: " + user.Username)
	fmt.Println("Home Dir: " + user.HomeDir)

}
