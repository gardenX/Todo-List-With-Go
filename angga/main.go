package main

import "fmt"

func dtAddList() {

}

func AddList(param int) {
	if param == 1 {
		var param string
		fmt.Print("Masukan Nama Tugas :")
		fmt.Scanln(&param)

		var add = [...]string{
			param,
		}
		var dtlist = add[:]
		fmt.Println(dtlist)
	} else {
		fmt.Println("2aa")
	}
}

func proccesAct(params int) {
	switch params {
	case 1:
		AddList(params)
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("D")
	case 5:
	}
}

func actList() {
	fmt.Println("=== To-DO List ====")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Daftar Tugas")
	fmt.Println("3. Tandai Tugas Selesai")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("5. Keluar")

	var act int
	fmt.Print("Pilih Operasi (1/2/3/4/5) :")
	fmt.Scanln(&act)
	proccesAct(act)
}

func main() {
	actList()
}
