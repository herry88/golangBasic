package main

import "fmt"

func main() {

	var s = []string{"abc", "cde", "abc", "cde", "cde"}
	var m = make(map[string]int)

	// dataRecive := []Data{"1"}

	for _, st := range s {
		m[st]++
	}
	fmt.Println(m)
}
