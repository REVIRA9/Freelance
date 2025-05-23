package main

import "fmt"

const NMAX int = 3

type tracking struct {
	proyek        string
	klien, status string
	deadlineDay, deadlineMonth, deadlineYear int
	bayaran int
}

type tabTrack [NMAX]tracking

func main(){
	var Data tabTrack
	var nData int
	var Option int

	for Option != 9{
		Menu()
		fmt.Print("Masukan Opsi: ")
		fmt.Scan(&Option)
		switch Option{
		case 1: InputData(&Data, &nData)
		case 2:	ReadData(Data, nData)
		case 9: fmt.Println("Program Selesai")
		default : fmt.Println("Opsi tidak valid, silahkan masukan ulang opsi: ")
		}
	}
}

func Menu(){
	fmt.Println("===============================================")
	fmt.Println("|         APLIKASI TRACKING FREELANCE         |")
	fmt.Println("===============================================")	
	fmt.Println("| 1. Input Data Proyek                        |")
	fmt.Println("| 2. Tampilkan Daftar Proyek                  |")
	fmt.Println("| 3. Ubah Status Proyek                       |")
	fmt.Println("| 4. Hapus Proyek                             |")
	fmt.Println("| 5. Cari Proyek (Sequential / Binary Search) |")
	fmt.Println("| 6. Urutkan Proyek (Deadline / Bayaran)      |")
	fmt.Println("| 9. Keluar                                   |")
	fmt.Println("===============================================")
}

func InputData(A *tabTrack, n *int){
	var inputBaru, inputLama int
	if *n < NMAX{
	fmt.Print("Masukan Jumlah Proyek yang ingin di input: ")
	fmt.Scan(&inputBaru)
	inputBaru += *n
	inputLama += *n

	if inputBaru <=NMAX{
		for inputLama<inputBaru{
			fmt.Println("========================================")
			fmt.Printf("Input Proyek ke-%d\n", inputLama+1)
			fmt.Println("========================================")

			fmt.Print("📌 Nama Proyek   : ")
			fmt.Scan(&A[inputLama].proyek)

			fmt.Print("👤 Nama Klien    : ")
			fmt.Scan(&A[inputLama].klien)

			fmt.Print("📍 Status        : ")
			fmt.Scan(&A[inputLama].status)

			fmt.Print("📆 Deadline (dd mm yyyy): ")
			fmt.Scan(&A[inputLama].deadlineDay, &A[inputLama].deadlineMonth, &A[inputLama].deadlineYear)
		inputLama++
	}
	*n = inputBaru
	fmt.Println("✅ Data proyek berhasil ditambahkan.")
	} else{
		fmt.Println("INPUTAN MELEBIHI KAPASITAS DATA❗❗")		
		}	
	} else {
		fmt.Println("DATA PROYEK SUDAH PENUH❗❗")	
	}

}

func ReadData(A tabTrack, n int){
	fmt.Println("==========================================================================================")
	fmt.Printf("| %-3s | %-20s | %-15s | %-10s | %-12s | %-8s |\n", "No", "Nama Proyek", "Nama Klien", "Status", "Deadline", "Bayaran")
	fmt.Println("==========================================================================================")
	j:=1
	for i:=0;i<n;i++{
		fmt.Printf("| %-3d | %-20s | %-15s | %-10s | %d-%d-%-8d |\n",j , A[i].proyek, A[i].klien, A[i].status,  A[i].deadlineDay, A[i].deadlineMonth, A[i].deadlineYear)
		j++
	}
	fmt.Println("==========================================================================================")
}

func UbahData(A *tabTrack, n int){
	var nama string
	fmt.Print("Masukkan nama proyek yang ingin diubah statusnya: ")
	fmt.Scan(&nama)

	found := false
	for i := 0; i < n; i++ {
		if len(nama) == len(A[i].proyek) {
			sama := true
			for j := 0; j < len(nama); j++ {
				a := nama[j]
				b := A[i].proyek[j]
				if a >= 'A' && a <= 'Z' {
					a += 64
				}
				if b >= 'A' && b <= 'Z' {
					b += 64
				}
				if a != b {
					sama = false
					break
				}
			}
			if sama {
				fmt.Printf("Status saat ini: %s\n", A[i].status)
				fmt.Print("Masukkan status baru: ")
				fmt.Scan(&A[i].status)
				fmt.Println("Status proyek telah diperbarui.")
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func hapusData(A *tabTrack, n *int) {
	var nama string
	fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i := 0; i < *n; i++ {
		match := true
		if len(nama) != len(A[i].proyek) {
			match = false
		} else {
			for j := 0; j < len(nama); j++ {
				a := nama[j]
				b := A[i].proyek[j]

				if a >= 'A' && a <= 'Z' {
					a += 32
				}
				if b >= 'A' && b <= 'Z' {
					b += 32
				}
				if a != b {
					match = false
					break
				}
			}
		}

		if match {
			
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("Proyek telah berhasil dihapus.")
			return
		}
	}

	fmt.Println("Proyek tidak ditemukan.")
}


