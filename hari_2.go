package main

import (
	"fmt"
)

func main() {
	if a := 4; a > 7 {
		fmt.Println("if")
	} else if a > 5 {
		fmt.Println("else1")
	} else if a > 3 {
		fmt.Println("else2")
	} else {
		fmt.Println("else3")
	}

}
