package main

import (
	"fmt"
	"os"
)

func main() {
	mode := os.Args[1]
	file := os.Args[2]
	s, _ := os.Executable()
	fmt.Println(mode, file, s)
}
