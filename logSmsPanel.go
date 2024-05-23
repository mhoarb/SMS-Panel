package main

import "fmt"

type LoggingSmsPanel struct {
	kavehNegar  *Kavehnegar
	alborzNegar *AlborzNegar
	mmdNegar    *MmdNegar
}

func (lsp *LoggingSmsPanel) SendKavehNegarMessage(number int, text string) error {
	fmt.Printf("LOG: sending SMS with kavehnegar panel to %d with text %s\n", number, text)
	lsp.kavehNegar.SendMessage(number, text)

	return nil
}

func(lsp *LoggingSmsPanel)SendAlborzNegarMessage(number int , text string) error {
	fmt.Printf("LOG: sending SMS with Alborznegar panel to %d with text %s\n", number, text)
	lsp.alborzNegar.SendMessage(number ,text)
	return nil 
}

func (lsp *LoggingSmsPanel) SendMmdNegarMessage(number int, text string) error {
	fmt.Printf("LOG: sending SMS with Mmdnegar panel to %d with text %s\n", number, text)
	lsp.kavehNegar.SendMessage(number, text)

	return nil
}


