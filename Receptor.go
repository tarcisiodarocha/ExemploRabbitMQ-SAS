// Exemplo de receptor RabbitMQ em Go
package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	
	// Estabelecendo da conexao com as credenciais do administrador 
	conn, err := amqp.Dial("amqp://usuario:senha@hostname:5672/")
	failOnError(err, "Falha de conexao com RabbitMQ")
	defer conn.Close()

    // Obtendo canal de comunicacao 
	ch, err := conn.Channel()
	failOnError(err, "Falha ao obter canal")
	defer ch.Close()

    // Declarando uma nova fila chamada "sas" 
	q, err := ch.QueueDeclare(
		"sas", // nome da fila
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

    // Definindo um consumidor para a fila "sas"
	msgs, err := ch.Consume(
		q.Name, // nome da fila
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

    // Consumindo todas as mensagens da fila SAS
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" Pressione CTRL+C para sair!")
	<-forever
}