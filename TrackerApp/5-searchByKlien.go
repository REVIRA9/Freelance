package main

import "fmt"

func CariProyekByKlien(DaftarProyek tabTrack, jumlahProyek int) {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek nih")
	} else {
		var namaKlienDicari string
	var modeCari int

	fmt.Print("Mau nyari klien apa? : ")
	fmt.Scan(&namaKlienDicari)

	fmt.Println("\nOke, mau dicari pakai cara apa nih?")
	fmt.Println("  1. Sequential (Santai, kita cari satu satu)")
	fmt.Println("  2. Binary (Lebih cepet tapi kudu urut datanya)")
	fmt.Print("Pilihanmu (1 atau 2)? : ")
	fmt.Scan(&modeCari)

	switch modeCari {
	case 1:
		fmt.Println("\nKita cari satu satu ya!")
		SeqSearchKlien(DaftarProyek, jumlahProyek, namaKlienDicari)
	case 2:
		fmt.Println("\nSip! Kita cari cepet ya!")
		fmt.Println("Cara ini cuma bener kalau daftar proyeknya DIURUTIN berdasarkan NAMA KLIEN.")
		fmt.Println("Kalau belum urut, hasil bisa salah ya.")
		BinarySearchKlien(DaftarProyek, jumlahProyek, namaKlienDicari)
	default:
		fmt.Println("Perasaan cuma ada pilihan 1 atau 2 dah")
	}
	}
}

func SeqSearchKlien(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
	apakahKetemu := false

	for i := 0; i < jumlahProyek && !apakahKetemu; i++ {
		if DaftarProyek[i].klien == namaYangDicari {
			tampilkanInfoProyekLengkap(DaftarProyek[i], i)
			apakahKetemu = true
		}
	}

	if !apakahKetemu {
		fmt.Printf("Duh, maaf nih, klien dengan nama '%s' kayaknya nggak ada di daftar.\n", namaYangDicari)
	}
}

func BinarySearchKlien(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
	indeksAwal := 0
	indeksAkhir := jumlahProyek - 1
	apakahKetemu := false

	for indeksAwal <= indeksAkhir && !apakahKetemu {
		indeksTengah := indeksAwal + (indeksAkhir-indeksAwal)/2
		
		kataKlienDiTengah := DaftarProyek[indeksTengah].klien
		kataCari := namaYangDicari
		var hasilBanding int 

		panjangKata1 := len(kataKlienDiTengah)
		panjangKata2 := len(kataCari)
		panjangMinimal := panjangKata1
		if panjangKata2 < panjangMinimal {
			panjangMinimal = panjangKata2
		}

		sudahAdaHasilBanding := false
		for k := 0; k < panjangMinimal && !sudahAdaHasilBanding; k++ {
			huruf1 := kataKlienDiTengah[k]
			huruf2 := kataCari[k]
			if huruf1 >= 'A' && huruf1 <= 'Z' {
				huruf1 += 32
			}
			if huruf2 >= 'A' && huruf2 <= 'Z' {
				huruf2 += 32
			}
			if huruf1 < huruf2 {
				hasilBanding = -1
				sudahAdaHasilBanding = true 
			} else if huruf1 > huruf2 { 
				hasilBanding = 1
				sudahAdaHasilBanding = true 
			}
		}

		if !sudahAdaHasilBanding { 
			if panjangKata1 < panjangKata2 {
				hasilBanding = -1
			} else if panjangKata1 > panjangKata2 {
				hasilBanding = 1
			} else {
				hasilBanding = 0 
			}
		}

		if hasilBanding == 0 {
			tampilkanInfoProyekLengkap(DaftarProyek[indeksTengah], indeksTengah)
			apakahKetemu = true 
		} else if hasilBanding < 0 { 
			indeksAwal = indeksTengah + 1
		} else { 
			indeksAkhir = indeksTengah - 1
		}
	}

	if !apakahKetemu {
		fmt.Printf("Sayang sekali, klien '%s' nggak ketemu.\n", namaYangDicari)
		fmt.Println("Pastikan nama klien sudah benar dan data diurutkan jika pakai Binary Search.")
	}
}
