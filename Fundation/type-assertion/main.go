package main

import "fmt"

func main() {
	var minhaVar interface{} = "João"
	println(minhaVar.(string))

	// res é o valor e ok é se a conversão deu certo, sempre um boolean
	res, ok := minhaVar.(int)
	fmt.Println(res, ok)
}
