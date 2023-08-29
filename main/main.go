package main

import (
	"fmt"
	"os"

	handlers "github.com/tliddle1/calc-apps/handlers"
	"github.com/tliddle1/calcy-lib/calcy"
)

func main() {
	// Ensure that there are exactly two command line arguments
	//if len(os.Args) != 3 {
	//	fmt.Println("Usage: go run main.go <num1> <num2>")
	//	panic("Too many arguments")
	//	return
	//}
	myHandler := handlers.New(os.Stdout, calcy.Addition{})
	err := myHandler.Handle([]string{"1", "2"})
	if err != nil {
		fmt.Println(err)
		return
	}
}
