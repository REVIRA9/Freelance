package main

import "fmt"

func cekUrut(p1, p2 tracking) bool {
	if p1.deadlineYear < p2.deadlineYear {
		return true
	} else if p1.deadlineYear > p2.deadlineYear {
		return false
	} else {
		if p1.deadlineMonth < p2.deadlineMonth {
			return true
		} else if p1.deadlineMonth > p2.deadlineMonth {
			return false
		} else {
			if p1.deadlineDay < p2.deadlineDay {
				return true
			} else {
				return false
			}
		}
	}
}

func UrutkanProyek(A *tabTrack, n *int) {
	if *n == 0 {
		fmt.Println("Proyek masih kosong")
	} else if *n==1{
		fmt.Println("Cuman ada satu proyek, belum bisa di urutin")
	} else{
	var pilihanKriteria, pilihanUrutan int
	fmt.Println("--- Urutkan Proyek Berdasarkan Apa? ---")
	fmt.Println("1. Deadline (Tanggal Tenggat Waktu)")
	fmt.Println("2. Bayaran (Nominal)")
	fmt.Print("Pilih kriteria (1-2): ")
	fmt.Scan(&pilihanKriteria)

	fmt.Println("\n--- Mau Diurutkan Bagaimana? ---")
	fmt.Println("1. Menaik (Deadline Terawal ke Terakhir)")
	fmt.Println("2. Menurun (Deadline Terakhir ke Terawal)")
	fmt.Print("Pilih urutan (1-2): ")
	fmt.Scan(&pilihanUrutan)
	
	for i := 0; i < *n-1; i++ {
		indeksPilihan := i
		
		for j := i + 1; j < *n; j++ {
			JDulu := false

			if pilihanKriteria == 1 {
				if pilihanUrutan == 1 {
					JDulu = cekUrut(A[j], A[indeksPilihan])
				} else {
					JDulu = cekUrut(A[indeksPilihan], A[j])
				}
			} else {
				if pilihanUrutan == 1 {
					JDulu = A[j].bayaran < A[indeksPilihan].bayaran
				} else {
					JDulu = A[j].bayaran > A[indeksPilihan].bayaran
				}
			}

			if JDulu {
				indeksPilihan = j
			}
		}

		if indeksPilihan != i {
			A[i], A[indeksPilihan] = A[indeksPilihan], A[i]
		}
	}

	fmt.Println("\nData proyek berhasil diurutkan!")
	ReadData(*A, *n) 
	}
}