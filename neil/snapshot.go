package main

import "fmt"

func main() {
	volumes := make(map[string]string)

	volumes["master"] = "vol-1593"

	fmt.Println(volumes)

}
