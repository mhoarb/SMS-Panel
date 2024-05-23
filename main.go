package main

func main() {
	Kavehnegar := Kavehnegar{}
	alborzNegar := AlborzNegar{}
	mmdNegar := MmdNegar{}

	loggingSmsPanel := LoggingSmsPanel{
		kavehNegar: &Kavehnegar,
		alborzNegar: &alborzNegar,
		mmdNegar: &mmdNegar,
	}

	loggingSmsPanel.SendAlborzNegarMessage(989121212 , "hello")

	loggingSmsPanel.SendKavehNegarMessage(98912458 , "hi")

	loggingSmsPanel.SendMmdNegarMessage(98912458 , "this is mmd negar panel")
}

