package services

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	// "golang.org/x/text/message"
)

type MessageQueue chan string

type Config struct {
	URL           string
    FetchInterval time.Duration
    MessageQueue MessageQueue
}

func FetchTitle(config Config) {
	tick := time.NewTicker(config.FetchInterval)

	for range tick.C {
		doc, err := goquery.NewDocument(config.URL)
		if err != nil {
			fmt.Printf("Error fetching data from %s: %v\n", config.URL, err)
			continue
		}

		title := doc.Find("title").Text()
		message := fmt.Sprintf("Title from %s: %s\n", config.URL, time.Now().Format("05"), title)
		config.MessageQueue <- message
	}
}

func FetchData(config Config) {
	for {
		message := <-config.MessageQueue
		data := fetchDataFromMessage(config.URL)
		if data != "" {
			fmt.Printf(message+"Data from %s: %s\n", config.URL, time.Now().Format("05"), data)
		}
	}
}

func fetchDataFromMessage(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("Error fetching data from %s: %v\n", url, err)
		return ""
	}
	return doc.Find("h1").Text()
}
