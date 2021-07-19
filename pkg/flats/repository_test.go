package flats_test

import (
	"context"
	_ "database/sql"
	"flatApp/pkg/platform/flat"
	"flatApp/tests/database"
	"flatApp/tests/testData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		body flat.Flat
		expectedError bool
	}{
		{
			name: "should create flat",
			body: testData.GiveTrueDataForRepo,
			expectedError: false,
		},
		{
			name: "shouldn't create flat without street",
			body: testData.GiveDataWithoutStreet,
			expectedError: true,
		},
		{
			name: "shouldn't create flat without house_number",
			body: testData.GiveDataWithoutHouseNumber,
			expectedError: true,
		},
		{
			name: "shouldn't create flat without room_number",
			body: testData.GiveDataWithoutRoomNumber,
			expectedError: true,
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("creatFlat")
	defer cleanup()

	expectedResult := flat.Flat{
		ID:          0,
		Street:      "Koroleva",
		HouseNumber: "12A",
		RoomNumber:  61,
		Description: "good flat in Minsk",
		City: flat.City{
			ID:      1,
			Country: "",
			Name:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Create(ctx, tc.body)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResult,result)
			}
		})
	}
}

func TestRepository_ReadAll(t *testing.T) {
	ctx := context.Background()

	repo, cleanup := database.CreateTestFlatsRepository("testReadAll")
	defer cleanup()

	expectedResult := []flat.Flat{
		{
			ID:          1,
			Street:      "Lenina",
			HouseNumber: "77A",
			RoomNumber:  33,
			Description: "good flat",
			City: flat.City{
				ID:      1,
				Country: "Belarus",
				Name:    "Minsk",
			},
		},
		{
			ID:          2,
			Street:      "Tolstogo",
			HouseNumber: "13",
			RoomNumber:  71,
			Description: "",
			City: flat.City{
				ID:      2,
				Country: "Belarus",
				Name:    "Brest",
			},
		},
		{
			ID:          3,
			Street:      "Dimitrova",
			HouseNumber: "13",
			RoomNumber:  6,
			Description: "bad flat",
			City: flat.City{
				ID:      3,
				Country: "Belarus",
				Name:    "Gomel",
			},
		},
	}

	result, err := repo.ReadAll(ctx)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, result)
}

func TestRepository_Read(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		id   string
		expectedError bool
	}{
		{
			name: "should return flat",
			id:   "1",
			expectedError: false,
		},
		{
			name: "shouldn't return flat",
			id:   "12",
			expectedError: true,
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("readFlat")
	defer cleanup()

	expectedResult := flat.Flat{
		ID:          1,
		Street:      "Lenina",
		HouseNumber: "77A",
		RoomNumber:  33,
		Description: "good flat",
		City: flat.City{
			ID:      1,
			Country: "Belarus",
			Name:    "Minsk",
		},
	}

	for _, tc := range tests {
		result, err := repo.Read(ctx, tc.id)
		if tc.expectedError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, expectedResult, result)
		}
	}
}

func TestRepository_Update(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		id   string
		body flat.Flat
	}{
		{
			name: "should update flat",
			id:   "1",
			body: testData.GiveTrueDataForRepo,
		},
		{
			name: "shouldn't update flat",
			id:   "12",
			body: testData.GiveTrueDataForRepo,
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("updateFlat")
	defer cleanup()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Update(ctx, tc.id, tc.body)
			if err != nil {
				assert.Errorf(t, err, "Incorrect result")
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		id   string
	}{
		{
			name: "should delete flat",
			id:   "2",
		},
		{
			name: "shouldn't delete flat",
			id:   "12",
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("deleteFlat")
	defer cleanup()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Delete(ctx, tc.id)
			if err != nil {
				assert.Errorf(t, err, "Incorrect result")
			}
		})
	}
}
