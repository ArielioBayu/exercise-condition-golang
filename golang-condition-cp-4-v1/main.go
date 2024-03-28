package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {

	qty := (VIP + regular + student)
	totalPembelian := ((VIP * 30) + (regular * 20) + (student * 10))

	if totalPembelian >= 100 {
		if day%2 != 0 {
			if qty >= 5 {
				return (float32(totalPembelian)) - (float32(totalPembelian) * 25 / 100)
			} else {
				return (float32(totalPembelian)) - (float32(totalPembelian) * 15 / 100)
			}
		} else {
			if qty >= 5 {
				return (float32(totalPembelian)) - (float32(totalPembelian) * 20 / 100)

			} else {
				return (float32(totalPembelian)) - (float32(totalPembelian) * 10 / 100)
			}
		}
	} else {
		return (float32(totalPembelian)) // TODO: replace this
	}

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
	fmt.Println(GetTicketPrice(4, 0, 0, 21))
}
