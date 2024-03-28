package main

import "fmt"

func TicketPlayground(height, age int) int {

	var price = 10000

	if age < 5 {
		return -1
	}

	if age >= 5 || height > 120 {
		price += 5000
	}
	if age >= 8 || height > 135 {
		price += 10000
	}
	if age >= 10 || height > 150 {
		price += 15000
	}
	if age >= 12 || height > 160 {
		price += 20000
	}
	if age > 12 {
		price = 100000
	}
	return price // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
	fmt.Println(TicketPlayground(165, 10))
	fmt.Println(TicketPlayground(140, 13))
	fmt.Println(TicketPlayground(85, 4))
}
