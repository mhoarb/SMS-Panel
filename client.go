package main

import "fmt"

type iClient interface {
	SendMessage()
}

type Client struct {
}

func (c *Client) SendClientMessage(ic iClient) error {
	fmt.Println("client Ssend message to ...")
	ic.SendMessage()
	return nil
}