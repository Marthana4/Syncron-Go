package main

import (
    "github.com/Marthana4/gomessage/services"
    "time"
)

func main() {
    messageQueue := make(services.MessageQueue)

	titleConfig := services.Config{
		URL:           "https://go.dev/doc/",
		FetchInterval: 5 * time.Second,
		MessageQueue:  messageQueue,
	}

	dataConfig := services.Config{
		URL:           "https://go.dev/doc/",
		FetchInterval: 5 * time.Second,
		MessageQueue:  messageQueue,
	}


	go services.FetchTitle(titleConfig)
	go services.FetchData(dataConfig)

    select {}
}
