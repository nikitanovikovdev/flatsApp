package testData

import (
	"encoding/json"
	"flatApp/pkg/platform/flat"
	"log"
)

var GiveTrueDataForRepo = flat.Flat{
	Street: "Koroleva",
	HouseNumber: "12A",
	RoomNumber: 61,
	Description: "good flat in Minsk",
	City: flat.City{
		ID: 1,
	},
}

var GiveDataWithoutStreet = flat.Flat{
		Street: " ",
		HouseNumber: "12A",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}


var GiveDataWithoutHouseNumber = flat.Flat{
		Street: "Koroleva",
		HouseNumber: " ",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}


var GiveDataWithoutRoomNumber = flat.Flat{
		Street: "Koroleva",
		HouseNumber: "12A ",
		RoomNumber: 0,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}


var GiveTrueDataForService = flat.Flat{
	Street: "Mira",
	HouseNumber: "99",
	RoomNumber: 98,
	Description: "test description",
	City: flat.City{
		ID: 2,
	},
}


func ConvertToBytes(f flat.Flat) []byte{
	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

