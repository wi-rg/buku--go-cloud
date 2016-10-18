package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println("Inisialisasi nilai dengan angka 100")
	fmt.Println("=> nilai := 100")
	fmt.Println("Setelah ini, 'nilai' berada di suatu lokasi memory")
	fmt.Println("dan ada angka int yang menempati lokasi memory tersebut")
	nilai := 100
	fmt.Println("Isi nilai = ", nilai)
	fmt.Println("Nilai tersebut berada di lokasi memory = ", &nilai)
	fmt.Println()

	fmt.Println("Memesan tempat di memory bertipe int")
	fmt.Println("bernama ptrNilai")
	fmt.Println("=> var ptrNilai *int")
	fmt.Println("Setelah ini, 'ptrNilai' berada di suatu lokasi memory")
	fmt.Println("tapi belum ada nilai yang menempati lokasi memory tsb")
	var ptrNilai *int
	fmt.Println("ptrNilai berada di lokasi memory", &ptrNilai)
	fmt.Println("ptrNilai sudah berada di lokasi memory tapi belum menunjuk ke mana2")
	fmt.Println("Isi ptrNilai = ", ptrNilai)
	fmt.Println("Jika memaksa mengambil nilai di lokasi tersebut, maka panic.")
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x401628]
	// uncomment di bawah ini jika ingin merasakan error:
	//fmt.Println("Isi dari lokasi memory tersebut = ", *ptrNilai)

	fmt.Println("ptrNilai menunjuk ke memory tempat nilai berada")
	fmt.Println("ptrNilai = &nilai")
	// &nilai berisi lokasi memory tempat menyimpan
	// nilai (100).
	ptrNilai = &nilai
	fmt.Println()

	fmt.Println("Hasilnya adalah sbb:")
	fmt.Println("--------------------")
	fmt.Println("Isi dari nilai = ", nilai)
	fmt.Println("nilai tersebut berada pada lokasi memory nilai = ", &nilai)
	fmt.Println("ptrNilai sudah menunjuk ke lokasi memory nilai, yaitu ", ptrNilai)
	fmt.Println("Isi lokasi memory ptrNilai (", ptrNilai, ") = ", *ptrNilai)
	fmt.Println()

	fmt.Println("Saat isi nilai diubah menjadi 150,")
	nilai = 150
	fmt.Println("hasilnya adalah sbb:")
	fmt.Println("--------------------")
	fmt.Println("Isi dari nilai = ", nilai)
	fmt.Println("nilai tersebut berada pada lokasi memory nilai = ", &nilai)
	fmt.Println("ptrNilai sudah menunjuk ke lokasi memory nilai, yaitu ", ptrNilai)
	fmt.Println("Isi lokasi memory ptrNilai (", ptrNilai, ") = ", *ptrNilai)
	fmt.Println()

}
