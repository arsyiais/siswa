package main

import (
	"fmt"
	"siswa/model"
)

// import "fmt"

// type Siswa struct {
// 	Id     int
// 	Nama   string
// 	Kelas  int
// 	Alamat string
// 	Nilai  float32
// }

// type TableSiswa struct {
// 	Data Siswa
// 	Next *TableSiswa
// }

// var DBsiswa TableSiswa

// func InsertSiswa(id int, nama string, kelas int, alamat string, nilai float32) {
// 	newDataSiswa := TableSiswa{
// 		Data: Siswa{id, nama, kelas, alamat, nilai},
// 		Next: nil,
// 	}
// 	var tempLL *TableSiswa
// 	tempLL = &DBsiswa
// 	if DBsiswa.Next == nil {
// 		DBsiswa.Next = &newDataSiswa
// 	} else {
// 		for tempLL.Next != nil {
// 			tempLL = tempLL.Next
// 		}
// 		tempLL.Next = &newDataSiswa
// 	}
// }

// func ReadAllSiswa() {
// 	var tempLL *TableSiswa
// 	tempLL = DBsiswa.Next
// 	if DBsiswa.Next == nil {
// 		fmt.Println("tidak ada data")
// 	} else {
// 		for tempLL != nil {
// 			fmt.Println("------------------------------------")
// 			fmt.Println("Id Siswa :", tempLL.Data.Id)
// 			fmt.Println("Nama Siswa :", tempLL.Data.Nama)
// 			fmt.Println("Kelas Siswa :", tempLL.Data.Kelas)
// 			fmt.Println("Alamat Siswa :", tempLL.Data.Alamat)
// 			fmt.Println("Nilai Siswa :", tempLL.Data.Nilai)
// 			tempLL = tempLL.Next
// 			fmt.Println("------------------------------------")
// 		}
// 	}
// }

// func SearchSiswa(id int) {
// 	var tempLL *TableSiswa
// 	tempLL = DBsiswa.Next
// 	cek := false
// 	if DBsiswa.Next == nil {
// 		fmt.Println("DATA SISWA KOSONG")
// 	} else {
// 		for tempLL != nil {
// 			if id == tempLL.Data.Id {
// 				cek = true
// 				fmt.Println("------------------------------------")
// 				fmt.Println("Id Siswa :", tempLL.Data.Id)
// 				fmt.Println("Nama Siswa :", tempLL.Data.Nama)
// 				fmt.Println("Kelas Siswa :", tempLL.Data.Kelas)
// 				fmt.Println("Alamat Siswa :", tempLL.Data.Alamat)
// 				fmt.Println("Nilai Siswa :", tempLL.Data.Nilai)
// 			}
// 			tempLL = tempLL.Next
// 		}
// 		if cek != true {
// 			fmt.Println("Nama tidak ditemukan")
// 		}
// 	}
// }

// func cariSiswa(id int) *TableSiswa {
// 	var tempLL *TableSiswa
// 	tempLL = DBsiswa.Next
// 	cek := false
// 	if DBsiswa.Next == nil {
// 		return nil
// 	} else {
// 		for tempLL != nil {
// 			if id == tempLL.Data.Id {
// 				cek = true
// 				return tempLL
// 			}
// 			tempLL = tempLL.Next
// 		}
// 		if cek != true {
// 			return nil
// 		}
// 	}
// 	return nil
// }

// func UpdateSiswa(id int, nama string, kelas int, alamat string, nilai float32) {
// 	siswa := cariSiswa(id)
// 	if siswa != nil {
// 		siswa.Data.Id = id
// 		siswa.Data.Nama = nama
// 		siswa.Data.Kelas = kelas
// 		siswa.Data.Alamat = alamat
// 		siswa.Data.Nilai = nilai
// 		fmt.Println("update berhasil")
// 	} else {
// 		fmt.Println("tidak ada data yang dicari")
// 	}
// }

// func DeleteSiswa(id int) {
// 	var tempLL *TableSiswa
// 	tempLL = &DBsiswa
// 	if tempLL.Next == nil {
// 		//cek data kosong
// 		fmt.Println("data kosong")
// 	} else {
// 		for tempLL.Next != nil {
// 			//fmt.Println(tempLL.Next.Data)
// 			if tempLL.Next.Data.Id == id {
// 				tempLL.Next = tempLL.Next.Next
// 				return
// 			}
// 			tempLL = tempLL.Next
// 		}
// 	}

// }

func main() {
	model.InsertSiswa(1, "Habibie", 6, "Sidoajo", 75)
	model.InsertSiswa(2, "Zanki", 1, "Sidoajo", 89)
	// model.DeleteSiswa(2)
	siswas := model.ReadAllSiswa()

	// model.SearchSiswa(1)
	// model.UpdateSiswa(2, "Zanki", 2, "Sidoajo", 83)

	// fmt.Println(siswas)
	// fmt.Println(sepatus.Next)

	if siswas != nil {
		for siswas.Next != nil {
			fmt.Println(siswas.Next.Data)
			siswas = siswas.Next
		}
	}

	// // SearchSiswa(2)
	// // UpdateSiswa(5, "Abi", 7, "Sidoajo", 93)
	// // DeleteSiswa(2)
	// ReadAllSiswa()
}
