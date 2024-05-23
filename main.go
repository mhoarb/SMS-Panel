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

	loggingSmsPanel.SendAlborzNegarMessage(1020 , "hello")
}
