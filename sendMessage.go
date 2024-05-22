package main

type SendMessage interface{
	SendMessage(number int , text string)error
}

