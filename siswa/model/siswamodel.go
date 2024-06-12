package model

import (
	"fmt"
	"siswa/database"
	"siswa/node"
)

func InsertSiswa(id int, nama string, kelas int, alamat string, nilai float32) {
	newDataSiswa := node.TableSiswa{
		Data: node.Siswa{id, nama, kelas, alamat, nilai},
		Next: nil,
	}
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if database.DBsiswa.Next == nil {
		database.DBsiswa.Next = &newDataSiswa
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataSiswa
	}
}
func ReadAllSiswa() *node.TableSiswa {
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if database.DBsiswa.Next == nil {
		return nil
	} else {
		return tempLL

	}
}
func SearchSiswa(id int) {
	var tempLL *node.TableSiswa
	tempLL = database.DBsiswa.Next
	cek := false
	if database.DBsiswa.Next == nil {
		fmt.Println("DATA SISWA KOSONG")
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				fmt.Println("------------------------------------")
				fmt.Println("Id Siswa :", tempLL.Data.Id)
				fmt.Println("Nama Siswa :", tempLL.Data.Nama)
				fmt.Println("Kelas Siswa :", tempLL.Data.Kelas)
				fmt.Println("Alamat Siswa :", tempLL.Data.Alamat)
				fmt.Println("Nilai Siswa :", tempLL.Data.Nilai)
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			fmt.Println("Nama tidak ditemukan")
		}
	}
}

func cariSiswa(id int) *node.TableSiswa {
	var tempLL *node.TableSiswa
	tempLL = database.DBsiswa.Next
	cek := false
	if database.DBsiswa.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}
func UpdateSiswa(id int, nama string, kelas int, alamat string, nilai float32) {
	siswa := cariSiswa(id)
	if siswa != nil {
		siswa.Data.Id = id
		siswa.Data.Nama = nama
		siswa.Data.Kelas = kelas
		siswa.Data.Alamat = alamat
		siswa.Data.Nilai = nilai
		fmt.Println("update berhasil")
	} else {
		fmt.Println("tidak ada data yang dicari")
	}
}
func DeleteSiswa(id int) {
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if tempLL.Next == nil {
		//cek data kosong
		fmt.Println("data kosong")
	} else {
		for tempLL.Next != nil {
			//fmt.Println(tempLL.Next.Data)
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return
			}
			tempLL = tempLL.Next
		}
	}

}
