package repo_test

import (
	"context"
	_ "database/sql"
	"flatApp/pkg/platform/flat"
	"flatApp/tests/database"
	"flatApp/tests/testData"
	"net/http"
	"testing"
)

var ctx context.Context

func TestCreate(t *testing.T) {
	tests := []struct{
		name string
		body flat.Flat
		status int
	}{
		{
			name: "should create flat",
			body: testData.GiveTrueData(),
			status: http.StatusCreated,
		},
		//{
		//	name: "shouldn't create flat without street",
		//	body: testData.GiveDataWithoutStreet(),
		//	status: http.StatusBadRequest,
		//},
		//{
		//	name: "shouldn't create flat without house_number",
		//	body:testData.GiveDataWithoutHouseNumber(),
		//	status: http.StatusBadRequest,
		//},
		//{
		//	name: "shouldn't create flat without room_number",
		//	body:testData.GiveDataWithoutRoomNumber(),
		//	status: http.StatusBadRequest,
		//},
	}

	repo, cleanup := database.CreateTestFlatsRepository(t,"TestFlat")
	defer cleanup()

	expectedResult := flat.Flat{
		ID: 0,
		Street: "Koroleva",
		HouseNumber: "12A",
		RoomNumber: 61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID: 1,
			Country: "",
			Name: "",
		},
	}

	for _, tc := range tests{
		t.Run(tc.name, func(tt *testing.T) {
			result, _ := repo.Create(ctx, tc.body)
			if result != expectedResult {
				t.Errorf("Create() result = %v, want = %v", result, expectedResult)
			}
		})
	}
}


