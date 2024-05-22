package main

import "fmt"

type Kavehnegar struct {
	
}

func(k *Kavehnegar)SendMessage(number int , text string)error {
	fmt.Println("sending sms with KavehNegar panel")
	fmt.Printf("to: %d" , number)
	fmt.Printf(" with text :%s\n" , text)
	return nil
}

