package service

import (
	"time"

	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/infra"
)

type RetriveReservationResult struct {
	ID                  int
	RoomName            string
	ReservationDateTime time.Time
	ConfirmedDateTime   time.Time
}

type RetriveReservation func(id int) (*RetriveReservationResult, error)

func NewRetriveReservation(
	retriveReservation infra.RetriveReservation,
	retriveRoom infra.RetriveRoom,
) RetriveReservation {
	return func(id int) (*RetriveReservationResult, error) {
		return retriveReservationImpl(id, retriveReservation, retriveRoom)
	}
}

func retriveReservationImpl(
	id int,
	retriveReservation infra.RetriveReservation,
	retriveRoom infra.RetriveRoom,
) (*RetriveReservationResult, error) {
	reservation, err := retriveReservation(id)
	if err != nil {
		return nil, err
	}
	
	room, err := retriveRoom(reservation.RoomID)
	if err != nil {
		return nil, err
	}
	
	return &RetriveReservationResult{
		ID:                  reservation.ID,
		RoomName:            room.Name,
		ReservationDateTime: reservation.ReservationDateTime,
		ConfirmedDateTime:   reservation.ConfirmedDateTime,
	}, nil
}
