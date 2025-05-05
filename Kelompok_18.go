package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Isi_Ide struct {
	Judul     string
	Deskripsi string
	Kategori  string
	Rating    int
	Tanggal   time.Time
}

var Daftar_Ide []Isi_Ide

func main() {

	var pilih int

	for {
		fmt.Println("   APLIKASI PENGELOLAAN IDE    ")
		fmt.Println("1.Tampil Daftar Ide")
		fmt.Println("2.Tambahkan Ide")
		fmt.Println("3.Mengedit Ide")
		fmt.Println("4.Hapus Ide")
		fmt.Println("5.Kasih Rating Ide")
		fmt.Println("6.Cari Ide")
		fmt.Println("7.Urutan Ide")
		fmt.Println("8.Keluar")

		fmt.Print("Pilih Menu (1-8) : ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			Lihat_Daftar_Ide()
		} else if pilih == 2 {
			Tambah_Ide_Proyek()
		} else if pilih == 3 {
			Mengubah_Ide_Proyek()
		} else if pilih == 4 {
			Hapus_Ide_Proyek()
		} else if pilih == 5 {
			Kasih_Rating_ide_Proyek()
		} else if pilih == 6 {
			Cari_Ide()
		} else if pilih == 7 {
			Urutan_ide()
		} else if pilih == 8 {
			fmt.Println("Terima Kasih")
			return
		} else {
			fmt.Print("Menu Enggak Ada")
		}
	}
}

func Lihat_Daftar_Ide() {

	if len(Daftar_Ide) == 0 {
		fmt.Printf("\nBelum Ada Daftar Ide\n\n")
		return
	}

	fmt.Println("Daftar Ide")

	for i := 0; i < len(Daftar_Ide); i++ {

		Ide := Daftar_Ide[i]

		fmt.Printf("\n%d. %s\n", i+1, Ide.Judul)

		fmt.Println("Deskripsi Ide\t: ", Ide.Deskripsi)
		fmt.Println("Kategori\t: ", Ide.Kategori)
		fmt.Println("Rating\t\t: ", Ide.Rating)
		fmt.Println("Tanggal\t\t: ", Ide.Tanggal)

		fmt.Println()
	}

}

func Tambah_Ide_Proyek() {

	Input := bufio.NewReader(os.Stdin)

	var Ide Isi_Ide

	fmt.Print("Masukkan Judul Ide : ")
	Judul, _ := Input.ReadString('\n')
	Ide.Judul = strings.TrimSpace(Judul)

	fmt.Print("Masukkan Deskripsi Ide : ")
	Deskripsi, _ := Input.ReadString('\n')
	Ide.Deskripsi = strings.TrimSpace(Deskripsi)

	fmt.Print("Masukkan Kategori : ")
	Kategori, _ := Input.ReadString('\n')
	Ide.Kategori = strings.TrimSpace(Kategori)

	Ide.Rating = 0

	Ide.Tanggal = time.Now()

	Daftar_Ide = append(Daftar_Ide, Ide)
	fmt.Printf("Ide Berhasil Ditambahkan\n\n")

}

func Mengubah_Ide_Proyek() {

	if len(Daftar_Ide) == 0 {
		fmt.Println("Belum Ada Daftar Ide")
		return
	}

	Lihat_Daftar_Ide()

	var Nomor_Ide int
	fmt.Print("Masukkan Nomor Ide Yang Ingin Di Ubah : ")
	fmt.Scanln(&Nomor_Ide)

	if Nomor_Ide < 1 || Nomor_Ide > len(Daftar_Ide) {
		fmt.Println("Nomor Tidak Ada Di Daftar Ide")
		return
	}

	Input := bufio.NewReader(os.Stdin)
	Ide := &Daftar_Ide[Nomor_Ide-1]

	fmt.Printf("\nUBAH DATA IDE\n\n")
	fmt.Println("Judul Sekarang : ", Ide.Judul)
	fmt.Print("Judul Yang Baru  (Kosongkan Jika Tidak Ingin Mengubah Judul) : ")
	Judul, _ := Input.ReadString('\n')
	Judul = strings.TrimSpace(Judul)
	if Judul != "" {
		Ide.Judul = Judul
	}

	fmt.Println()

	fmt.Println("Deskripsi Sekarang : ", Ide.Deskripsi)
	fmt.Print("Deskripsi Yang baru (Kosongkan Jika Tidak Ingin Mengubah Deskripsi) : ")
	Deskripsi, _ := Input.ReadString('\n')
	Deskripsi = strings.TrimSpace(Deskripsi)
	if Deskripsi != "" {
		Ide.Deskripsi = Deskripsi
	}

	fmt.Println()

	fmt.Println("Kategori Sekarang : ", Ide.Kategori)
	fmt.Print("Deskripsi Yang baru (Kosongkan Jika Tidak Ingin Mengubah Deskripsi) : ")
	Kategori, _ := Input.ReadString('\n')
	Kategori = strings.TrimSpace(Kategori)
	if Kategori != "" {
		Ide.Kategori = Kategori
	}

	fmt.Printf("\nIde Berhasil Diperbarui\n\n")

}

func Hapus_Ide_Proyek() {

	if len(Daftar_Ide) == 0 {
		fmt.Println("Belum Ada Daftar Ide")
		return
	}

	Lihat_Daftar_Ide()

	var Nomor_Ide int
	fmt.Print("Masukkan Nomor Ide Yang Ingin Dihapus : ")
	fmt.Scanln(&Nomor_Ide)

	if Nomor_Ide < 1 || Nomor_Ide > len(Daftar_Ide) {
		fmt.Println("Nomor Tidak Ada")
		return
	}

	Ide_Yang_Dihapus := Daftar_Ide[Nomor_Ide-1].Judul
	Daftar_Ide = append(Daftar_Ide[:Nomor_Ide-1], Daftar_Ide[Nomor_Ide:]...)
	fmt.Println("Judul ", Ide_Yang_Dihapus, " Telah Dihapus")
	fmt.Println()

}

func Kasih_Rating_ide_Proyek() {

	if len(Daftar_Ide) == 0 {
		fmt.Println("Belum Ada Ide Proyek")
		return
	}

	Lihat_Daftar_Ide()

	var Pilih_Nomor_Ide, Masukkan_Rating int
	fmt.Print("Masukkan Nomor Ide Yang Ingin Di kasih Rating : ")
	fmt.Scanln(&Pilih_Nomor_Ide)

	if Pilih_Nomor_Ide < 1 || Pilih_Nomor_Ide > len(Daftar_Ide) {
		fmt.Println("Nomor Tidak Ada Pada Ide")
		return
	}

	fmt.Print("Masukkan Rating (1-10): ")
	fmt.Scanln(&Masukkan_Rating)

	if Masukkan_Rating < 0 || Masukkan_Rating > 10 {
		fmt.Println("Hanya Bisa Ngasih 0 - 10 ")
		return
	}

	Daftar_Ide[Pilih_Nomor_Ide-1].Rating = Masukkan_Rating
	fmt.Printf("Rating %d/10 berhasil dimasukkan untuk ide '%s'\n\n", Masukkan_Rating, Daftar_Ide[Pilih_Nomor_Ide-1].Judul)

}

func Binary_Search(cariJudul string) int {

	Selection_Sort() // fungsi Sorting berdasarkan Judul

	rendah := 0
	tinggi := len(Daftar_Ide) - 1

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2
		cari := strings.ToLower(cariJudul)
		target := strings.ToLower(Daftar_Ide[tengah].Judul)

		if cari == target {
			return tengah
		} else if target < cari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	return -1 
}

func Squential_Search(cariJudul string) int{
	
	for i := 0; i < len(Daftar_Ide); i++ {
        if strings.Contains(strings.ToLower(Daftar_Ide[i].Judul), strings.ToLower(cariJudul)) {
            return i
        }
    }
    return -1
}

func Cari_Ide() {

	if len(Daftar_Ide) == 0 {
		fmt.Print("Belum Ada Daftar Ide")
		return
	}
	input := bufio.NewReader(os.Stdin)

	
	fmt.Print("Masukkan Judul Yang dicari : ")
	Carijudul, _ := input.ReadString('\n')
	Carijudul = strings.TrimSpace(Carijudul)

	var pilih int 
	fmt.Println("Pilih Menu Search")
	fmt.Println("1.Binary Search")
	fmt.Println("2.Squential Search")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilih)

	var index int
	
	if pilih == 1 {
		index = Binary_Search(Carijudul) 
		ide := Daftar_Ide[index]

			if index == -1 {
				fmt.Println("Tidak ada")
			}else {
				fmt.Println("\nIde ditemukan")
				fmt.Println("Judul\t:", ide.Judul)
				fmt.Println("Deskripsi\t:", ide.Deskripsi)
				fmt.Println("Rating\t:", ide.Rating)
			}
	}else if pilih == 2 {
		index = Squential_Search(Carijudul)
		ide := Daftar_Ide[index]

			if index == -1 {
				fmt.Println("Tidak ada")
			}else {
				fmt.Println("\nIde ditemukan:")
				fmt.Println("Judul\t:", ide.Judul)
				fmt.Println("Deskripsi\t:", ide.Deskripsi)
				fmt.Println("Rating\t:", ide.Rating)
			}
	}else {
		fmt.Println("Pilihan tidak Ada")
		return
	}
}

func Selection_Sort() {		// Buat nge sort Judul

	for i := 0 ; i < len(Daftar_Ide)-1 ; i++{
		temp := i

		for j := i+1 ; j < len(Daftar_Ide) ; j++{
			if strings.ToLower(Daftar_Ide[j].Judul) < strings.ToLower(Daftar_Ide[temp].Judul){
				temp = j
			}
		}
		Daftar_Ide[i] , Daftar_Ide[temp] = Daftar_Ide[temp], Daftar_Ide[i]
	}
}

func Insertion_Sort() {  	// Buat nge sort Rating

	for i := 1 ; i < len(Daftar_Ide) ; i++{
		temp := Daftar_Ide[i]

		var j int
		j = i

		for j > 0 && Daftar_Ide[j-1].Rating < temp.Rating{
			Daftar_Ide[j] = Daftar_Ide[j-1]
			j--
		}
		Daftar_Ide[j] = temp	
	}
}

func Urutan_ide() {

	if len(Daftar_Ide) == 0{
		fmt.Println("Belum ada Ide")
		return
	}

	var pilih int 
	fmt.Printf("\nMenu Sorting\n\n ")
	fmt.Println("1.Insertion Sort (Rating terbesar ke terkecil) ")
	fmt.Println("2.Selection Sort (Urutan Judul (Urutan Alfabet) ")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		Insertion_Sort()
	}else if pilih == 2 {
		Selection_Sort()
	}else{
		fmt.Println("Pilihan tidak ada")
		return
	}

	Lihat_Daftar_Ide()
}