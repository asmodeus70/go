package main

import "fmt"

func main() {
	volumes := make(map[string]string)

	volumes["master"] = "vol-1593"
	volumes["slave01"] = "vol-2051"
	volumes["slave02"] = "vol-3594"

	fmt.Println(volumes)

}
