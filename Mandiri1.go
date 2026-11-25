package main
import "fmt"

func main() {
	var jenis string
	fmt.Println("A. Mobil")
	fmt.Println("B. Motor")
	fmt.Println("C. Sepeda")
	fmt.Println("D. Bus")

	fmt.Scan(&jenis)
	switch jenis {
	case "A":
		fmt.Println("MOBIL")
	case "B":
		fmt.Println("MOTOR")
	case "C":
		fmt.Println("SEPEDA")
	case "D":
		fmt.Println("BUS")
	default:
		fmt.Println("tidak ada kendaraan")
	}
}

// MUHAMMAD NABIL HAKFANI 1090826000111
