package infra

type RetriveRoom func(id int) (*Room, error)

func NewRetriveRoom() RetriveRoom {
	return retriveRoom
}

func retriveRoom(id int) (*Room, error) {
	room, ok := rooms[id]
	if !ok {
		return nil, ErrNotFound
	}
	return room, nil
}
