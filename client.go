package main

import "fmt"

type iClient interface {
	SendMessage() error
}
type Client struct {
}

func (c *Client) SensMessage(client iClient) error {
	fmt.Println("sending message from client")
	client.SendMessage()
	return nil
}
