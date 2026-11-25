package main
import "fmt"

func main() {
	var kode,jumlah int
	var genre string
	var harga, total int
	fmt.Println("=== BIOSKOP ===")	
	fmt.Println("1. Action")
	fmt.Println("2. Comedy")

	fmt.Print("Pilih Gendre Film :")
	fmt.Scan(&kode)
	fmt.Print("Jumlah Tiket :")
	fmt.Scan(&jumlah)

	switch kode {
	case 1 :
		harga = 50000
		genre = "Action"
	case 2 :
		harga = 40000
		genre = "Comedy"
	default :
		fmt.Println("Pilihan Tidak Tersedia")
		return
	}
	
	if jumlah <= 0 {
		fmt.Println("Jumlah Tiket Tidak Valid")
		return
	}
	total = harga * jumlah
	fmt.Println("Genre :", genre)
	fmt.Println("Harga :", harga)
	fmt.Println("Jumlah Tiket :", jumlah)
	fmt.Println("Total Pembayaran :", total)
	
}

// MUHAMMAD NABIL HAKFANI 109082600111
