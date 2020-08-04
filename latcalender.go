package main

import (
	"fmt"
	"time"
)

func hawal() {
	for a := 1; a < 13; a++ {
		time.Date(2020, time.Month(a), 1, 0, 0, 0, 0, time.UTC)
	}
}

func main() {

	//date := time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	jumlahhari := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for i := 0; i < len(jumlahhari); i++ {
		fmt.Println("\n=====", time.Month(i+1), "=====")
		fmt.Println("S  S  R  K  J  S  M")

		m := 0
		for l := 1; l <= jumlahhari[i]; l++ {
			dt := time.Date(2020, time.Month(i), (l), 0, 0, 0, 0, time.UTC)
			k2 := int(dt.Weekday())
			k := dt.Day()
			if l == 1 {
				for x := 0; x <= k2+1; x++ {
					fmt.Print("-- ")
					m++
					if m%7 == 0 {
						println("")
						m = 0
					}
				}
			}

			if k < 10 {
				print(k, "  ")
			} else {
				print(k, " ")
			}
			m++
			if m%7 == 0 {
				println("")
				m = 0
			}

		}
	}
}
