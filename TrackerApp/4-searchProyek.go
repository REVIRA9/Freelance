package main

import "fmt"

func CariProyek(DaftarProyek tabTrack, jumlahProyek int) {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek nih")
	} else{
			var namaProyekDicari string
	var modeCari int

	fmt.Print("Mau nyari proyek apa? : ")
	fmt.Scan(&namaProyekDicari)

	fmt.Println("\nOke, mau dicari pakai cara apa nih?")
	fmt.Println("  1. Sequential (Santai, kita cari satu satu)")
	fmt.Println("  2. Binary (Lebih cepet tapi kudu urut datanya)")
	fmt.Print("Pilihanmu (1 atau 2)? : ")
	fmt.Scan(&modeCari)

	switch modeCari {
	case 1:
		fmt.Println("\nKita cari satu satu ya!")
		SeqSearch(DaftarProyek, jumlahProyek, namaProyekDicari)
	case 2:
		fmt.Println("\nSip! Kita cari cepet ya!")
		fmt.Println("Cara ini cuma bener kalau daftar proyeknya DIURUTIN kyk NAMA PROYEK.")
		fmt.Println("Kalau belum urut ntar ga akurat, mending urutin dulu dah di opsi menu no 6")
		idTarget :=-1
		for i := 0; i < jumlahProyek; i++ {
			if DaftarProyek[i].proyek == namaProyekDicari {
				idTarget = DaftarProyek[i].id
			}
		}

		if idTarget == -1 {
			fmt.Println("❌ Proyek tidak ditemukan.")
			return
		}

		UrutkanBerdasarkanID(&DaftarProyek, jumlahProyek)
		BinarySearchByID(DaftarProyek, jumlahProyek, idTarget)
	default:
		fmt.Println("Perasaan cuma ada pilihan 1 atau 2 dah")
	}
	}
}

func SeqSearch(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
	var	apakahKetemu bool = false

	for i := 0; i < jumlahProyek && !apakahKetemu; i++ {
		if DaftarProyek[i].proyek == namaYangDicari {
			tampilkanInfoProyekLengkap(DaftarProyek[i], i)
			apakahKetemu = true
		}
	}

	if !apakahKetemu {
		fmt.Printf("Duh, maaf nih, proyek dengan nama '%s' kayaknya nggak ada di daftar.\n", namaYangDicari)
	}
}

func BinarySearchByID(DaftarProyek tabTrack, jumlahProyek int, idTarget int) {
	kiri := 0
	kanan := jumlahProyek - 1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if DaftarProyek[tengah].id == idTarget {
			tampilkanInfoProyekLengkap(DaftarProyek[tengah], tengah)
			ketemu = true
		} else if DaftarProyek[tengah].id < idTarget {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if !ketemu {
		fmt.Println("❌ Proyek tidak ditemukan setelah disortir berdasarkan ID.")
	}
}


func tampilkanInfoProyekLengkap(p tracking, nomorData int) {
	fmt.Println("==============================================")
	fmt.Printf("Data ke           : %d\n", nomorData+1)
	fmt.Printf("Nama Proyek       : %s\n", p.proyek)
	fmt.Printf("Nama Klien        : %s\n", p.klien)
	fmt.Printf("Status Proyek     : %s\n", p.status)
	tanggalDeadline := fmt.Sprintf("%02d-%02d-%04d", p.deadlineDay, p.deadlineMonth, p.deadlineYear)
	fmt.Printf("Deadline          : %s\n", tanggalDeadline)
	fmt.Printf("Perkiraan Bayaran : Rp %d\n", p.bayaran)
	fmt.Println("==============================================")
}


func UrutkanBerdasarkanID(A *tabTrack, n int) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].id < A[idx].id {
				A[j], A[idx] = A[idx], A[j]
			}
		}
	}
}
