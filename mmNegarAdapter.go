package main

import "fmt"

type MmdNegarAdapter struct {
	mmdNegar *MmdNegar
}

func (m *MmdNegarAdapter) SendMessage() {
	fmt.Println("sending message from mmdNegar panel")
	m.mmdNegar.SendMessage(
	
}