package main

import "fmt"

func UpdateStatus(A *tabTrack, n int) {
	var nama string
	var status int
	var valid bool = false
	fmt.Print("Masukkan nama proyek yang ingin diubah statusnya: ")
	fmt.Scan(&nama)

	found := false

	for i := 0; i < n; i++ {
		if !found && len(nama) == len(A[i].proyek) {
			sama := true
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
					sama = false
				}
			}

			if sama {
				fmt.Printf("Status saat ini: %s\n", A[i].status)
				for !valid{
				fmt.Println("Ubah Status Proyek:")
				fmt.Println("1. Pending")
				fmt.Println("2. Sedang_dikerjakan")
				fmt.Println("3. Selesai")
				fmt.Print("Pilih Status (1-3) : ")
				fmt.Scan(&status)

				if status == 1{
					A[i].status = "Pending"
					valid = true
				} else if status == 2{
					A[i].status = "Sedang_dikerjakan"
					valid = true
				} else if status == 3{
					A[i].status = "Selesai"
					valid = true
				} else {
					fmt.Println("❌ Pilihan tidak valid! Silakan masukkan ulang (1-3).")
				}
			}
				fmt.Println("✅ Status proyek telah diperbarui.")
				found = true
			}
		}
	}

	if !found {
		fmt.Println("❌ Proyek tidak ditemukan.")
	}
}



func HapusData(A *tabTrack, n *int) {
	var nama string
	fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
	fmt.Scan(&nama)

	indeksDihapus := -1

	for i := 0; i < *n; i++ {
		match := true

		if len(nama) == len(A[i].proyek) {
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
				}
			}
		} else {
			match = false
		}

		if indeksDihapus == -1 && match {
			indeksDihapus = i
		}
	}

	if indeksDihapus != -1 {
		for j := indeksDihapus; j < *n-1; j++ {
			A[j] = A[j+1]
		}
		*n--
		fmt.Println("✅ Proyek telah berhasil dihapus.")
	} else {
		fmt.Println("❌ Proyek tidak ditemukan.")
	}
}