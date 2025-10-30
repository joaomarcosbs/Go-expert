package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerResponse struct {
	Bid string `json:"bid"`
}

func saveToFile(filename string, bidValue string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", bidValue))
	if err != nil {
		return err
	}

	return nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("[CLIENT] Erro ao criar requisição: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("[CLIENT] Erro ao fazer requisição ao servidor: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("[CLIENT] Erro ao ler resposta do servidor: %v", err)
	}

	var serverResponse ServerResponse
	err = json.Unmarshal(body, &serverResponse)
	if err != nil {
		log.Fatalf("[CLIENT] Erro ao decodificar JSON da resposta: %v", err)
	}

	err = saveToFile("../cotacao.txt", serverResponse.Bid)
	if err != nil {
		log.Fatalf("[CLIENT] Erro ao salvar cotação no arquivo: %v", err)
	}

	log.Println("[CLIENT] Cotação salva com sucesso em cotacao.txt!")
}
