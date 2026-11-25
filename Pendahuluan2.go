package main
import "fmt"

func main() {
	var nilai int
	var predikat string
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	switch {
	case nilai >= 85 && nilai <= 100:
		predikat = "Predikat A"
	case nilai >= 75 && nilai < 85:
		predikat = "Predikat B"
	case nilai >= 65 && nilai < 75:
		predikat = "Predikat C"
	case nilai >= 50 && nilai < 65:
		predikat = "Predikat D"
	case nilai >= 0 && nilai < 50:
		predikat = "Predikat E"
	default:
		predikat = "Tidak Ada Nilai"
	}
	fmt.Println(predikat)
}

// MUHAMMAD NABIL HAKFANI 109082600111
