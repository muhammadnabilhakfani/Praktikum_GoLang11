package main
import "fmt"

func main() {
	var ph float64
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6 :
			fmt.Println("Air Layak Minum")
	case (ph < 6.5 || ph > 8.6 ) && (ph >= 0 && ph <= 14) :
			fmt.Println("Air Tidak Layak Minum")
	default:
		fmt.Println("Kadar Ph tidak terdeteksi")
	}
}

// MUHAMMAD NABIL HAKFANI 109082600111
