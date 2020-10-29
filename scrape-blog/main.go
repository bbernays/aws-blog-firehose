package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"log"
	"net/http"
)

func downloadData(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		panic(err)
	}

	return string(body), err
}

func handler() (err error) {
	fp := gofeed.NewParser()

	json_urls := []string{"https://aws.amazon.com/api/dirs/items/search?item.directoryId=blog-posts&sort_by=item.additionalFields.createdDate&sort_order=desc&size=15&item.locale=en_US"}
	rss_urls := []string{"https://aws.amazon.com/new/feed/"}
	for _, url := range rss_urls {
		feed, _ := fp.ParseURL(url)
		for _, story := range feed.Items {
			log.Print(story.Title)
			log.Print(story.Link)
			log.Print(story.GUID)
		}
	}

	for _, url := range json_urls {
		data, _ := downloadData(url)
		var blogsData blogsData
		json.Unmarshal([]byte(data), &blogsData)
		for _, story := range blogsData.Items {
			log.Print(story.Item.AdditionalFields.Link)
			log.Print(story.Item.AdditionalFields.Title)
			log.Print(story.Item.ID)
		}
	}
	return err
}

func main() {
	lambda.Start(handler)
}
