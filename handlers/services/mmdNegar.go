package services

/*
*
Mmd negar is a sms panel like Kavehnager that have a method SendMessage
but its call by logging system
*
*/
type MmdNegar struct {
}

func (m *MmdNegar) SendMessage(number int, text string) error {
	// fmt.Println("sending sms with Mmdnegar sms panel")
	// fmt.Printf("to: %d " , number)
	// fmt.Printf(" with text :%s\n" , text)

	return nil
}
