package angga

import (
	"fmt"
	"os"
)

//func AddList(param int) {
//	if param == 1 {
//		var param string
//		fmt.Print("Masukan Nama Tugas :")
//		fmt.Scanln(&param)
//
//		var add = [...]string{
//			param,
//		}
//		var dtlist = add[:]
//		fmt.Println(dtlist)
//	} else {
//		fmt.Println("2")
//	}
//}
//
//func proccesAct(params int) {

//}

func AddList(slc1 []string) []string {
	var listNew string
	fmt.Print("Masukan Nama Tugas :")
	fmt.Scanln(&listNew)
	slc1 = append(slc1, listNew)
	return slc1
}

func ViewList(slc1 []string) []string {
	fmt.Println("Daftar Tugas:")
	for i := 0; i < len(slc1); i++ {
		fmt.Println(i+1, ".", slc1[i], "- Belum Selesai")
	}
	return slc1
}

func MarkList(slc1 []string) {
	ViewList(slc1)

	fmt.Print("Pilih Daftar Tugas :")
	var NumList int
	fmt.Scanln(&NumList)

	slc1[NumList] = "Selesai"
	fmt.Print(slc1)
}

func actList() {
	fmt.Println("=== To-DO List ====")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Daftar Tugas")
	fmt.Println("3. Tandai Tugas Selesai")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("5. Keluar")
}

func Main() {

	var slc1 []string

	for {
		actList()

		var act int
		fmt.Print("Pilih Operasi (1/2/3/4/5) :")
		fmt.Scanln(&act)

		switch act {
		case 1:
			slc1 = AddList(slc1)
			fmt.Println("Tugas Di Tambahkan !")
			fmt.Println("")
			fmt.Println("")
		case 2:
			ViewList(slc1)
			fmt.Println("")
			fmt.Println("")
		case 3:
			MarkList(slc1)
			fmt.Println("C")
		case 4:
			fmt.Println("D")
		case 5:
			os.Exit(0)
		}
	}
}
