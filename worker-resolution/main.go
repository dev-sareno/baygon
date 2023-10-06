package main

import (
	"fmt"
	"github.com/dev-sareno/ginamus/worker-resolution/dns"
	"github.com/dev-sareno/ginamus/worker-resolution/mq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
)

func main() {
	// setup RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer func() {
		if err := ch.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	mq.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func test() {
	domain := "github.com"

	ipResolver := dns.IpResolver{}
	cnameResolver := dns.CnameResolver{Child: &ipResolver}
	recordResolver := dns.RecordResolver{Child: &cnameResolver}
	recordResolver.SetValue(domain)
	result, err := recordResolver.Resolve()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var items []string
	for _, item := range result {
		items = append(items, item.Value)
	}
	fmt.Printf("%v\n", strings.Join(items, "\n"))
}
