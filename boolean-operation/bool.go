package main

import "fmt"

func main() {
	var nilai_ujian = 80
	var nilai_absensi = 85

	var kkm_nilai_ujian bool = nilai_ujian > 75
	var kkm_nilai_absensi bool = nilai_absensi > 80

	var naik_kelas bool = kkm_nilai_ujian && kkm_nilai_absensi
	fmt.Println("Naik Kelas : ", naik_kelas)
}