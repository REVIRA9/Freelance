package main

import "fmt"

func InputData(A *tabTrack, n *int){
	var inputBaru, inputLama, status int
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

			valid := false
			for !valid{
				fmt.Println("📍 Status Proyek:")
				fmt.Println("1. Pending")
				fmt.Println("2. Sedang_dikerjakan")
				fmt.Println("3. Selesai")
				fmt.Print("Pilih Status (1-3) : ")
				fmt.Scan(&status)

				if status == 1{
					A[inputLama].status = "Pending"
					valid = true
				} else if status == 2{
					A[inputLama].status = "Sedang_dikerjakan"
					valid = true
				} else if status == 3{
					A[inputLama].status = "Selesai"
					valid = true
				} else {
					fmt.Println("❌ Pilihan tidak valid! Silakan masukkan ulang (1-3).")
				}
			}
			fmt.Print("📆 Deadline (dd mm yyyy): ")
			fmt.Scan(&A[inputLama].deadlineDay, &A[inputLama].deadlineMonth, &A[inputLama].deadlineYear)

			fmt.Print("💰 Masukan Bayaran : Rp")
			fmt.Scan(&A[inputLama].bayaran)
		inputLama++
		}
	*n = inputBaru
	fmt.Println("✅ Data proyek berhasil ditambahkan.\n\n")
	} else{
		fmt.Println("INPUTAN MELEBIHI KAPASITAS DATA❗❗\n\n")		
		}	
	} else {
		fmt.Println("DATA PROYEK SUDAH PENUH❗❗\n\n")	
	}

}

func ReadData(A tabTrack, n int){
	fmt.Println("=============================================================================================================")
	fmt.Println("|                                 APLIKASI TRACKING FREELANCE                                               |")
	fmt.Println("=============================================================================================================")	
	fmt.Printf("| %-3s | %-20s | %-15s | %-18s | %-12s | %-22s |\n", "No", "Nama Proyek", "Nama Klien", "Status", "Deadline", "Bayaran")
	fmt.Println("=============================================================================================================")
	j:=1
	for i:=0;i<n;i++{
		fmt.Printf("| %-3d | %-20s | %-15s | %-18s | %02d-%02d-%04d   | Rp%-21d|\n",j , A[i].proyek, A[i].klien, A[i].status,  A[i].deadlineDay, A[i].deadlineMonth, A[i].deadlineYear, A[i].bayaran)
		j++
	}
	fmt.Println("=============================================================================================================")
}