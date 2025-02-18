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

type CotacaoDolarResponse struct {
	Usdbrl struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	var cotacao CotacaoDolarResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	err = saveCotacaoToFile(cotacao.Usdbrl.Bid)
	if err != nil {
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	fmt.Println("Cotação salva com sucesso em cotacao.txt")

	select {
	case <-ctx.Done():
		log.Println("Timeout ao receber resposta do servidor")
		return
	default:
	}

}

func saveCotacaoToFile(bid string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", bid))
	if err != nil {
		return err
	}

	return nil
}
