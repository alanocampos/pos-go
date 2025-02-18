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

type CotacaoDolar struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type CotacaoDolarModel struct {
	gorm.Model
	Bid string
}

func main() {

	db, err := gorm.Open(sqlite.Open("cotacao.db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&CotacaoDolarModel{})
	if err != nil {
		panic("failed to migrate database")
	}

	http.HandleFunc("/cotacao", handleCotacao(db))
	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}

func handleCotacao(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cotacao, err := getCotacao(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao obter cotação"))
			log.Println("Erro ao obter cotação:", err)
			return
		}

		err = saveCotacao(ctx, db, cotacao.USDBRL.Bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao salvar cotação"))
			log.Println("Erro ao salvar cotação:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	}
}

func getCotacao(ctx context.Context) (*CotacaoDolar, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cotacao CotacaoDolar
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout ao obter cotação da API")
		return nil, ctx.Err()
	default:
		return &cotacao, nil
	}
}

func saveCotacao(ctx context.Context, db *gorm.DB, bid string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	cotacaoModel := CotacaoDolarModel{Bid: bid}

	result := db.WithContext(ctx).Create(&cotacaoModel)
	if result.Error != nil {
		select {
		case <-ctx.Done():
			log.Println("Timeout ao salvar cotação no banco de dados")
			return ctx.Err()
		default:
			log.Println("Erro ao salvar cotação no banco de dados:", result.Error) // Log the specific GORM error
			return result.Error
		}
	}

	return nil
}
