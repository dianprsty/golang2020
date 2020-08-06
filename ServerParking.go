package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Parking = make(map[string]time.Time)

type SendIn struct {
	Status  int64
	Id      string
	Time    time.Time
	Message string
}

type SendOut struct {
	Status  int64
	Id      string
	Message string
}
type Out struct {
	Id     int64  `json:"id"`
	Tipe   string `json:"tipe"`
	PlatNo string `json:"plat"`
}

func main() {
	http.HandleFunc("/get_id", generateId)
	http.HandleFunc("/get_total", checkOut)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
func generateId(w http.ResponseWriter, r *http.Request) {
	id := ParkIn()
	x := timeIn()
	Parking[id] = x
	send_id := SendIn{200, id, x, "Success!!!, welcome to our Parking Area!"}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(send_id)
	w.Write(js)
}
func checkOut(w http.ResponseWriter, r *http.Request) {
	resp, _ := ioutil.ReadAll(r.Body)
	var out Out
	err := json.Unmarshal(resp, &out)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := ParkOut(out.Tipe, out.PlatNo, strconv.Itoa(int(out.Id)))
	var result = SendOut{200, strconv.Itoa(int(out.Id)), res}
	js, _ := json.Marshal(result)
	w.Write(js)
}

func ParkIn() string {
	a := strconv.Itoa(rand.Intn(1000))
	return a
}

func timeIn() time.Time {
	timeIn := time.Now()
	time := timeIn.String()
	fmt.Println(time)
	return timeIn

}

func ParkOut(tipe, platno, id_parkir string) string {
	c := time.Now()
	time2 := c.String()
	e := Parking[id_parkir]
	f := c.Sub(e)
	g := int(f / time.Second)
	bill := 0
	fmt.Println("Time In :", e)
	fmt.Println("Time Out :", c)
	fmt.Println("Duration :", f)
	fmt.Println("time 2:", time2)

	if _, found := Parking[id_parkir]; found {
		switch tipe {
		case "car":
			a := (3000 * g) + 5000
			if a >= 100000 {
				bill = 100000
				fmt.Println("Your Parking Bill Is : ", bill)
			} else {
				bill = a
				fmt.Println("Your Parking Bill Is : ", bill)
			}
			delete(Parking, id_parkir)
		case "motor":
			a := (2000 * g) + 3000
			if a >= 50000 {
				bill = 50000
				fmt.Println("Your Parking Bill Is : ", bill)
			} else {
				bill = a
				fmt.Println("Your Parking Bill Is : ", bill)
			}
			delete(Parking, id_parkir)
		}
		return "Your Parking Bill Is:" + strconv.Itoa(bill)
	} else {
		return "Parking Id not found please input another Id"
	}

}
