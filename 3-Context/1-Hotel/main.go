package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// uma vez iniciado, esta rodando na nossa aplicação
	ctx := context.Background()

	// adicionar timeout de 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	//defer para rodar por ultimo
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	// quando passar os 3 segundos
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
	// se der 5 segundos e o contexto nao foi cancelado
	//case <-time.After(5 * time.Second):
	//	fmt.Println("Hotel booked.")
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
