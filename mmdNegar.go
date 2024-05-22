package main

import "fmt"

type MmdNegar struct {
}

func (m *MmdNegar) SendMessage(number int, text string) error {
	fmt.Println("sending sms with Mmdnegar sms panel")
	fmt.Printf("to: %d " , number)
	fmt.Printf(" with text :%s\n" , text)

	return nil
}