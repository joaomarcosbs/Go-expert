package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	tamanho, err := f.Write([]byte{'a', 'b', 'c', 'd', 'e', 'f'})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tamanho)

	f.Close()

	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// ler arquivos aos poucos, em caso de ter arquivos muito grandes para a memória

	arquivo2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Remover arquivo
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
