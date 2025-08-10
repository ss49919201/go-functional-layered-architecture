package infra

type RetriveRoom func(id int)

func NewRetriveRoom() RetriveRoom {
	return retriveRoom
}

func retriveRoom(id int) {}
