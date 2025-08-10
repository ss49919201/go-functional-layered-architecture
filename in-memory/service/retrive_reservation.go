package service

import (
	"time"

	"github.com/ss49919201/go-functional-layerd-architecture/in-memory/infra"
)

type RetriveReservationResult struct {
	ID                  int
	RoomName            int
	ReservationDateTime time.Time
	ConfirmedDateTime   time.Time
}

type RetriveReservation func(id int)

func NewRetriveReservation(
	retriveReservation infra.RetriveReservation,
	retriveRoom infra.RetriveRoom,
) RetriveReservation

func retriveReservation(id int) {}
