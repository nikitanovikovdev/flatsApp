package testData

import "flatApp/pkg/platform/flat"

func GiveTrueData() flat.Flat{
	return flat.Flat{
		Street: "Koroleva",
		HouseNumber: "12A",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}
}

func GiveDataWithoutStreet() flat.Flat{
	return flat.Flat{
		Street: " ",
		HouseNumber: "12A",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}
}

func GiveDataWithoutHouseNumber() flat.Flat{
	return flat.Flat{
		Street: "Koroleva",
		HouseNumber: " ",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}
}

func GiveDataWithoutRoomNumber() flat.Flat{
	return flat.Flat{
		Street: "Koroleva",
		HouseNumber: "12A ",
		RoomNumber: 0,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
		},
	}
}