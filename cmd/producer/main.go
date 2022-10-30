package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"math/rand"
)

type Order struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
}

func GenerateOrder() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
	}
}

func Notify(ch *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct", // exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 100000; i++ {
		order := GenerateOrder()
		err = Notify(ch, order)
		if err != nil {
			panic(err)
		}
		fmt.Println("Order created:", order)
	}

}
