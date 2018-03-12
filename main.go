package main

import (
	"os";
	"fmt"
)

func main() {
	argsWithProg := os.Args
	
	fmt.Println(argsWithProg[3])
}