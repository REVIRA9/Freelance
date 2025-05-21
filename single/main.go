package main

import "fmt"

const NMAX int = 1024

type tracking struct {
	proyek        string
	klien, status string
	deadlineDay, deadlineMonth, deadlineYear      int
}

type tabTrack [NMAX]tracking

func main(){
	var Data tabTrack
	var nData int
	var output int

	for output != 9{
		Menu()
		fmt.Print("Masukan Opsi: ")
		fmt.Scan(&output)
		switch output{
		case 1: InputData(&Data, &nData)
		case 2:	ReadData(Data, nData)
		default : fmt.Println("Opsi tidak valid, silahkan masukan ulang opsi: ")
		}
	}
}

func Menu(){
	fmt.Println("Menu")
	fmt.Println("1- Input Data")
	fmt.Println("2- List Data")
}

func InputData(A *tabTrack, n *int){
	var inputBaru, inputLama int
	fmt.Print("Masukan Jumlah Proyek yang ingin di input: ")
	fmt.Scan(&inputBaru)
	inputBaru += *n
	inputLama += *n
	
	for inputLama<inputBaru{
		fmt.Print("Masukana nama proyek: ")
		fmt.Scan(&A[inputLama].proyek)
		fmt.Print("Masukana nama klien: ")
		fmt.Scan(&A[inputLama].klien) 
		fmt.Print("Masukana status: ")
		fmt.Scan(&A[inputLama].status)
		fmt.Print("Masukan Deadline: ")
		fmt.Scan(&A[inputLama].deadlineDay, &A[inputLama].deadlineMonth, &A[inputLama].deadlineYear)
		inputLama++
	}
	*n = inputBaru
	fmt.Println("Proyek berhasil ditambahkan")
	
}

func ReadData(A tabTrack, n int){
	fmt.Println("\t LIST-LIST PROYEK")
	fmt.Printf("%-3s %-15s %-15s %-10s %-4s\n", "No", "Proyek", "Klien", "Status", "Deadline")
	j:=1
	for i:=0;i<n;i++{
		fmt.Printf("%-3d %-15s %-15s %-10s %d-%d-%d\n",j , A[i].proyek, A[i].klien, A[i].status,  A[i].deadlineDay, A[i].deadlineMonth, A[i].deadlineYear)
		j++
	}
	fmt.Println()
}