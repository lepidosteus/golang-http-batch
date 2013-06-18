package main

import (
	"github.com/lepidosteus/golang-http-batch/batch"
	"fmt"
)

func main() {
	b := batch.New()

	b.SetMaxConcurrent(8)

	b.AddEntry("http://www.google.com", func (url string, body string, data batch.CallbackData, err error) {
		fmt.Printf("Result from: %s\n", url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Body's length: %d\n", len(body))
	})

	b.AddEntry("http://www.aol.com", func (url string, body string, data batch.CallbackData, err error) {
		fmt.Printf("Result from: %s\n", url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Body's length: %d\n", len(body))
	})

	b.AddEntry("http://www.some-error-domain-that-fail.com", func (url string, body string, data batch.CallbackData, err error) {
		fmt.Printf("Result from: %s\n", url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Body's length: %d\n", len(body))
	})

	b.AddEntry("http://www.reddit.com", func (url string, body string, data batch.CallbackData, err error) {
		fmt.Printf("Result from: %s\n", url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Body's length: %d\n", len(body))
	})

	b.AddEntryWithData("http://www.google.com", func (url string, body string, data batch.CallbackData, err error) {
		fmt.Printf("Result from: %s\n", url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("data[foo]: %s, Body's length: %d\n", data["foo"].(string), len(body))
	}, batch.CallbackData{
		"foo": "bar",
	})

	b.Run()
}
