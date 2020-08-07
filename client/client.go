package main

import (
	"context"
	"fmt"
	"log"
	"parking/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
)

func main() {
	//Set Up Connection
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := model.NewParkingServiceClient(conn)

	//Menu Parkir
	var x int
	i := true
	for i {
		fmt.Println("Menu")
		fmt.Println("1. Park In")
		fmt.Println("2. Park Out")
		fmt.Println("Select [1/2] = ")
		fmt.Println("[Press any other number key to exit]")
		fmt.Scan(&x)
		switch x {
		case 1:
			a, b := GetId(c)
			fmt.Println("Id :", a)
			fmt.Println("Time In :", b)

		case 2:
			var (
				id           string
				platno, tipe string
			)
			fmt.Println("Input Id")
			fmt.Scan(&id)
			fmt.Println("Input vehicle type [car/motor]")
			fmt.Scan(&tipe)
			fmt.Println("Input platno")
			fmt.Scan(&platno)
			input := &model.InputData{
				Id:     id,
				Tipe:   tipe,
				Platno: platno,
			}
			result := CheckOut(c, input)
			fmt.Println(result)

		default:
			i = false
		}

	}

}

func GetId(c model.ParkingServiceClient) (string, string) {
	resp, err := c.GenerateId(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf(err.Error())
	}
	return resp.Id, resp.Time
}

func CheckOut(c model.ParkingServiceClient, input *model.InputData) string {
	resp, err := c.ParkOut(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return "Id :" + resp.Id + "\n" +
		"Platno :" + resp.Platno + "\n" +
		"Duration :" + resp.Duration + "\n" +
		"Message :" + resp.Message
}
