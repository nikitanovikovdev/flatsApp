package flat

type Flat struct {
	ID          int    `json:"id"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	RoomNumber  int    `json:"room_number"`
	Description string `json:"description"`
	City        City   `json:"city"`
}

type City struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	Name    string `json:"name"`
}
