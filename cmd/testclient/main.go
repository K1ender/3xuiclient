package main

import (
	"context"
	"fmt"
	"os"
	"time"

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
			ExpiryTime: time.Now().Add(time.Hour).UnixMilli(),
			TotalGB:    1024 * 1024 * 1024,
			LimitIP:    1000,
			Email:      "test@test.test",
			Enable:     false,
		},
		InboundIDS: []int64{
			id,
		},
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	err = client.UpdateClient(context.Background(), "test@test.test", xuiclient.UpdateClientRequest{
		CreateClient: xuiclient.CreateClient{
			Email:      "test@test.test",
			ExpiryTime: time.Now().Add(time.Hour * 24).UnixMilli(),
			Enable:     true,
			TotalGB:    1024 * 1024 * 1024 * 10,
		},
	})
	if err != nil {
		panic(err)
	}

	clientRes, err := client.GetClient(context.Background(), "test@test.test")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", clientRes.Data)
}
