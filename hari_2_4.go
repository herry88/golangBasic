package main

import "fmt"

func divide(x, y float64) (r float64, err string, d string) {
	if y == 0 {
		return 0, "zero division", "Something"
	}
	return x / y, "", "Something else"
}
func main() {
	r, _, d := divide(5, 0.)
	fmt.Println(r, d)
}
