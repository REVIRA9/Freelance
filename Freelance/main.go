package main

import (
	"Freelance/function"
	"fmt"
)



func main(){
	var output int

	for output != 0{
		function.Menu()
		fmt.Scan(&output)
		switch output{
		case 1:
		case 2:
		default : fmt.Println("Opsi tidak valid, silahkan masukan ulang opsi: ")
		}
	}
}