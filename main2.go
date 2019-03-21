package main

import "fmt"
import "os"

func main() {
	argsWithprog := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithprog)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
