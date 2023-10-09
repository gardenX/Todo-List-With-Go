package angga

import (
	"fmt"
	"os"
)

func AddList(slc1 []string, status []bool) ([]string, []bool) {
	var listNew string
	fmt.Print("Masukan Nama Tugas :")
	fmt.Scanln(&listNew)

	slc1 = append(slc1, listNew)
	status = append(status, false)

	return slc1, status
}

func ViewList(slc1 []string, status []bool) []string {
	fmt.Println("Daftar Tugas:")
	for i := 0; i < len(slc1); i++ {
		if status[i] {
			fmt.Println(i+1, slc1[i], "- Selesai")
		} else {
			fmt.Println(i+1, slc1[i], "- Belum Selesai")
		}
	}
	return slc1
}

func MarkList(slc1 []string, status []bool) ([]string, []bool) {
	//ViewList(slc1)

	fmt.Print("Pilih Daftar Tugas :")
	var NumList int
	fmt.Scanln(&NumList)
	NumList = NumList - 1

	slc1[NumList] = slc1[NumList]

	status[NumList] = !status[NumList]

	return slc1, status

}

func filterList(slc1 []string, status []bool) ([]string, []bool) {
	fmt.Println("1. Tugas Selesai")
	fmt.Println("2. Tugas Belum Selesai")
	fmt.Print("Pilih Filter :")

	var numList int
	fmt.Scan(&numList)

	var tmpSlc1 []string
	var tmpStatus []bool

	for i := 0; i < len(status); i++ {
		if status[i] == (numList == 1) {
			tmpSlc1 = append(tmpSlc1, slc1[i])
			tmpStatus = append(tmpStatus, status[i])
		}
	}
	return tmpSlc1, tmpStatus
}

func actList() {
	fmt.Println("=== To-DO List ====")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Daftar Tugas")
	fmt.Println("3. Tandai Tugas Selesai")
	fmt.Println("4. Filter Tugas")
	fmt.Println("5. Hapus Tugas")
	fmt.Println("6. Keluar")
}

func delList(slc1 []string, status []bool) ([]string, []bool) {
	fmt.Print("Pilih Daftar Tugas :")
	var NumList int
	fmt.Scanln(&NumList)

	var indDel int = NumList - 1

	slc1 = append(slc1[:indDel], slc1[indDel+1:]...)
	status = append(status[:indDel], status[indDel+1:]...)

	return slc1, status
}

func Main() {
	var slc1 []string
	var status []bool

	for {
		actList()

		var act int
		fmt.Print("Pilih Operasi (1/2/3/4/5) :")
		fmt.Scanln(&act)

		switch act {
		case 1:
			slc1, status = AddList(slc1, status)
			fmt.Println("Tugas Di Tambahkan !")
			fmt.Println("")
			fmt.Println("")
		case 2:
			ViewList(slc1, status)
			fmt.Println("")
			fmt.Println("")
		case 3:
			slc1, status = MarkList(slc1, status)
			ViewList(slc1, status)
			fmt.Println("")
			fmt.Println("")
		case 4:
			ViewList(filterList(slc1, status))
			fmt.Println("")
			fmt.Println("")
		case 5:
			slc1, status = delList(slc1, status)
			ViewList(slc1, status)
			fmt.Println("")
			fmt.Println("")
		case 6:
			os.Exit(0)
		}
	}
}
