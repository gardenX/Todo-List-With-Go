package main

import "fmt"

func AddList(param1 string) {
	var add = [...]string{
		param1,
	}
	var dtlist = add[:]
	fmt.Println(dtlist)
}

func proccesAct(params int) {
	for {
		switch params {
		case 1:
			var param string
			fmt.Print("Masukan Nama Tugas :")
			fmt.Scanln(&param)
			AddList(param)
			fmt.Println("ditambahkan sebagai Tugas !")
			actList()
		case 2:
			fmt.Println("B")
		case 3:
			fmt.Println("C")
		case 4:
			fmt.Println("D")
		case 5:
		}
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
