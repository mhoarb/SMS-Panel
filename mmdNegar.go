package main

import "fmt"

type MmdNegar struct {
}


func(mn *MmdNegar) SendMessage() {
	fmt.Println("sending message from Mmd negar panel")
}
