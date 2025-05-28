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

func CariProyek(DaftarProyek tabTrack, jumlahProyek int) {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek nih")
		return
	}

	var namaProyekDicari string
	var modeCari int

	fmt.Print("Mau nyari proyek apa?")
	fmt.Scan(&namaProyekDicari)

	fmt.Println("\nOke, mau dicari pakai cara apa nih?")
	fmt.Println("  1. Sequential (Santai, kita cari satu satu)")
	fmt.Println("  2. Binary (Lebih cepet tapi kudu urut datanya)")
	fmt.Print("Pilihanmu (1 atau 2)? ")
	fmt.Scan(&modeCari)

	switch modeCari {
	case 1:
		fmt.Println("\nKita cari satu satu ya!")
		cariSatuSatu(DaftarProyek, jumlahProyek, namaProyekDicari)
	case 2:
		fmt.Println("\nSip! Kita cari cepet ya!")
		fmt.Println("Cara ini cuma bener kalau daftar proyeknya DIURUTIN kyk NAMA PROYEK.")
		fmt.Println("Kalau belum urut ntar ga akurat, mending urutin dulu dah di opsi menu no 6")
		cariBagiDua(DaftarProyek, jumlahProyek, namaProyekDicari)
	default:
		fmt.Println("Perasaan cuma ada pilihan 1 atau 2 dah")
	}
}

func cariSatuSatu(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
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

func cariBagiDua(DaftarProyek tabTrack, jumlahProyek int, namaYangDicari string) {
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
		fmt.Println("   Pastikan namanya bener dan datanya udah diurutin ya kalau mau pakai cara ini.")
	}
}

func tampilkanInfoProyekLengkap(p tracking, nomorData int) {
	fmt.Println("-------------------------------------------")
	fmt.Printf("Proyeknya ketemu nih di data nomor %d \n", nomorData+1)
	fmt.Printf("   Nama Proyek    : %s\n", p.proyek)
	fmt.Printf("   Nama Klien     : %s\n", p.klien)
	fmt.Printf("   Status Saat Ini : %s\n", p.status)
	tanggalDeadline := fmt.Sprintf("%02d-%02d-%04d", p.deadlineDay, p.deadlineMonth, p.deadlineYear)
	fmt.Printf("   Deadline       : %s\n", tanggalDeadline)
	fmt.Printf("   Perkiraan Bayaran: Rp %d\n", p.bayaran)
	fmt.Println("-------------------------------------------")
}

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

func UrutkanProyek(A *tabTrack, n int) {
	if n == 0 {
		fmt.Println("Proyek masih kosong")
		return
	}
	if n == 1 {
		fmt.Println("Cuman ada satu proyek, belum bisa di urutin")
		return
	}

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
	
	for i := 0; i < n-1; i++ {
		indeksPilihan := i
		
		for j := i + 1; j < n; j++ {
			JDulu := false

			if pilihanKriteria == 1 {
				if pilihanUrutan == 1 {
					JDulu = isD1SebelumD2(A[j], A[indeksPilihan])
				} else {
					JDulu = isD1SebelumD2(A[indeksPilihan], A[j])
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
	ReadData(*A, n) 
}
