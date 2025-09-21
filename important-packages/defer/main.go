package main

import "fmt"

func main() {

	//O defer deixa o comando que ele está precedendo por último na execução, muito usado para fechar conexões http

	fmt.Println("Hello world")
	defer fmt.Println("Hello world 2")
	fmt.Println("Hello world 3")
}
