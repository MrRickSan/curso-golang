package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Numero de processadores: ", runtime.NumCPU())
}
