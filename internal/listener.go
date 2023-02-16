package internal

import (
	"context"
	"fmt"
	"github.com/Azure/go-amqp"
	"log"
)

func ListenMessage(numberListener int) {

	ctx := context.Background()
	addr := "amqp://artemis:simetraehcapa@localhost:61616/"

	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatal("Dialing AMQP server failed: ", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("could not create AMQP session:", err)
	}
	defer session.Close(ctx)

	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress("/queue/myqueue"),
	)
	if err != nil {
		log.Fatal("could not create AMQP receiver:", err)
	}
	defer receiver.Close(ctx)

	for {
		// Receive the next message
		message, err := receiver.Receive(ctx)
		if err != nil {
			log.Fatal("could not receive AMQP message:", err)
		}
		// Accept the message
		message.Accept(ctx)

		// Print the message body
		fmt.Println(numberListener, " Received message: ", bytesToString(message.Data))
	}
}

func bytesToString(bytes [][]byte) string {
	var merged []byte
	for _, b := range bytes {
		merged = append(merged, b...)
	}
	return string(merged)
}
