package main

import "fmt"

func main() {
	var Data tabTrack
	var nData int
	var Option int

	for Option != 7 {
		Menu()
		fmt.Print("Masukan Opsi: ")
		fmt.Scan(&Option)
		switch Option {
		case 1:
			InputData(&Data, &nData)
		case 2:
			ReadData(Data, nData)
		case 3:
			UpdateStatus(&Data, nData)
		case 4:
			HapusData(&Data, &nData)
		case 5:
			CariProyek(Data, nData)
		case 6:
			UrutkanProyek(&Data, &nData)
		case 7:
			fmt.Println("Program Selesai")
		default:
			fmt.Println("Opsi tidak valid, silahkan masukan ulang.")
		}
	}
}
