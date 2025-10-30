package main

import (
	"context"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"time"
)

type Bid struct {
	Value string
	gorm.Model
}

type Conversion struct {
	USDBRL Dolar
}

type Dolar struct {
	Bid string `json:"bid"`
}

func CotacaoHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[SERVER] Erro ao buscar cotação: %v", err)
			http.Error(w, "[SERVER] Erro ao buscar cotação: timeout da API excedido", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var dolar Conversion
		err = json.Unmarshal(body, &dolar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("[SERVER] Cotação recebida com sucesso: %s", dolar.USDBRL.Bid)

		bid := Bid{Value: dolar.USDBRL.Bid}

		err = saveBid(db, &bid)
		if err != nil {
			http.Error(w, "[SERVER] Erro ao salvar no banco: timeout do banco de dados excedido", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dolar.USDBRL)
	}
}

func saveBid(db *gorm.DB, bid *Bid) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := db.WithContext(ctx).Create(bid).Error
	if err != nil {
		log.Printf("[SERVER] Erro ao salvar cotação no banco: %v", err)
		return err
	}

	log.Println("[SERVER] Cotação salva com sucesso no banco de dados.")
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("../cotacao.db"), &gorm.Config{})
	if err != nil {
		panic("[SERVER] Falha ao conectar com o banco de dados")
	}

	db.AutoMigrate(&Bid{})

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler(db))

	log.Println("[SERVER] Servidor iniciado na porta 8080")
	muxErr := http.ListenAndServe(":8080", mux)
	if muxErr != nil {
		log.Fatal(muxErr)
	}
}
