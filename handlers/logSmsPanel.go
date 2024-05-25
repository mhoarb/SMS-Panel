package handlers

import (
	"Sms-panel/handlers/services"
	"fmt"
)

/*
*
logging sms panel is a struct and have all my panels
*/
type LoggingSmsPanel struct {
	KavehNegar  *services.Kavehnegar
	AlborzNegar *services.AlborzNegar
	MmdNegar    *services.MmdNegar
}

func (lsp *LoggingSmsPanel) SendKavehNegarMessage(number int, text string) error {
	fmt.Printf("LOG: sending SMS with kavehnegar panel to %d with text %s\n", number, text)
	lsp.KavehNegar.SendMessage(number, text)

	return nil
}

func (lsp *LoggingSmsPanel) SendAlborzNegarMessage(number int, text string) error {
	fmt.Printf("LOG: sending SMS with Alborznegar panel to %d with text %s\n", number, text)
	lsp.AlborzNegar.SendMessage(number, text)
	return nil
}

func (lsp *LoggingSmsPanel) SendMmdNegarMessage(number int, text string) error {
	fmt.Printf("LOG: sending SMS with Mmdnegar panel to %d with text %s\n", number, text)
	lsp.KavehNegar.SendMessage(number, text)

	return nil
}
