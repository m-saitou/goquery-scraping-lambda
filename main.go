package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	URL string `json:"url"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
	url := event.URL

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	doc.Find("p").Each(func(_ int, s *goquery.Selection) {
		span := s.Find("span").First()
		fmt.Println(span.Text())

		a := s.Find("a").First()
		href, _ := a.Attr("href")
		fmt.Println(href)
	})

	return MyResponse{Message: "success"}, nil
}

func main() {
	lambda.Start(hello)
}
