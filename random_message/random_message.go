package random_message

import (
	"math/rand"
)

var listOfMessages = []string{
	"Hello",
	"Hello World",
	"Hello Wts",
	"WATT",
	"Change requests user",
	"Exit in circle",
	"May",
}

func RandMessage() string {

	index := rand.Intn(len(listOfMessages))
	message := listOfMessages[index]
	return message
}
