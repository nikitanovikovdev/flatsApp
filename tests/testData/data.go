package testData

import (
	"encoding/json"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/flat"
	"log"
)

var us = flat.Username{
	Username: "lera",
}

var flTrue = flat.Flat{
	Street:      "Koroleva",
	HouseNumber: "12A",
	RoomNumber:  61,
	Description: "good flat in Minsk",
	City: flat.City{
		ID: 1,
	},
}

var flWithoutStreet = flat.Flat{
	Street:      " ",
	HouseNumber: "12A",
	RoomNumber:  61,
	Description: "good flat in Minsk",
	City: flat.City{
		ID: 1,
	},
}

var flWithoutHouseNumber = flat.Flat{
	Street:      "Koroleva",
	HouseNumber: " ",
	RoomNumber:  61,
	Description: "good flat in Minsk",
	City: flat.City{
		ID: 1,
	},
}

var flWithoutRoomNumber = flat.Flat{
	Street:      "Koroleva",
	HouseNumber: "12A ",
	RoomNumber:  0,
	Description: "good flat in Minsk",
	City: flat.City{
		ID: 1,
	},
}

var GiveTrueDataForRepo = flat.FlatWithUsername{
	flTrue,
	us,
}

var GiveDataWithoutStreet = flat.FlatWithUsername{
	flWithoutStreet,
	us,
}

var GiveDataWithoutHouseNumber = flat.FlatWithUsername{
	flWithoutHouseNumber,
	us,
}

var GiveDataWithoutRoomNumber = flat.FlatWithUsername{
	flWithoutRoomNumber,
	us,
}

var GiveCorrectUsername = flat.Username{
	Username: "lera",
}

var GiveIncorrectUsername = flat.Username{
	Username: "petr",
}

var GiveTrueDataForService = flat.Flat{
	Street:      "Mira",
	HouseNumber: "99",
	RoomNumber:  98,
	Description: "test description",
	City: flat.City{
		ID: 2,
	},
}

var GiveIncorrectDataForService = flat.Flat{
	Street:      "",
	HouseNumber: "99",
	RoomNumber:  98,
	Description: "test description",
	City: flat.City{
		ID: 2,
	},
}

func ConvertToBytes(f flat.Flat) []byte {
	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	return body
}