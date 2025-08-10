package infra

import "time"

var reservations = map[int]*Reservation{
	1: {
		ID:                  1,
		RoomID:              1,
		ReservationDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		ConfirmedDateTime:   time.Date(2024, 1, 1, 5, 0, 0, 0, time.UTC),
	},
	2: {
		ID:                  2,
		RoomID:              1,
		ReservationDateTime: time.Date(2024, 2, 1, 10, 0, 0, 0, time.UTC),
		ConfirmedDateTime:   time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC),
	},
	3: {
		ID:                  3,
		RoomID:              2,
		ReservationDateTime: time.Date(2024, 5, 1, 10, 0, 0, 0, time.UTC),
		ConfirmedDateTime:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
	},
}

var rooms = map[int]*Room{
	1: {
		ID:   1,
		Name: "normal",
	},
	2: {
		ID:   2,
		Name: "special",
	},
}
