package main
/**
sending message interface that define a function " SendMessage(int, string)error "
**/
type SendMessage interface{
	SendMessage(int ,string)error
}

