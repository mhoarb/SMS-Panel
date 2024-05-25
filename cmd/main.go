package main

import (
	"Sms-panel/handlers"
	"Sms-panel/handlers/services"
)

func main() {
	Kavehnegar := services.Kavehnegar{}
	alborzNegar := services.AlborzNegar{}
	mmdNegar := services.MmdNegar{}

	loggingSmsPanel := handlers.LoggingSmsPanel{
		KavehNegar:  &Kavehnegar,
		AlborzNegar: &alborzNegar,
		MmdNegar:    &mmdNegar,
	}

	err := loggingSmsPanel.SendAlborzNegarMessage(989121212, "hello")
	if err != nil {
		return
	}

	err = loggingSmsPanel.SendKavehNegarMessage(98912458, "hi")
	if err != nil {
		return
	}

	err = loggingSmsPanel.SendMmdNegarMessage(98912458, "this is mmd negar panel")
	if err != nil {
		return
	}
}
