package main

import (
	"artemisAMQlistenerAMQP/internal"
	"time"
)

func main() {
	go internal.ListenMessage(1)
	time.Sleep(100 * time.Microsecond)
	go internal.ListenMessage(2)
	time.Sleep(100 * time.Microsecond)
	go internal.ListenMessage(3)
	time.Sleep(100 * time.Microsecond)
	go internal.ListenMessage(4)
	time.Sleep(100 * time.Microsecond)
	internal.SendMessages()
	time.Sleep(3 * time.Second)
}
