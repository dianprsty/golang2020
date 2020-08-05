package main

import (
	"fmt"
	"time"
)

type park struct {
	id           int
	platno, tipe string
	time         time.Time
}

func main() {
	Parking := make(map[int]time.Time)
	var x int
	y := 1
	i := true
	for i {
		fmt.Println("Menu")
		fmt.Println("1. Park In")
		fmt.Println("2. Park Out")
		fmt.Println("3. List Park")
		fmt.Println("Select [1/2/3] = ")
		fmt.Println("[Press other number key to exit]")
		fmt.Scan(&x)
		switch x {
		case 1:
			a := generateId(y)
			b := time.Now()
			Parking[a] = b
			fmt.Println("id parkir : ", a, "\nTime In : ", b)
			y++
			i = true
		case 2:
			var (
				id           int
				platno, tipe string
			)
			c := time.Now()
			fmt.Println("Input Id")
			fmt.Scan(&id)
			fmt.Println("Input vehicle type [car/motor]")
			fmt.Scan(&tipe)
			fmt.Println("Input platno")
			fmt.Scan(&platno)
			e := Parking[id]
			f := c.Sub(e)
			g := int(f / time.Second)

			fmt.Println("Time In :", e)
			fmt.Println("Time Out :", c)
			fmt.Println("Duration :", f)

			switch tipe {
			case "car":
				a := (3000 * g) + 5000
				if a >= 100000 {
					bill := 100000
					fmt.Println("Your Parking Bill Is : ", bill)
				} else {
					bill := a
					fmt.Println("Your Parking Bill Is : ", bill)
				}
				delete(Parking, id)
			case "motor":
				a := (2000 * g) + 3000
				if a >= 50000 {
					bill := 50000
					fmt.Println("Your Parking Bill Is : ", bill)
				} else {
					bill := a
					fmt.Println("Your Parking Bill Is : ", bill)
				}
				delete(Parking, id)
			default:
				fmt.Println("input vehicle type is wrong!")
				break
			}
		case 3:
			for key, value := range Parking {
				fmt.Println("Id :", key, "\n", "Time In :", value)
			}
		default:
			i = false
		}

	}
}
func generateId(id int) int {
	id = id
	return id
}

//==============================DIAN PRASETYO==========================================
