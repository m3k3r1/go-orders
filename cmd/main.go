package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/m3k3r1/go-orders/internal/order/infra/database"
	"github.com/m3k3r1/go-orders/internal/usecase"
	"github.com/m3k3r1/go-orders/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"math/rand"
	"sync"
)

func main() {
	maxWorkers := 3
	wg := sync.WaitGroup{}
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go Worker(out, uc, i)
	}
	wg.Wait()
}

func Worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}
		input.Tax = rand.Float64() * 8
		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("Error saving order", err)
		}
		msg.Ack(false)
		fmt.Println("Worker", workerId, "processed order", input.ID)
	}
}
