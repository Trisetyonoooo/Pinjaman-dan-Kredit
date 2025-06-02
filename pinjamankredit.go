package main

import (
	"fmt"
	"sort"
	"strings"
)

type Peminjam struct {
	Nama            string
	JumlahPinjaman  float64
	Tenor           int
	Bunga           float64
	Status          string
	TotalPembayaran float64
	SisaPembayaran  float64
}

var dataPeminjam []Peminjam
var pilihan int

func tambahPeminjam() {
	var p Peminjam
	var totalKewajiban float64

	fmt.Println("\n========== Tambah Data Peminjam ==========")
	fmt.Print("Masukkan nama peminjam       : ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Masukkan jumlah pinjaman     : ")
	fmt.Scanln(&p.JumlahPinjaman)
	fmt.Print("Masukkan tenor (bulan)       : ")
	fmt.Scanln(&p.Tenor)
	fmt.Print("Masukkan bunga pinjaman (%)  : ")
	fmt.Scanln(&p.Bunga)

	p.Status = "Belum Lunas"
	totalKewajiban = p.JumlahPinjaman + (p.JumlahPinjaman * p.Bunga / 100)
	p.SisaPembayaran = totalKewajiban
	p.TotalPembayaran = 0

	dataPeminjam = append(dataPeminjam, p)
	fmt.Println("\n✅ Data peminjam berhasil ditambahkan!")
}

func ubahPeminjam() {
	var nama string
	fmt.Println("\n========== Ubah Data Peminjam ==========")
	fmt.Print("Masukkan nama peminjam yang ingin diubah: ")
	fmt.Scanln(&nama)

	var i int
	for i = 0; i < len(dataPeminjam); i++ {
		if strings.EqualFold(dataPeminjam[i].Nama, nama) {
			fmt.Print("Masukkan jumlah pinjaman baru: ")
			fmt.Scanln(&dataPeminjam[i].JumlahPinjaman)
			fmt.Print("Masukkan tenor baru (bulan)  : ")
			fmt.Scanln(&dataPeminjam[i].Tenor)
			fmt.Print("Masukkan bunga baru (%)      : ")
			fmt.Scanln(&dataPeminjam[i].Bunga)

			var totalKewajiban float64 = dataPeminjam[i].JumlahPinjaman + (dataPeminjam[i].JumlahPinjaman * dataPeminjam[i].Bunga / 100)
			dataPeminjam[i].SisaPembayaran = totalKewajiban - dataPeminjam[i].TotalPembayaran

			if dataPeminjam[i].SisaPembayaran <= 0 {
				dataPeminjam[i].Status = "Lunas"
				dataPeminjam[i].SisaPembayaran = 0
			} else {
				dataPeminjam[i].Status = "Belum Lunas"
			}
			fmt.Println("\n✅ Data berhasil diperbarui!")
			return
		}
	}
	fmt.Println("\n❌ Data tidak ditemukan.")
}

func hapusPeminjam() {
	var nama string
	var i int

	fmt.Println("\n========== Hapus Data Peminjam ==========")
	fmt.Print("Masukkan nama peminjam yang ingin dihapus: ")
	fmt.Scanln(&nama)

	for i = 0; i < len(dataPeminjam); i++ {
		if strings.EqualFold(dataPeminjam[i].Nama, nama) {
			dataPeminjam = append(dataPeminjam[:i], dataPeminjam[i+1:]...)
			fmt.Println("\n✅ Data berhasil dihapus!")
			return
		}
	}
	fmt.Println("\n❌ Data tidak ditemukan.")
}

func hitungCicilan(p Peminjam) float64 {
	var total float64 
	var cicilan float64 
    total = p.JumlahPinjaman + (p.JumlahPinjaman * p.Bunga / 100)
    cicilan = total / float64(p.Tenor)
	return cicilan
}

func cariSequential(nama string) {
	var i int
	fmt.Println("\n========== Hasil Pencarian ==========")
	for i = 0; i < len(dataPeminjam); i++ {
		if strings.EqualFold(dataPeminjam[i].Nama, nama) {
			var p Peminjam = dataPeminjam[i]
			fmt.Println("✅ Data ditemukan!")
			fmt.Println("--------------------------------------")
			fmt.Println("Nama Peminjam     :", p.Nama)
			fmt.Println("Jumlah Pinjaman   :", formatUang(p.JumlahPinjaman))
			fmt.Println("Tenor (bulan)     :", p.Tenor)
			fmt.Println("Bunga (%)         :", p.Bunga)
			fmt.Println("Status Pembayaran :", p.Status)
			fmt.Println("Total Dibayar     :", formatUang(p.TotalPembayaran))
			fmt.Println("Sisa Pembayaran   :", formatUang(p.SisaPembayaran))
			fmt.Println("--------------------------------------")
			return
		}
	}
	fmt.Println("❌ Data tidak ditemukan.")
}

func cariBinary(nama string) {
	var low, high, mid int

	sort.Slice(dataPeminjam, func(a int, b int) bool {
		return strings.ToLower(dataPeminjam[a].Nama) < strings.ToLower(dataPeminjam[b].Nama)
	})

	low = 0
	high = len(dataPeminjam) - 1

	for low <= high {
		mid = (low + high) / 2
		if strings.EqualFold(dataPeminjam[mid].Nama, nama) {
			cariSequential(dataPeminjam[mid].Nama)
			return
		} else if strings.ToLower(dataPeminjam[mid].Nama) < strings.ToLower(nama) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("❌ Data tidak ditemukan.")
}

func selectionSortJumlah() {
	var i, j, minIdx int
	var temp Peminjam

	for i = 0; i < len(dataPeminjam)-1; i++ {
		minIdx = i
		for j = i + 1; j < len(dataPeminjam); j++ {
			if dataPeminjam[j].JumlahPinjaman < dataPeminjam[minIdx].JumlahPinjaman {
				minIdx = j
			}
		}
		temp = dataPeminjam[i]
		dataPeminjam[i] = dataPeminjam[minIdx]
		dataPeminjam[minIdx] = temp
	}
	fmt.Println("\n✅ Data berhasil diurutkan berdasarkan jumlah pinjaman!")
}

func insertionSortTenor() {
	var i, j int
	var key Peminjam

	for i = 1; i < len(dataPeminjam); i++ {
		key = dataPeminjam[i]
		j = i - 1
		for j >= 0 && dataPeminjam[j].Tenor > key.Tenor {
			dataPeminjam[j+1] = dataPeminjam[j]
			j--
		}
		dataPeminjam[j+1] = key
	}
	fmt.Println("\n✅ Data berhasil diurutkan berdasarkan tenor!")
}

func bayarCicilan() {
	var nama string
	var pembayaran float64
	var i int

	fmt.Println("\n========== Pembayaran Cicilan ==========")
	fmt.Print("Masukkan nama peminjam: ")
	fmt.Scanln(&nama)

	for i = 0; i < len(dataPeminjam); i++ {
		if strings.EqualFold(dataPeminjam[i].Nama, nama) {
			fmt.Println("Sisa pembayaran:", formatUang(dataPeminjam[i].SisaPembayaran))
			fmt.Print("Masukkan jumlah pembayaran: ")
			fmt.Scanln(&pembayaran)
			if pembayaran <= 0 {
				fmt.Println("❌ Jumlah tidak valid.")
				return
			}
			dataPeminjam[i].TotalPembayaran += pembayaran
			dataPeminjam[i].SisaPembayaran -= pembayaran
			if dataPeminjam[i].SisaPembayaran <= 0 {
				dataPeminjam[i].Status = "Lunas"
				dataPeminjam[i].SisaPembayaran = 0
				fmt.Println("\nCicilan lunas! Status diperbarui menjadi 'Lunas'.")
			} else {
				fmt.Println("✅ Pembayaran berhasil. Sisa pembayaran:", formatUang(dataPeminjam[i].SisaPembayaran))
			}
			return
		}
	}
	fmt.Println("❌ Data tidak ditemukan.")
}

func tampilkanLaporan() {
	var i, spasiKiri, totalLebarTabel int
	var judul, spasi string 

	totalLebarTabel = 135
	judul = "LAPORAN DATA PEMINJAM"
	spasiKiri = (totalLebarTabel - len(judul)) / 2
	spasi = ""

	for i = 0; i < spasiKiri; i++ {
		spasi += " "
	}

	fmt.Println()
	for i = 0; i < totalLebarTabel; i++ {
		fmt.Print("=")
	}
	fmt.Println()
	fmt.Println(spasi + judul)
	for i = 0; i < totalLebarTabel; i++ {
		fmt.Print("=")
	}
	fmt.Println()

	if len(dataPeminjam) == 0 {
		fmt.Println("Belum ada data peminjam.")
		return
	}


	fmt.Printf("| %-20s | %-20s | %-6s | %-20s | %-12s | %-20s | %-20s |\n",
		"Nama", "Jumlah", "Tenor", "Cicilan/Bulan", "Status", "Total Dibayar", "Sisa Pembayaran")

	for i = 0; i < totalLebarTabel; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i = 0; i < len(dataPeminjam); i++ {
		var p Peminjam = dataPeminjam[i]
		var cicilan float64 = hitungCicilan(p)

		fmt.Printf("| %-20s | %20s | %5dbl | %20s | %-12s | %20s | %20s |\n",
			p.Nama,
			formatUang(p.JumlahPinjaman),
			p.Tenor,
			formatUang(cicilan),
			p.Status,
			formatUang(p.TotalPembayaran),
			formatUang(p.SisaPembayaran))
	}

	for i = 0; i < totalLebarTabel; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}


func formatUang(x float64) string {
	var bil int64
	var sisa int64
	var strAngka string
	var i int
	var digit int64
	var charDigit string

	bil = int64(x)
	sisa = int64((x - float64(bil)) * 100)
	if sisa < 0 {
		sisa = 0
	}

	strAngka = ""
	i = 1

	for bil > 0 {
		digit = bil % 10
		charDigit = string(digit + 48)
		strAngka = charDigit + strAngka
		if i%3 == 0 && bil/10 > 0 {
			strAngka = "." + strAngka
		}
		i++
		bil = bil / 10
	}

	if strAngka == "" {
		strAngka = "0"
	}

	var formatted string
	formatted = fmt.Sprintf("Rp. %s,%02d", strAngka, sisa)
	return fmt.Sprintf("%20s", formatted)
}



func tampilkanMenu() {
	fmt.Println("\n========================================")
	fmt.Println("        APLIKASI SIMULASI PINJAMAN")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Data Peminjam")
	fmt.Println("2. Ubah Data Peminjam")
	fmt.Println("3. Hapus Data Peminjam")
	fmt.Println("4. Cari Nama (Sequential Search)")
	fmt.Println("5. Cari Nama (Binary Search)")
	fmt.Println("6. Urutkan Berdasarkan Jumlah Pinjaman")
	fmt.Println("7. Urutkan Berdasarkan Tenor")
	fmt.Println("8. Bayar Cicilan")
	fmt.Println("9. Tampilkan Laporan")
	fmt.Println("10. Keluar")
	fmt.Println("========================================")
	fmt.Print("Pilih menu [1-10]: ")
}

func main() {
	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			tambahPeminjam()
		} else if pilihan == 2 {
			ubahPeminjam()
		} else if pilihan == 3 {
			hapusPeminjam()
		} else if pilihan == 4 {
			var nama string
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&nama)
			cariSequential(nama)
		} else if pilihan == 5 {
			var nama string
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&nama)
			cariBinary(nama)
		} else if pilihan == 6 {
			selectionSortJumlah()
		} else if pilihan == 7 {
			insertionSortTenor()
		} else if pilihan == 8 {
			bayarCicilan()
		} else if pilihan == 9 {
			tampilkanLaporan()
		} else if pilihan == 10 {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}