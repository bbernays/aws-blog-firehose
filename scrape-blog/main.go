package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
	"log"
)

func handler() (err error) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://aws.amazon.com/new/feed/")
	for _, story := range feed.Items {
		log.Print(story.Title)
		log.Print(story.Link)
		log.Print(story.GUID)

	}
	// fmt.Println(feed)
	return err
}

func main() {
	lambda.Start(handler)
}
