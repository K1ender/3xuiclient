package main

import (
	"context"
	"fmt"
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

	fmt.Printf("%+v\n", res.Data[2].Protocol)
}
