package infra

type RetriveReservation func(id int) (*Reservation, error)

func NewRetriveReservation() RetriveReservation {
	return retriveReservation
}

func retriveReservation(id int) (*Reservation, error) {
	reservation, ok := reservations[id]
	if !ok {
		return nil, ErrNotFound
	}
	return reservation, nil
}
