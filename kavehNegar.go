package main
/**
Kavehnegar is a sms panel  that have a method SendMessage
but its call by logging system
**/
type Kavehnegar struct {
}

func (k *Kavehnegar) SendMessage(number int, text string) error {
	// fmt.Println("sending sms with KavehNegar panel")
	// fmt.Printf("to: %d" , number)
	// fmt.Printf(" with text :%s\n" , text)
	return nil
}
