package main

import "fmt"

func main() {
	switch 99 {
	case 1, 2, 3:
		fmt.Println("1 2 3")
		fmt.Println("xxx")
	case 4, 5, 7:
		fmt.Println("4 5 7")
		fmt.Println("xxx")
	default:
		fmt.Println("sonstwas")
	}
}
