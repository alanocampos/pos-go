package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CotacaoDolar struct {
	Usdbrl struct {
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	//ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

	var c CotacaoDolar
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(body))

	//select {
	//case <-time.After(time.Second * 5):
	//	log.Println("Request processada com sucesso - Cmd")
	//	w.Write([]byte(body))
	//case <-ctx.Done():
	//	log.Println("Request cancelada pelo cliente - Cmd")
	//}

	//req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/get", nil)

}
