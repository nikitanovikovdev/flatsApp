package testData

import "flatApp/pkg/platform/flat"

var TrueData = flat.Flat{
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
