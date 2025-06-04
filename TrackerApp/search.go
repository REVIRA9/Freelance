package main

import "fmt"

func CariProyek(DaftarProyek tabTrack, jumlahProyek int) {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek nih")
		return	
	}

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
		BinarySearch(DaftarProyek, jumlahProyek, namaProyekDicari)
	default:
		fmt.Println("Perasaan cuma ada pilihan 1 atau 2 dah")
	}
}

func SeqSearch(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
	apakahKetemu := false
	for i := 0; i < jumlahProyek && !apakahKetemu; i++ {
		kataProyek := DaftarProyek[i].proyek
		kataCari := namaYangDicari
		kondisiCocok := false

		if len(kataProyek) == len(kataCari) {
			semuaHurufSama := true
			for j := 0; j < len(kataProyek) && semuaHurufSama; j++ {
				huruf1 := kataProyek[j]
				huruf2 := kataCari[j]
				if huruf1 >= 'A' && huruf1 <= 'Z' {
					huruf1 += 32
				}
				if huruf2 >= 'A' && huruf2 <= 'Z' {
					huruf2 += 32
				}
				if huruf1 != huruf2 {
					semuaHurufSama = false
				}
			}
			if semuaHurufSama {
				kondisiCocok = true
			}
		}

		if kondisiCocok {
			tampilkanInfoProyekLengkap(DaftarProyek[i], i)
			apakahKetemu = true
		}
	}

	if !apakahKetemu {
		fmt.Printf("Duh, maaf nih, proyek dengan nama '%s' kayaknya nggak ada di daftar.\n", namaYangDicari)
	}
}

func BinarySearch(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
	indeksAwal := 0
	indeksAkhir := jumlahProyek - 1
	apakahKetemu := false

	for indeksAwal <= indeksAkhir && !apakahKetemu {
		indeksTengah := indeksAwal + (indeksAkhir-indeksAwal)/2
		
		kataProyekDiTengah := DaftarProyek[indeksTengah].proyek
		kataCari := namaYangDicari
		var hasilBanding int 

		panjangKata1 := len(kataProyekDiTengah)
		panjangKata2 := len(kataCari)
		panjangMinimal := panjangKata1
		if panjangKata2 < panjangMinimal {
			panjangMinimal = panjangKata2
		}

		sudahAdaHasilBanding := false
		for k := 0; k < panjangMinimal && !sudahAdaHasilBanding; k++ {
			huruf1 := kataProyekDiTengah[k]
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
		fmt.Printf("Sayang sekali, setelah dicari pakai cara bagi dua, proyek '%s' tetap nggak ketemu.\n", namaYangDicari)
		fmt.Println("Pastikan namanya bener dan datanya udah diurutin ya kalau mau pakai cara ini.")
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
