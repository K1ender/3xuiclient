package main

import (
	"context"
	"os"

	xuiclient "github.com/k1ender/3xuiclient"
)

func main() {
	baseurl := os.Getenv("XUI_BASEURL")
	token := os.Getenv("XUI_TOKEN")

	client := xuiclient.New(baseurl, token)

	res, err := client.GetInbounds(context.Background())
	if err != nil {
		panic(err)
	}

	id := res.Data[0].ID

	err = client.AddClient(context.Background(), xuiclient.CreateClientRequest{
		Client: xuiclient.CreateClient{
			Email: "test",
		},
		InboundIDS: []int64{
			id,
		},
	})
	if err != nil {
		panic(err)
	}
}
