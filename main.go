package main

import "fmt"

func main() {
	ip := 4
	var grade string
	if ip == 4 {
		grade = "A"
		fmt.Println("Grade IP kamu adalah : ", grade)
		if grade == "A" {
			fmt.Println("Selamat kamu wisuda dengan tag Summa Cumlaude.")
		}
	} else if float64(ip) >= 3.5 {
		grade = "B+"
		fmt.Println("Grade IP kamu adalah : ", grade)
	} else if float64(ip) >= 3.0 {
		grade = "B"
		fmt.Println("Grade IP kamu adalah : ", grade)
	} else if float64(ip) >= 2.5 {
		grade = "C+"
		fmt.Println("Grade IP kamu adalah : ", grade)
	} else if float64(ip) >= 2.3 {
		grade = "C"
		fmt.Println("Grade IP kamu adalah : ", grade)
	} else if float64(ip) <= 2.3 {
		grade = "D"
		fmt.Println("Grade IP kamu adalah : ", grade)
		fmt.Println("Kamu tidak bisa mengikuti ujian.")
	} else {
		fmt.Println("Grade tidak diketahui.")
	}
}
