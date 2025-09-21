package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a > b { // > < >= <= == != && ||
		fmt.Println("a > b")
	} else { // Não tem else if
		fmt.Println("b > a")
	}

	switch b {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("a < b")
	}
}
