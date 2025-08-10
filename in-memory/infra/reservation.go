package infra

import "time"

type Reservation struct {
	ID                  int
	RoomID              int
	ReservationDateTime time.Time
	ConfirmedDateTime   time.Time
}
