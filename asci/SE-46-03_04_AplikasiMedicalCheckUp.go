package main

import "fmt"

type pasien struct {
	nama, id, golonganDarah      string
	tglLahir, blnLahir, thnLahir int
	jumlahMCUPasien              int
}
type dataPasien [200]pasien
type Mcu struct {
	namaPasienMCU, idPasien, paketMcu, rekapMCU string
	tglMcu, blnMcu, thnMcu, SortMcu, opsi       int
}
type dataMcu [1000]Mcu

func main() {
	var tPasien dataPasien
	var tMcu dataMcu
	var nPasien, nMcu, x int
	var asci string
	asci = `
-----------------------------------------------------------------------------------------------------	


                                              __________
                                            /____________\
                _____                     /________________\
           ____/     \___                   |            |  
     _____/              \____              |     ||     |                            
    /_________________________\             |  ===||===  |                   _______                       
                                            |     ||     |                  /        \_____
                                        ____|____________|____            /_________________\                          
                                      /________________________\                                     
                                    /____________________________\                                        
                                      |   __   __    __   __   |                 ___   ___                 
         ___   ___             _______|  |  | |  |  |  | |  |  |_______             \_/                       
            \_/              /________|  |__| |__|  |__| |__|  |________\                            
              __  __        /_________|   __   __    __   __   |_________\                           
                \/           |   __   |  |  | |  |  |  | |  |  |   __   |     __  __                 
                             |  |  |  |  |__| |__|  |__| |__|  |  |  |  |       \/                     
        /\           /\      |  |__|  |   __   __    __   __   |  |__|  |                            
       /  \         /  \     |   __   |  |  | |  |  |  | |  |  |   __   |                 
      / || \       / || \    |  |  |  |  |__| |__|  |__| |__|  |  |  |  |                      
      / || \       / || \    |  |__|  |   __              __   |  |__|  |    _-__________-___
     /  ||  \     /  ||  \   |   __   |  |  |  ________  |  |  |   __   |   |          ___   | 
    /___||___\   /___||___\  |  |  |  |  |__| | = || = | |__|  |  |  |  |   |AMBULANCE| = |  \___      
        ||           ||      |  |__|  |       |   ||   |       |  |__|  |   |  __     |   |  __  \
        ||           ||      |        |       |   ||   |       |        |   |_(  )____|___|_(  )__|       
-----------------------------------------------------------------------------------------------------							                

	`
	for x != 13 {
		fmt.Println(asci)
		fmt.Println("-----------------------------------------------------------------------------------------------------")
		opsiMenu()
		fmt.Scan(&x)
		for !(x >= 1 && x <= 13) {
			fmt.Println("masukan opsi yang valid")
			fmt.Scan(&x)
		}
		switch x {
		case 1:
			tambahDataPasien(&tPasien, &nPasien)
		case 2:
			TambahMcu(&tMcu, &nMcu, &tPasien, nPasien)
		case 3:
			editDataPasien(&tPasien, nPasien, &tMcu, nMcu)
		case 4:
			HapusMcu(&tMcu, &nMcu, &tPasien, nPasien)
		case 5:
			editMcu(&tMcu, nMcu)
		case 6:
			menghapusDataPasien(&tPasien, &nPasien, &tMcu, &nMcu)
		case 7:
			melihatPasienYanMemilihPaket(tMcu, nMcu)
		case 8:
			melihatMcuTerutberdasarkanPriode(tMcu, nMcu)
		case 9:
			tampilkanPasienBerdasrakanNama(tPasien, nPasien)
		case 10:
			melihatMcuTerutberdasarkantanggalMcu(tMcu, nMcu)
		case 11:
			melihatMcuTerutberdasarkanPaketMcu(tMcu, nMcu)
		case 12:
			tampilkanDataPasienYangTerdaftar(tPasien, nPasien)
			var x string
			fmt.Println("ketik apapun untuk kembali")
			fmt.Scan(&x)
		case 13:
			fmt.Println(" - - - - - - - - - - - - TERIMA KASIH KARENA TELAH MENGGUNAKAN APLIKASI KAMI - - - - - - - - - - - - ")
			fmt.Println("                                                                                                     ")
			fmt.Println(" ___________________________________________________________________________________________________ ")
			fmt.Println("|                                                                                                   |")
			fmt.Println("|                                     KELOMPOK 6 ALPRO SE-46-03                                     |")
			fmt.Println("|                                                                                                   |")
			fmt.Println("|                                             ANGGOTA:                                              |")
			fmt.Println("|                                 KEVIN ALBANY JUNAIDI (1302223027)                                 |")
			fmt.Println("|                                                DAN                                                |")
			fmt.Println("|                                 HASAN NURRAHMAN PANE (1302220031)                                 |")
			fmt.Println("|                                                                                                   |")
			fmt.Println("|___________________________________________________________________________________________________|")
			fmt.Println(" ")
		}
	}

}

func tambahDataPasien(T *dataPasien, nPasien *int) {
	var idx int
	fmt.Println("Masukkan nama pasien")
	fmt.Println("Catatan: Gunakan underscore (_) untuk menambahkan spasi pada nama")
	fmt.Println("Contoh nama: Budi Ahmad ")
	fmt.Println("Contoh penulisan: Budi_Ahmad ")
	fmt.Scan(&T[*nPasien].nama)
	fmt.Println("Masukan id Pasien")
	fmt.Scan(&T[*nPasien].id)
	idx = CariIdPasien(*T, *nPasien, (T[*nPasien].id))
	if idx != -1 {
		fmt.Println("maaf id sudah terdaftar")
	} else {
		fmt.Println("masukan golongan darah pasien")
		fmt.Println("gunakan huruf besar untuk menuliskan golongan darah")
		fmt.Scan(&T[*nPasien].golonganDarah)
		for T[*nPasien].golonganDarah != "A" && T[*nPasien].golonganDarah != "B" && T[*nPasien].golonganDarah != "AB" && T[*nPasien].golonganDarah != "O" {
			fmt.Println("Masukan golongan dara yang valid")
			fmt.Println("gunakan huruf besar untuk menuliskan golongan darah")
			fmt.Scan(&T[*nPasien].golonganDarah)

		}
		fmt.Println(" ")
		fmt.Println("Masukkan tangal, bulan dan tahun Tahir")
		fmt.Println("Note: jangan gunakan 0 didepannya khususnya untuk 08 dan 09 contoh: 08 ")
		fmt.Println("Contoh penulisan 1 Agustus 2010")
		fmt.Println("Penulisan pada program akan menjadi seperti: 1 8 2010")
		fmt.Scan(&T[*nPasien].tglLahir, &T[*nPasien].blnLahir, &T[*nPasien].thnLahir)
		for !validasiTanggal(T[*nPasien].tglLahir, T[*nPasien].blnLahir, T[*nPasien].thnLahir) {
			fmt.Println("Masukkan tanggal, bulan, dan tahun yang valid")
			fmt.Scan(&T[*nPasien].tglLahir, &T[*nPasien].blnLahir, &T[*nPasien].thnLahir)
		}

		fmt.Println("                          SELAMAT ANDA TELAH BERHASIL MENAMBAHKAN DATA PASIEN BARU                          ")
		*nPasien++
	}

	fmt.Println("                                      takan tombol apapun untuk kembali                                  ")
	var x string
	fmt.Scan(&x)

}
func opsiMenu() {
	fmt.Println("+ + + + + + + + + + + + + + SELAMAT DATANG PADA APLIKASI MEDICAL CHECK UP + + + + + + + + + + + + + +")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("                                                                                                     ")
	fmt.Println("                                        1. Tambah data pasien                                        ")
	fmt.Println("                                         2. Tambah data Mcu                                          ")
	fmt.Println("                                         3. Edit data pasien                                         ")
	fmt.Println("                                           4. Hapus data mcu                                         ")
	fmt.Println("                                            5. Edit data Mcu                                         ")
	fmt.Println("                                         6. Hapus data pasien                                        ")
	fmt.Println("                              7. Tampilkan datamcu berdasarkan paket mcu                             ")
	fmt.Println("                       8. Tampilkan data pasien terurut berdasarkan priode                           ")
	fmt.Println("                              9. Tampilkan data pasien berdasarkan nama                              ")
	fmt.Println("                       10. Tampilkan data mcu terurut berdasarkan tanggal                            ")
	fmt.Println("                       11. Tampilkan data pasien terurut berdasarkan paket mcu                       ")
	fmt.Println("                       12. Tampilkan data pasien yang terdaftar di aplikasi mcu                      ")
	fmt.Println("                                                                                                     ")
	fmt.Println("                                              13. KELUAR                                             ")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("                             Pilihlah opsi yang valid diantara 1 hingga 13                           ")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
}
func TambahMcu(T *dataMcu, nMcu *int, Tpasien *dataPasien, nPasien int) {
	var idx int
	var jawaban string
	var pilihan int
	var namaPasien string
	var berenti bool
	var x string
	berenti = false
	if nPasien == 0 {
		fmt.Println("maaf data pasien masih kosong")
	} else {
		fmt.Println("apakah anda ingin melihat data pasien?")
		fmt.Println("ketik 1 untuk iya dan liannya untuk tidak")
		fmt.Scan(&x)
		if x == "1" {
			tampilkanDataPasienYangTerdaftar(*Tpasien, nPasien)
		}

		fmt.Println("Masukan nama pasien")
		fmt.Scan(&namaPasien)
		for i := 0; i < nPasien && berenti == false; i++ {
			if Tpasien[i].nama == namaPasien {
				fmt.Println("Nama Pasien: ", Tpasien[i].nama)
				fmt.Println("ID Pasien: ", Tpasien[i].id)
				fmt.Println("Golongan Darah: ", Tpasien[i].golonganDarah)
				fmt.Println("Tanggal Lahir: ", Tpasien[i].tglLahir, " ", Tpasien[i].blnLahir, " ", Tpasien[i].thnLahir)
				fmt.Println("Apakah pasien tersebut yang anda cari")
				fmt.Println("Ketik 1 untuk iya dan lainnya untuk mencari pasien lainnya")
				fmt.Scan(&jawaban)
				if jawaban == "1" {
					berenti = true
					idx = i
				}
			}

		}
		if berenti == false {
			idx = -1
		}

		if idx == -1 {
			fmt.Println("maaf nama pasien tidak ditemukan")
		} else {
			T[*nMcu].namaPasienMCU = Tpasien[idx].nama
			T[*nMcu].idPasien = Tpasien[idx].id
			fmt.Println("masukan tanggal mcu")
			fmt.Scan(&T[*nMcu].tglMcu, &T[*nMcu].blnMcu, &T[*nMcu].thnMcu)
			for !validasiTanggal(T[*nMcu].tglMcu, T[*nMcu].blnMcu, T[*nMcu].thnMcu) {
				fmt.Println("Masukkan tanggal, bulan, dan tahun yang valid")
				fmt.Scan(&T[*nMcu].tglMcu, &T[*nMcu].blnMcu, &T[*nMcu].thnMcu)
			}
			T[*nMcu].SortMcu = T[*nMcu].thnMcu*10000 + T[*nMcu].blnMcu*100 + T[*nMcu].tglMcu
			OpsipaketMcu()
			fmt.Println("masukan opsi mcu yang valid antara 1 -7")
			fmt.Scan(&T[*nMcu].opsi)
			for T[*nMcu].opsi <= 0 || T[*nMcu].opsi >= 8 {
				OpsipaketMcu()
				fmt.Println(" ! ! ! ! ! Mohon maaf paketMcu yang tersebut tidak tersedia, silahkan pilih paketMcu kembali ! ! ! ! ! ")
				fmt.Println(" ")
				fmt.Scan(&T[*nMcu].opsi)
			}
			if T[*nMcu].opsi == 1 {
				T[*nMcu].paketMcu = "Cek fungsi jantung"
			} else if T[*nMcu].opsi == 2 {
				T[*nMcu].paketMcu = "Pemeriksaan radiologi"
			} else if T[*nMcu].opsi == 3 {
				T[*nMcu].paketMcu = "Pemeriksaan laboratorium"
			} else if T[*nMcu].opsi == 4 {
				T[*nMcu].paketMcu = "Pemeriksaan kolesterol"
			} else if T[*nMcu].opsi == 5 {
				T[*nMcu].paketMcu = "Pemeriksaan gula darah"
			} else if T[*nMcu].opsi == 6 {
				T[*nMcu].paketMcu = "Pemeriksaan fungsi hati"
			} else if T[*nMcu].opsi == 7 {
				T[*nMcu].paketMcu = "Pemeriksaan fungsi ginjal"
			}

			for pilihan <= 0 || pilihan >= 4 {
				fmt.Println("                       Masukkan hasil rekap mcu yang valid diantara 1 sampai 3                       ")
				fmt.Println("                                               1. Sehat                                              ")
				fmt.Println("                                           2. Kurang Sehat                                           ")
				fmt.Println("                                              3. Kritis                                              ")
				fmt.Scan(&pilihan)
			}
			if pilihan == 1 {
				T[*nMcu].rekapMCU = "Sehat"
			} else if pilihan == 2 {
				T[*nMcu].rekapMCU = "Kurang sehat"
			} else if pilihan == 3 {
				T[*nMcu].rekapMCU = "Kritis"
			}
			*nMcu++
			Tpasien[idx].jumlahMCUPasien++
		}

	}

	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&pilihan)
}

func editDataPasien(T *dataPasien, nPasien int, Tmcu *dataMcu, nMcu int) {
	var nama, opsi string
	var idx int
	var berenti bool
	berenti = false
	if nPasien == 0 {
		fmt.Println("maaf data pasien masih kosong")
	} else {
		fmt.Println("apakah anda ingin menampilkan data pasien")
		fmt.Println("tekan 1 untuk iya dan lainnya untuk tidak")
		fmt.Scan(&opsi)
		if opsi == "1" {
			tampilkanDataPasienYangTerdaftar(*T, nPasien)
		}
		fmt.Println("masukan nama pasien yang ingin diedit")
		fmt.Scan(&nama)
		idx = -1
		for i := 0; i < nPasien && berenti == false; i++ {
			if T[i].nama == nama {
				fmt.Println("nama pasien: ", T[i].nama)
				fmt.Println("idPasien: ", T[i].id)
				fmt.Println("golongan darah: ", T[i].golonganDarah)
				fmt.Println("tahun lahir: ", T[i].tglLahir, " ", T[i].blnLahir, " ", T[i].thnLahir)
				fmt.Println("apakah data tersebut yang anda cari")
				fmt.Println("ketik 1 untuk iya dan cari lainnya untuk cari pasien lain")
				fmt.Scan(&opsi)
				if opsi == "1" {
					berenti = true
					idx = i
				}

			}

		}
		if idx == -1 {
			fmt.Println("maaf data yang anda cari tidak ada")
		} else {
			fmt.Println("apakah anda ingin nama pasie?")
			fmt.Println("ketik 1 untuk iya dan lainnya tidak")
			fmt.Scan(&opsi)
			if opsi == "1" {
				fmt.Scan(&T[idx].nama)
				for i := 0; i < nMcu; i++ {
					if T[idx].id == Tmcu[i].idPasien {
						Tmcu[i].namaPasienMCU = T[idx].nama
						
					}
				}
			}
			fmt.Println("apakah anda ingin mengedit golongan darah?")
			fmt.Println("ketik 1 untuk iya dan lainnya untuk tidak")
			fmt.Scan(&opsi)
			if opsi == "1" {
				fmt.Scan(&T[idx].golonganDarah)
				for T[nPasien].golonganDarah != "A" && T[idx].golonganDarah != "B" && T[idx].golonganDarah != "AB" && T[idx].golonganDarah != "O" {
					fmt.Println("Masukan golongan dara yang valid")
					fmt.Println("gunakan huruf besar untuk menuliskan golongan darah")
					fmt.Scan(&T[idx].golonganDarah)

				}
			}
			fmt.Println("apakah anda ingin mengedit tanggal lahir")
			fmt.Println("ketik 1 untuk iya dan lainnya untuk tidak")
			fmt.Scan(&opsi)
			if opsi == "1" {
				fmt.Scan(&T[idx].tglLahir, &T[idx].blnLahir, &T[idx].thnLahir)
				for !validasiTanggal(T[idx].tglLahir, T[idx].blnLahir, T[idx].thnLahir) {
					fmt.Println("Masukkan tanggal, bulan, dan tahun yang valid")
					fmt.Scan(&T[idx].tglLahir, &T[idx].blnLahir, &T[idx].thnLahir)
				}
			}
		}
	}

	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&opsi)

}
func CariIdPasien(T dataPasien, n int, id string) int {
	for i := 0; i < n; i++ {
		if T[i].id == id {
			return i
		}
	}
	return -1
}
func OpsipaketMcu() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("                                                                                                     ")
	fmt.Println("                             Jenis - Jenis paketMcu yang Tersedia di MCU :                            ")
	fmt.Println("                                                                                                     ")
	fmt.Println("                                        1. Cek fungsi jantung                                        ")
	fmt.Println("                                       2. Pemeriksaan radiologi                                      ")
	fmt.Println("                                     3. Pemeriksaan laboratorium                                     ")
	fmt.Println("                                      4. Pemeriksaan kolesterol                                      ")
	fmt.Println("                                      5. Pemeriksaan gula darah                                      ")
	fmt.Println("                                     6. Pemeriksaan fungsi hati                                      ")
	fmt.Println("                                     7. Pemeriksaan fungsi ginjal                                    ")
	fmt.Println("                                                                                                     ")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("+ + + + + + + + + + + + + +  Pilih jenis paketMcu MCU diantara 1 sampai 7  + + + + + + + + + + + + + +")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println(" ")
}

func cekKabisat(x int) bool {
	//function ini digunakan untuk mengecek apakah tahun yang dimasukan adalah kabisat atau bukan
	var kabisat bool
	kabisat = x%400 == 0 || x%4 == 0 && x%100 != 0
	return kabisat
}

func validasiTanggal(tgl, bln, thn int) bool {
	//function ini digunakan untuk mengecek apakah tanggal yang dimasukan sudah valid atau belum
	var valid bool
	var tanggal int
	valid = true
	if cekKabisat(thn) && bln == 2 {
		tanggal = 29
	} else if !cekKabisat(thn) && bln == 2 {
		tanggal = 28
	} else if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 9 || bln == 10 || bln == 12 {
		tanggal = 31
	} else {
		tanggal = 30
	}
	if bln <= 0 || bln > 12 || tgl <= 0 || tanggal-tgl < 0 {
		valid = false
	}
	return valid

}

func HapusMcu(Tmcu *dataMcu, nMcu *int, Tpasien *dataPasien, nPasien int) {
	var berenti bool
	var namaPasien string
	var jawaban string
	var idx, idx1 int
	var x string
	berenti = false
	if *nMcu == 0 {
		fmt.Println("maaf data mcu masih kosong")
	} else {
		fmt.Println("apakah anda ingin melihat data pasien?")
		fmt.Println("ketik 1 untuk iya dan lainnya untuk tidak")
		fmt.Scan(&x)
		if x == "1" {
			tampilkanDataPasienYangTerdaftar(*Tpasien, nPasien)
		}
		fmt.Println("Masukan nama pasien")
		fmt.Scan(&namaPasien)
		for i := 0; i < nPasien && berenti == false; i++ {
			if Tpasien[i].nama == namaPasien {
				fmt.Println("Nama Pasien: ", Tpasien[i].nama)
				fmt.Println("ID Pasien: ", Tpasien[i].id)
				fmt.Println("Golongan Darah: ", Tpasien[i].golonganDarah)
				fmt.Println("Tanggal Lahir: ", Tpasien[i].tglLahir, " ", Tpasien[i].blnLahir, " ", Tpasien[i].thnLahir)
				fmt.Println("Apakah pasien tersebut yang anda cari")
				fmt.Println("Ketik 1 untuk iya dan lainnya untuk mencari pasien lainnya")
				fmt.Scan(&jawaban)
				if jawaban == "1" {
					berenti = true
					idx1 = i
				}
			}

		}
		if berenti == false {
			idx1 = -1
		}

		if idx1 == -1 {
			fmt.Println("maaf id pasien tidak ditemukan")
		} else {
			berenti = false
			for i := idx; i < *nMcu && berenti == false; i++ {
				if Tpasien[idx1].id == Tmcu[i].idPasien {
					fmt.Println("paket Mcu: ", Tmcu[i].paketMcu)
					fmt.Println("hasil rekap Mcu: ", Tmcu[i].rekapMCU)
					fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].blnMcu, " ", Tmcu[i].tglMcu)
					fmt.Println("apakah data mcu ini yang ingin anda hapus")
					fmt.Println("ketik 1 jika dan lainnya untuk mencari yang lainnya")
					fmt.Scan(&jawaban)
					if jawaban == "1" {
						berenti = true
						idx = i
					}
				}
			}
			if *nMcu < 1000 {
				for i := idx; i < *nMcu; i++ {
					Tmcu[i] = Tmcu[i+1]
				}
			}
			*nMcu--
			Tpasien[idx1].jumlahMCUPasien--

		}

	}
	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&x)

}

func editMcu(Tmcu *dataMcu, nMcu int) {
	var namaPasien, jawaban string
	var berenti bool
	var idx, pilihan int
	var x string
	berenti = false
	if nMcu == 0 {
		fmt.Println("maaf data mcu masih kosog")
	} else {
		fmt.Println("apakah anda ingin melihat data mcu pasien?")
		fmt.Println("ketik 1 untuk iya dan lainnya untuk tidak")
		fmt.Scan(&x)
		if x == "1" {
			for i := 0; i < nMcu; i++ {
				fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
				fmt.Println("Id pasien: ", Tmcu[i].idPasien)
				fmt.Println("paket Mcu: ", Tmcu[i].paketMcu)
				fmt.Println("Rekap Mcu: ", Tmcu[i].rekapMCU)
				fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu)
			}
		}
		idx = -1
		fmt.Println("masukan nama pasien")
		fmt.Scan(&namaPasien)
		for i := 0; i < nMcu && berenti == false; i++ {
			if namaPasien == Tmcu[i].namaPasienMCU {
				fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
				fmt.Println("Id pasien: ", Tmcu[i].idPasien)
				fmt.Println("paket Mcu: ", Tmcu[i].paketMcu)
				fmt.Println("Rekap Mcu: ", Tmcu[i].rekapMCU)
				fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu)
				fmt.Println("apakah data tersebut yang anda cari")
				fmt.Println("untuk iya dan lainnya untuk tidak")
				fmt.Scan(&jawaban)
				if jawaban == "1" {
					idx = i
					berenti = true
				}
			}
		}
		if idx == -1 {
			fmt.Println("maaf data yang anda cari tidak ada")

		} else {
			fmt.Println("apakah anda ingin mengedit paket mcu")
			fmt.Println("ketik 1 untuk iya dan tidak untuk lainnya")
			fmt.Scan(&jawaban)
			if jawaban == "1" {
				OpsipaketMcu()
				fmt.Scan(&Tmcu[idx].opsi)
				for Tmcu[idx].opsi <= 0 || Tmcu[idx].opsi >= 8 {
					OpsipaketMcu()
					fmt.Println(" ! ! ! ! ! Mohon maaf paketMcu yang tersebut tidak tersedia, silahkan pilih paketMcu kembali ! ! ! ! ! ")
					fmt.Println(" ")
					fmt.Scan(&Tmcu[idx].opsi)
				}
				if Tmcu[idx].opsi == 1 {
					Tmcu[idx].paketMcu = "Cek fungsi jantung"
				} else if Tmcu[idx].opsi == 2 {
					Tmcu[idx].paketMcu = "Pemeriksaan radiologi"
				} else if Tmcu[idx].opsi == 3 {
					Tmcu[idx].paketMcu = "Pemeriksaan laboratorium"
				} else if Tmcu[idx].opsi == 4 {
					Tmcu[idx].paketMcu = "Pemeriksaan kolesterol"
				} else if Tmcu[idx].opsi == 5 {
					Tmcu[idx].paketMcu = "Pemeriksaan gula darah"
				} else if Tmcu[idx].opsi == 6 {
					Tmcu[idx].paketMcu = "Pemeriksaan fungsi hati"
				} else if Tmcu[idx].opsi == 7 {
					Tmcu[idx].paketMcu = "Pemeriksaan fungsi ginjal"
				}
			}
			fmt.Println("apakah anda ingin mengedit tanggal Mcu")
			fmt.Println("ketik 1 untuk iya dan tidak untuk lainnya")
			fmt.Scan(&jawaban)
			if jawaban == "1" {
				fmt.Println("masukan tanggal, bulan, dan tahun mcu")
				fmt.Scan(&Tmcu[idx].tglMcu, &Tmcu[idx].blnMcu, &Tmcu[idx].thnMcu)
				for !validasiTanggal(Tmcu[idx].tglMcu, Tmcu[idx].blnMcu, Tmcu[idx].thnMcu) {
					fmt.Println("Masukkan tanggal, bulan, dan tahun yang valid")
					fmt.Scan(&Tmcu[idx].tglMcu, &Tmcu[idx].blnMcu, &Tmcu[idx].thnMcu)
				}

			}
			fmt.Println("apakah anda ingin mengedit hasil rekap mcu")
			fmt.Println("ketik 1 untuk iya dan tidak untuk lainnya")
			fmt.Scan(&jawaban)
			if jawaban == "1" {
				for pilihan <= 0 || pilihan >= 4 {
					fmt.Println("                       Masukkan hasil rekap mcu yang valid diantara 1 sampai 3                       ")
					fmt.Println("                                               1. Sehat                                              ")
					fmt.Println("                                           2. Kurang Sehat                                           ")
					fmt.Println("                                              3. Kritis                                              ")
					fmt.Scan(&pilihan)
				}
				if pilihan == 1 {
					Tmcu[idx].rekapMCU = "Sehat"
				} else if pilihan == 2 {
					Tmcu[idx].rekapMCU = "Kurang sehat"
				} else if pilihan == 3 {
					Tmcu[idx].rekapMCU = "Kritis"
				}
			}
		}
	}
	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&jawaban)

}

func hapusPasien(tMcu *dataMcu, id string, nMcu *int) {
	var idx int
	var berenti bool
	berenti = false
	for i := 0; i < *nMcu && berenti == false; i++ {
		if tMcu[i].idPasien == id {
			idx = i
			berenti = true
		}

	}
	if *nMcu < 1000 {
		for idx < *nMcu {
			tMcu[idx] = tMcu[idx+1]
			idx++
		}
	}
	*nMcu--

}

func menghapusDataPasien(Tpasien *dataPasien, nPasien *int, Tmcu *dataMcu, nMcu *int) {
	var namaPasien string
	var berenti bool
	var idx int
	var jawaban string
	var x string

	if *nPasien == 0 {
		fmt.Println("maaf data pasien masih kosong")
	} else {
		fmt.Println("apakah anda ingin melihat data pasien?")
		fmt.Println("ketik 1 untuk iya dan liannya untuk tidak")
		fmt.Scan(&x)
		if x == "1" {
			tampilkanDataPasienYangTerdaftar(*Tpasien, *nPasien)
		}
		berenti = false
		idx = -1

		fmt.Println("masukan nama pasien yang ingin dihapus")
		fmt.Scan(&namaPasien)
		for i := 0; i < *nPasien && berenti == false; i++ {
			if namaPasien == Tpasien[i].nama {
				fmt.Println("nama Pasien: ", Tpasien[i].nama)
				fmt.Println("Id pasien: ", Tpasien[i].id)
				fmt.Println("Golongan Darah: ", Tpasien[i].golonganDarah)
				fmt.Println("Tanggal Lahir: ", Tpasien[i].tglLahir, " ", Tpasien[i].blnLahir, " ", Tpasien[i].thnLahir)
				fmt.Println("apakah data tersebut yang anda cari")
				fmt.Println("untuk iya dan lainnya untuk tidak")
				fmt.Scan(&jawaban)
				if jawaban == "1" {
					idx = i
					berenti = true
				}
			}
		}

		if idx == -1 {
			fmt.Println("maaf data pasien tidak ditemukam")
		} else {
			for i := 0; i < Tpasien[idx].jumlahMCUPasien; i++ {
				hapusPasien(Tmcu, Tpasien[idx].id, nMcu)
			}

			if *nPasien < 200 {
				for i := idx; i < *nPasien; i++ {
					Tpasien[i] = Tpasien[i+1]
				}
			}
			*nPasien--
		}
	}

	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&namaPasien)

}

func melihatPasienYanMemilihPaket(Tmcu dataMcu, nMcu int) {
	var paketMcu int
	var ada bool
	ada = false
	if nMcu == 0 {
		fmt.Println("data mcu masih kosong")
	} else {
		OpsipaketMcu()
		fmt.Scan(&paketMcu)
		for paketMcu <= 0 || paketMcu >= 8 {
			fmt.Println("masukan paket macu yang valid antara 1 -7")
			fmt.Scan(&paketMcu)
		}
		for i := 0; i < nMcu; i++ {
			if Tmcu[i].opsi == paketMcu {
				fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
				fmt.Println("Id pasien: ", Tmcu[i].idPasien)
				fmt.Println("Paket Mcu: ", Tmcu[i].paketMcu)
				fmt.Println("Hasil rekap Mcu: ", Tmcu[i].rekapMCU)
				fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].thnMcu)
				ada = true
			}
		}
		if ada == false {
			fmt.Println("maaf tidak ada pasien yang melakukan pada layan mcu tersebut")
		}
	}
	var x string
	fmt.Println("tekan apapun untuk kembali")
	fmt.Scan(&x)

}

func melihatMcuTerutberdasarkanPriode(Tmcu dataMcu, nMcu int) {
	var tgl1, bln1, thn1, tgl2, bln2, thn2, kodeTanggal1, kodeTanggal2 int
	var x string
	var ada bool
	ada = false

	if nMcu == 0 {
		fmt.Println("Maaf data yang miliki masih kosong")
	} else {
		fmt.Println("Masukkan tanggal pertama, bln pertama, dan tahun pertama")
		fmt.Scan(&tgl1, &bln1, &thn1)
		for !validasiTanggal(tgl1, bln1, thn1) {
			fmt.Println("Masukkan tanggal pertama, bln pertama, dan tahun pertama yang valid")
			fmt.Scan(&tgl1, &bln1, &thn1)
		}
		fmt.Println("Masukkan tanggal kedua, bln kedua, dan tahun kedua")
		fmt.Scan(&tgl2, &bln2, &thn2)
		for !validasiTanggal(tgl2, bln2, thn2) {
			fmt.Println("Masukkan tanggal kedua, bln kedua, dan tahun kedua yang valid")
			fmt.Scan(&tgl1, &bln1, &thn1)
		}
		kodeTanggal1 = thn1*10000 + bln1*100 + tgl1
		kodeTanggal2 = thn2*10000 + bln2*100 + tgl2
		for kodeTanggal1 >= kodeTanggal2 {
			fmt.Println("Masukkan tanggal pertama, bln pertama, dan tahun pertama yang valid")
			fmt.Scan(&tgl1, &bln1, &thn1)
			for !validasiTanggal(tgl1, bln1, thn1) {
				fmt.Println("Masukkan tanggal pertama, bln pertama, dan tahun pertama yang valid")
				fmt.Scan(&tgl1, &bln1, &thn1)
			}
			fmt.Println("Masukkan tanggal kedua, bln kedua, dan tahun kedua")
			fmt.Scan(&tgl2, &bln2, &thn2)
			for !validasiTanggal(tgl2, bln2, thn2) {
				fmt.Println("Masukkan tanggal kedua, bln kedua, dan tahun kedua yang valid")
				fmt.Scan(&tgl1, &bln1, &thn1)
			}
			kodeTanggal1 = thn1*10000 + bln1*100 + tgl1
			kodeTanggal2 = thn2*10000 + bln2*100 + tgl2
		}
		var i, j int

		var temp Mcu

		for i = 1; i < nMcu; i++ { // insertion sort
			j = i
			for j > 0 && Tmcu[j-1].SortMcu > Tmcu[j].SortMcu {
				temp = Tmcu[j]
				Tmcu[j] = Tmcu[j-1]
				Tmcu[j-1] = temp
				j--
			}
		}
		var low, high int
		low = 0
		high = nMcu - 1

		for low <= high && !ada { // binary search
			mid := (low + high) / 2

			if Tmcu[mid].SortMcu < kodeTanggal1 {
				low = mid + 1
			} else if Tmcu[mid].SortMcu > kodeTanggal2 {
				high = mid - 1
			} else {
				ada = true

			}
		}

		if ada == false {
			fmt.Println("Data pada priode yang ada cari tidak ada")
		} else {
			for i := 0; i < nMcu; i++ {
				if Tmcu[i].SortMcu >= kodeTanggal1 && Tmcu[i].SortMcu <= kodeTanggal2 {
					fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
					fmt.Println("Id pasien: ", Tmcu[i].idPasien)
					fmt.Println("Paket Mcu: ", Tmcu[i].paketMcu)
					fmt.Println("Hasil rekap Mcu: ", Tmcu[i].rekapMCU)
					fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].thnMcu)

				}
			}

		}
	}
	fmt.Println("Tekan apapun untuk kembali ke menu")
	fmt.Scan(&x)
}

func tampilkanPasienBerdasrakanNama(T dataPasien, nPasien int) {
	var nama, opsi string
	var idx int
	var berenti bool
	berenti = false
	if nPasien == 0 {
		fmt.Println("maaf data pasien masih kosong")
	} else {
		fmt.Println("apakah anda ingin menampilkan data pasien")
		fmt.Println("tekan 1 untuk iya dan lainnya untuk tidak")
		fmt.Scan(&opsi)
		if opsi == "1" {
			tampilkanDataPasienYangTerdaftar(T, nPasien)
		}
		fmt.Println("masukan nama pasien yang ingin ditampilkan")
		fmt.Scan(&nama)
		idx = -1
		for i := 0; i < nPasien && berenti == false; i++ {
			if T[i].nama == nama {
				fmt.Println("nama pasien: ", T[i].nama)
				fmt.Println("idPasien: ", T[i].id)
				fmt.Println("golongan darah: ", T[i].golonganDarah)
				fmt.Println("tahun lahir: ", T[i].tglLahir, " ", T[i].blnLahir, " ", T[i].thnLahir)
				fmt.Println("apakah data tersebut yang anda cari")
				fmt.Println("ketik 1 untuk iya dan cari lainnya untuk cari pasien lain")
				fmt.Scan(&opsi)
				if opsi == "1" {
					berenti = true
					idx = i
				}

			}

		}
		if idx == -1 {
			fmt.Println("maaf data yang anda cari tidak ada")
		} else {
			fmt.Println("nama pasien: ", T[idx].nama)
			fmt.Println("idPasien: ", T[idx].id)
			fmt.Println("golongan darah: ", T[idx].golonganDarah)
			fmt.Println("tahun lahir: ", T[idx].tglLahir, " ", T[idx].blnLahir, " ", T[idx].thnLahir)
		}
	}
	var x string
	fmt.Println("ketik apapun untuk kembali")
	fmt.Scan(&x)
}

func melihatMcuTerutberdasarkantanggalMcu(Tmcu dataMcu, nMcu int) {
	var i, j, idx_min int
	var temp Mcu

	i = 1
	for i <= nMcu-1 {
		idx_min = i - 1
		j = i
		for j < nMcu {
			if Tmcu[idx_min].SortMcu > Tmcu[j].SortMcu {
				idx_min = j
			}
			j++
		}
		temp = Tmcu[idx_min]
		Tmcu[idx_min] = Tmcu[i-1]
		Tmcu[i-1] = temp
		i++

	}
	if nMcu == 0 {
		fmt.Println("Data pasien masih kosong")
	}
	for i := 0; i < nMcu; i++ {
		fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
		fmt.Println("Id pasien: ", Tmcu[i].idPasien)
		fmt.Println("Paket Mcu: ", Tmcu[i].paketMcu)
		fmt.Println("Hasil rekap Mcu: ", Tmcu[i].rekapMCU)
		fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].thnMcu)
	}
	fmt.Println("tekan apapun untuk kembali")
	var x string
	fmt.Scan(&x)
}

func melihatMcuTerutberdasarkanPaketMcu(Tmcu dataMcu, nMcu int) {
	var i, j, idx_min int
	var x string
	var temp Mcu

	if nMcu == 0 {
		fmt.Println("Maaf data pasien masih kosong")
	} else {
		i = 1
		for i <= nMcu-1 {
			idx_min = i - 1
			j = i
			for j < nMcu {
				if Tmcu[idx_min].opsi > Tmcu[j].opsi {
					idx_min = j
				}
				j++
			}
			temp = Tmcu[idx_min]
			Tmcu[idx_min] = Tmcu[i-1]
			Tmcu[i-1] = temp
			i++

		}

		for i := 0; i < nMcu; i++ {
			fmt.Println("nama Pasien: ", Tmcu[i].namaPasienMCU)
			fmt.Println("Id pasien: ", Tmcu[i].idPasien)
			fmt.Println("Paket Mcu: ", Tmcu[i].paketMcu)
			fmt.Println("Hasil rekap Mcu: ", Tmcu[i].rekapMCU)
			fmt.Println("Tanggal Mcu: ", Tmcu[i].tglMcu, " ", Tmcu[i].tglMcu, " ", Tmcu[i].thnMcu)
		}
	}
	fmt.Println("Tekan apapun untuk kembali ke menu")
	fmt.Scan(&x)
}

func tampilkanDataPasienYangTerdaftar(Tpasien dataPasien, nPasien int) {
	if nPasien == 0 {
		fmt.Println("maaf data pasien masih kosong")
	} else {
		for i := 0; i < nPasien; i++ {
			fmt.Println("nama pasien: ", Tpasien[i].nama)
			fmt.Println("idPasien: ", Tpasien[i].id)
			fmt.Println("golongan darah: ", Tpasien[i].golonganDarah)
			fmt.Println("tahun lahir: ", Tpasien[i].tglLahir, " ", Tpasien[i].blnLahir, " ", Tpasien[i].thnLahir)
		}

	}

}
