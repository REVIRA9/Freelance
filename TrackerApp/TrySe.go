package main

//Bhian
/*func BhiSeqSearch(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
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
}*/