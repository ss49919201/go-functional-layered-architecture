package infra

type RetriveReservation func(id int)

func NewRetriveReservation() RetriveReservation {
	return retriveReservation
}

func retriveReservation(id int) {}
