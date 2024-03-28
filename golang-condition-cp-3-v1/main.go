package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {

	grade := ((math + science + english + indonesia) / 4)
	fmt.Println(grade)

	var result string

	if grade == 100 {
		result = "Sempurna"
	} else if grade >= 90 {
		result = "Sangat Baik"
	} else if grade >= 80 {
		result = "Baik"
	} else if grade >= 70 {
		result = "Cukup"
	} else if grade >= 60 {
		result = "Kurang"
	} else {
		result = "Sangat kurang"
	}

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
