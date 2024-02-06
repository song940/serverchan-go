package main

import (
	"log"

	"github.com/song940/serverchan-go/serverchan"
)

func main() {
	client := serverchan.New(&serverchan.Config{
		ApiKey: "SCT82268Td2xus2Ms6YDvKbVIqaNNELzx",
	})
	resp, err := client.Send(&serverchan.Message{
		Title: "hello",
		Desp:  "hahah",
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
