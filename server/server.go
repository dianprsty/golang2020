package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"parking/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":9999"
)

type ParkingService struct{}

var Parking = make(map[string]time.Time)

func (ParkingService) GenerateId(ctx context.Context, empty *empty.Empty) (*model.ParkIn, error) {
	a := strconv.Itoa(rand.Intn(10000))
	b := time.Now()
	c := b.String()
	Parking[a] = b
	log.Println("Id :", a, "\t", "Time :", b)
	message := &model.ParkIn{
		Id:   a,
		Time: c,
	}
	return message, nil

}

func (ParkingService) ParkOut(ctx context.Context, input *model.InputData) (*model.Output, error) {
	id := input.Id
	c := time.Now()
	e := Parking[id]
	f := c.Sub(e)
	g := int(f / time.Second)
	strduration := f.String()
	var bill string
	tipe := input.Tipe
	platno := input.Platno
	message := ""

	if _, found := Parking[id]; found {
		switch tipe {
		case "car":
			a := (3000 * g) + 5000
			if a >= 100000 {
				bill = strconv.Itoa(100000)
				message = "Your Parking Bill Is : " + bill
			} else {
				bill = strconv.Itoa(a)
				message = "Your Parking Bill Is : " + bill
			}
			delete(Parking, id)
		case "motor":
			a := (2000 * g) + 3000
			if a >= 50000 {
				bill = strconv.Itoa(50000)
				message = "Your Parking Bill Is : " + bill
			} else {
				bill = strconv.Itoa(a)
				message = "Your Parking Bill Is : " + bill
			}
			delete(Parking, id)
		}
	} else {
		message = "Parking Id not found please input another Id"
	}
	output := &model.Output{
		Id:       id,
		Platno:   platno,
		Duration: strduration,
		Message:  message,
	}

	return output, nil

}

func main() {
	log.Println("Server Running --> PORT", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failedd to listen: %v", err)
	}

	s := grpc.NewServer()
	var parkingService ParkingService
	model.RegisterParkingServiceServer(s, parkingService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print(s.Serve(lis))
}
