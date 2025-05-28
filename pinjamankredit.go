package main

import (
    "fmt"
    "strings"
    "sort"
)

type Peminjam struct {
    Nama             string
    JumlahPinjaman   float64
    Tenor            int
    Bunga            float64
    Status           string
    TotalPembayaran  float64
    SisaPembayaran   float64
}

var dataPeminjam []Peminjam

func tambahPeminjam() {
    var p Peminjam
    fmt.Print("Masukkan nama peminjam: ")
    fmt.Scanln(&p.Nama)
    fmt.Print("Masukkan jumlah pinjaman: ")
    fmt.Scanln(&p.JumlahPinjaman)
    fmt.Print("Masukkan tenor (bulan): ")
    fmt.Scanln(&p.Tenor)
    fmt.Print("Masukkan bunga pinjaman (%): ")
    fmt.Scanln(&p.Bunga)

    p.Status = "Belum Lunas"
    totalKewajiban := p.JumlahPinjaman + (p.JumlahPinjaman * p.Bunga / 100)
    p.SisaPembayaran = totalKewajiban
    p.TotalPembayaran = 0

    dataPeminjam = append(dataPeminjam, p)
    fmt.Println("Data berhasil ditambahkan.")
}

func ubahPeminjam() {
    var nama string
    fmt.Print("Masukkan nama peminjam yang ingin diubah: ")
    fmt.Scanln(&nama)
    for i := range dataPeminjam {
        if strings.EqualFold(dataPeminjam[i].Nama, nama) {
            fmt.Print("Masukkan jumlah pinjaman baru: ")
            fmt.Scanln(&dataPeminjam[i].JumlahPinjaman)
            fmt.Print("Masukkan tenor baru (bulan): ")
            fmt.Scanln(&dataPeminjam[i].Tenor)
            fmt.Print("Masukkan bunga baru (%): ")
            fmt.Scanln(&dataPeminjam[i].Bunga)

            totalKewajiban := dataPeminjam[i].JumlahPinjaman + (dataPeminjam[i].JumlahPinjaman * dataPeminjam[i].Bunga / 100)
            dataPeminjam[i].SisaPembayaran = totalKewajiban - dataPeminjam[i].TotalPembayaran
            if dataPeminjam[i].SisaPembayaran <= 0 {
                dataPeminjam[i].Status = "Lunas"
                dataPeminjam[i].SisaPembayaran = 0
            } else {
                dataPeminjam[i].Status = "Belum Lunas"
            }

            fmt.Println("Data berhasil diubah.")
            return
        }
    }
    fmt.Println("Data tidak ditemukan.")
}

func hapusPeminjam() {
    var nama string
    fmt.Print("Masukkan nama peminjam yang ingin dihapus: ")
    fmt.Scanln(&nama)
    for i := range dataPeminjam {
        if strings.EqualFold(dataPeminjam[i].Nama, nama) {
            dataPeminjam = append(dataPeminjam[:i], dataPeminjam[i+1:]...)
            fmt.Println("Data berhasil dihapus.")
            return
        }
    }
    fmt.Println("Data tidak ditemukan.")
}

func hitungCicilan(p Peminjam) float64 {
    total := p.JumlahPinjaman + (p.JumlahPinjaman * p.Bunga / 100)
    return total / float64(p.Tenor)
}

func cariSequential(nama string) {
    for i := 0; i < len(dataPeminjam); i++ {
        if strings.EqualFold(dataPeminjam[i].Nama, nama) {
            p := dataPeminjam[i]
            fmt.Println("Data ditemukan:")
            fmt.Println("Nama Peminjam     :", p.Nama)
            fmt.Printf("Jumlah Pinjaman   : Rp %.2f\n", p.JumlahPinjaman)
            fmt.Println("Tenor (bulan)     :", p.Tenor)
            fmt.Println("Bunga (%)         :", p.Bunga)
            fmt.Println("Status Pembayaran :", p.Status)
            fmt.Printf("Total Dibayar     : Rp %.2f\n", p.TotalPembayaran)
            fmt.Printf("Sisa Pembayaran   : Rp %.2f\n", p.SisaPembayaran)
            return
        }
    }
    fmt.Println("Data tidak ditemukan.")
}

func cariBinary(nama string) {
    sort.Slice(dataPeminjam, func(i, j int) bool {
        return strings.ToLower(dataPeminjam[i].Nama) < strings.ToLower(dataPeminjam[j].Nama)
    })
    low, high := 0, len(dataPeminjam)-1
    for low <= high {
        mid := (low + high) / 2
        if strings.EqualFold(dataPeminjam[mid].Nama, nama) {
            fmt.Println("Data ditemukan:", dataPeminjam[mid])
            return
        } else if strings.ToLower(dataPeminjam[mid].Nama) < strings.ToLower(nama) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    fmt.Println("Data tidak ditemukan.")
}

func selectionSortJumlah() {
    n := len(dataPeminjam)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if dataPeminjam[j].JumlahPinjaman < dataPeminjam[minIdx].JumlahPinjaman {
                minIdx = j
            }
        }
        dataPeminjam[i], dataPeminjam[minIdx] = dataPeminjam[minIdx], dataPeminjam[i]
    }
    fmt.Println("Data berhasil diurutkan berdasarkan jumlah pinjaman.")
}

func insertionSortTenor() {
    for i := 1; i < len(dataPeminjam); i++ {
        key := dataPeminjam[i]
        j := i - 1
        for j >= 0 && dataPeminjam[j].Tenor > key.Tenor {
            dataPeminjam[j+1] = dataPeminjam[j]
            j--
        }
        dataPeminjam[j+1] = key
    }
    fmt.Println("Data berhasil diurutkan berdasarkan tenor.")
}

func bayarCicilan() {
    var nama string
    var jumlah float64
    fmt.Print("Masukkan nama peminjam: ")
    fmt.Scanln(&nama)
    for i := range dataPeminjam {
        if strings.EqualFold(dataPeminjam[i].Nama, nama) {
            fmt.Printf("Sisa yang harus dibayar: Rp %.2f\n", dataPeminjam[i].SisaPembayaran)
            fmt.Print("Masukkan jumlah pembayaran: ")
            fmt.Scanln(&jumlah)
            if jumlah <= 0 {
                fmt.Println("Jumlah pembayaran tidak valid.")
                return
            }
            dataPeminjam[i].TotalPembayaran += jumlah
            dataPeminjam[i].SisaPembayaran -= jumlah

            if dataPeminjam[i].SisaPembayaran <= 0 {
                dataPeminjam[i].Status = "Lunas"
                dataPeminjam[i].SisaPembayaran = 0
                fmt.Println("Pembayaran lunas! Status diperbarui menjadi 'Lunas'.")
            } else {
                fmt.Printf("Pembayaran berhasil. Sisa pembayaran: Rp %.2f\n", dataPeminjam[i].SisaPembayaran)
            }
            return
        }
    }
    fmt.Println("Data peminjam tidak ditemukan.")
}

func tampilkanLaporan() {
    fmt.Println("\n===================== LAPORAN DATA PEMINJAM =====================")
    if len(dataPeminjam) == 0 {
        fmt.Println("Belum ada data peminjam.")
        return
    }

    fmt.Println("+----------------------+---------------+--------+------------------+--------------+-----------------+-----------------+")
    fmt.Printf("| %-20s | %-13s | %-6s | %-16s | %-12s | %-15s | %-15s |\n",
        "Nama",
        "Jumlah",
        "Tenor",
        "Cicilan/Bulan",
        "Status",
        "Total Dibayar",
        "Sisa Pembayaran")
    fmt.Println("+----------------------+---------------+--------+------------------+--------------+-----------------+-----------------+")

    for i := 0; i < len(dataPeminjam); i++ {
        p := dataPeminjam[i]
        cicilan := hitungCicilan(p)
        fmt.Printf("| %-20s | Rp%-10.2f | %4dbl | Rp%-13.2f | %-12s | Rp%-12.2f  | Rp%-12.2f  |\n",
            p.Nama,
            p.JumlahPinjaman,
            p.Tenor,
            cicilan,
            p.Status,
            p.TotalPembayaran,
            p.SisaPembayaran)
    }
    fmt.Println("+----------------------+---------------+--------+------------------+--------------+-----------------+-----------------+")
}

func main() {
    var pilihan int
    for {
        fmt.Println("\n--- Menu Aplikasi Simulasi Pinjaman ---")
        fmt.Println("1. Tambah Data Peminjam")
        fmt.Println("2. Ubah Data Peminjam")
        fmt.Println("3. Hapus Data Peminjam")
        fmt.Println("4. Cari Nama (Sequential Search)")
        fmt.Println("5. Cari Nama (Binary Search)")
        fmt.Println("6. Urutkan Berdasarkan Jumlah Pinjaman (Selection Sort)")
        fmt.Println("7. Urutkan Berdasarkan Tenor (Insertion Sort)")
        fmt.Println("8. Bayar Cicilan")
        fmt.Println("9. Tampilkan Laporan")
        fmt.Println("10. Keluar")
        fmt.Print("Pilih menu [1-10]: ")
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            tambahPeminjam()
        case 2:
            ubahPeminjam()
        case 3:
            hapusPeminjam()
        case 4:
            var nama string
            fmt.Print("Masukkan nama: ")
            fmt.Scanln(&nama)
            cariSequential(nama)
        case 5:
            var nama string
            fmt.Print("Masukkan nama: ")
            fmt.Scanln(&nama)
            cariBinary(nama)
        case 6:
            selectionSortJumlah()
        case 7:
            insertionSortTenor()
        case 8:
            bayarCicilan()
        case 9:
            tampilkanLaporan()
        case 10:
            fmt.Println("Terima kasih telah menggunakan aplikasi!")
            return
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}