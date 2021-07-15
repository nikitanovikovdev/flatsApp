package repo_test

import (
	"context"
	_ "database/sql"
	"flatApp/pkg/platform/flat"
	"flatApp/tests/database"
	"fmt"
	"reflect"
	"testing"
)

var ctx context.Context

//func TestCreate(t *testing.T) {
//
//	tests := []struct{
//		name string
//		body flat.Flat
//	}{
//		{
//			name: "should create flat",
//			body: testData.TrueData,
//		},
		//{
		//	name:   "shouldn't create flat without street",
		//	body:   testData.GiveDataWithoutStreet(),
		//	status: http.StatusBadRequest,
		//},
		//{
		//	name:   "shouldn't create flat without house_number",
		//	body:   testData.GiveDataWithoutHouseNumber(),
		//	status: http.StatusBadRequest,
		//},
		//{
		//	name:   "shouldn't create flat without room_number",
		//	body:   testData.GiveDataWithoutRoomNumber(),
		//	status: http.StatusBadRequest,
		//},
	//}

//	repo, cleanup := database.CreateTestFlatsRepository("creatFlat")
//	defer cleanup()
//
//	expectedResult := flat.Flat{
//		ID: 0,
//		Street: "Koroleva",
//		HouseNumber: "12A",
//		RoomNumber: 61,
//		Description: "good flat in Minsk",
//		City: flat.City{
//			ID: 1,
//			Country: "",
//			Name: "",
//		},
//	}
//
//	for _, tc := range tests {
//			fmt.Println(tc.name)
//			result, err := repo.Create(ctx, testData.TrueData)
//			if err!= nil {
//				t.Fatal(err)
//			}
//			if result != expectedResult {
//				t.Errorf("Create() result = %v, want = %v", result, expectedResult)
//			}
//			fmt.Printf("Correct result. got=%v want=%v", repo, expectedResult)
//
//	}
//
//}

func TestReadAll(t *testing.T) {

	repo, cleanup := database.CreateTestFlatsRepository("readAllFlats")
	defer cleanup()

	expectedResult := []flat.Flat{
		{
			ID: 1,
			Street: "Lenina",
			HouseNumber: "77A",
			RoomNumber: 33,
			Description: "good flat",
			City: flat.City{
				ID: 1,
				Country: "Belarus",
				Name: "Minsk",
			},
		},
		{
			ID: 2,
			Street: "Tolstogo",
			HouseNumber: "13",
			RoomNumber: 71,
			Description: "",
			City: flat.City{
				ID: 2,
				Country: "Belarus",
				Name: "Brest",
			},
		},
	}

	result, err := repo.ReadAll(ctx)
	if err != nil {
		t.Error(err)
	}

	compare := reflect.DeepEqual(result, expectedResult)

	if compare {
		fmt.Printf("Correct result. got=%v want=%v", repo, expectedResult)
	}
	t.Errorf("Create() result = %v, want = %v", result, expectedResult)

}

//func TestRead(t *testing.T) {
//	tests := []struct {
//		name string
//		id string
//	}{
//		{
//			name: "should return flat",
//			id: "1",
//		},
//		{
//			name: "shouldn't return flat",
//			id: "12",
//		},
//	}
//
//	repo, cleanup := database.CreateTestFlatsRepository("readFlat")
//	defer cleanup()
//
//	expectedResult := flat.Flat{
//			ID:          1,
//			Street:      "Lenina",
//			HouseNumber: "77A",
//			RoomNumber:  33,
//			Description: "good flat",
//			City: flat.City{
//				ID:      1,
//				Country: "Belarus",
//				Name:    "Minsk",
//		},
//	}
//
//	for _, tc := range tests {
//
//			result, err := repo.Read(ctx, tc.id)
//			if err != nil {
//				t.Error(err)
//			}
//
//			if result != expectedResult {
//				t.Errorf("ReadAll() result = %v, want = %v", result, expectedResult)
//			}
//			fmt.Printf("Correct result. got=%v want=%v", repo, expectedResult)
//	}
//}
//
//func TestUpdate(t *testing.T) {
//	tests := []struct {
//		name string
//		id string
//		body flat.Flat
//	}{
//		{
//			name: "should return flat",
//			id: "1",
//			body: testData.TrueData,
//		},
//		{
//			name: "shouldn't return flat",
//			id: "12",
//			body: testData.TrueData,
//		},
//	}
//
//	repo, cleanup := database.CreateTestFlatsRepository("updateFlat")
//	defer cleanup()
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			err := repo.Update(ctx, tc.id, tc.body)
//			if err != nil {
//				t.Errorf("Incorrect result %s",err)
//			}
//		})
//	}
//}

//func TestDelete(t *testing.T) {
//	tests := []struct {
//		name string
//		id string
//	}{
//		{
//			name: "should return flat",
//			id: "2",
//		},
//		{
//			name: "shouldn't return flat",
//			id: "12",
//		},
//	}
//
//	repo, cleanup := database.CreateTestFlatsRepository("deleteFlat")
//	defer cleanup()
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			err := repo.Delete(ctx, tc.id)
//			if err != nil {
//				t.Errorf("Incorrect result %s",err)
//			}
//		})
//	}
//}


