package internal

import (
	"context"
	"github.com/Azure/go-amqp"
	"log"
	"strconv"
)

func SendMessages() {
	for i := 1; i <= 20; i++ {
		message := "test " + strconv.Itoa(i)
		sendMessage(message)
	}
}

func sendMessage(messageInput string) {

	ctx := context.Background()

	conn, err := amqp.Dial("amqp://artemis:simetraehcapa@localhost:61616/")
	if err != nil {
		log.Fatal("could notconnect  to AMQP server:", err)
	}
	defer conn.Close()

	// Create a new AMQP session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("could not create AMQP session:", err)
	}
	defer session.Close(ctx)

	// Create a new AMQP sender
	sender, err := session.NewSender(
		amqp.LinkTargetAddress("/queue/myqueue"),
	)
	if err != nil {
		log.Fatal("could not create AMQP sender:", err)
	}
	defer sender.Close(ctx)

	// Create a new AMQP message
	message := amqp.NewMessage([]byte(messageInput))
	log.Println("message successful")

	// Send the message
	err = sender.Send(ctx, message)
	if err != nil {
		log.Fatal("could not send AMQP message:", err)
	}

}
