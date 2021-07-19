package flats_test

import (
	"context"
	"flatApp/pkg/flats"
	"flatApp/pkg/platform/flat"
	"flatApp/tests/database"
	"flatApp/tests/testData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()

	tests := []struct{
		name string
		body flat.Flat
		expectedError bool
	} {
		{
			name: "should create flat",
			body: testData.GiveTrueDataForService,
			expectedError: false,
		},
		{
			name: "shouldn't create flat",
			body: testData.GiveDataWithoutStreet,
			expectedError: true,
		},
	}

	expectedResult := flat.Flat{
		Street: "Mira",
		HouseNumber: "99",
		RoomNumber: 98,
		Description: "test description",
		City: flat.City{
			ID: 2,
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("createFlat")
	defer cleanup()
	service := flats.NewService(repo)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := service.Create(ctx,testData.ConvertToBytes(tc.body))

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResult, result)
			}
		})
	}
}

func TestService_Read(t *testing.T) {
	ctx := context.Background()

	tests := []struct{
		name string
		id string
		expectedError bool
	}{
		{
			name: "should return flat",
			id : "1",
			expectedError: false,
		},
		{
			name: "shouldn't return flat",
			id : "11",
			expectedError: true,
		},
	}

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
	repo, cleanup := database.CreateTestFlatsRepository("readFlat")
	service := flats.NewService(repo)
	defer cleanup()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := service.Read(ctx, tc.id)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResult, result)
			}
		})
	}
}

func TestService_ReadAll(t *testing.T) {
	ctx := context.Background()

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

	repo, cleanup := database.CreateTestFlatsRepository("testReadAll")
	service := flats.NewService(repo)
	defer cleanup()

	result, err := service.ReadAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestService_Update(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		id string
		body flat.Flat
	}{
		{
			name: "should update flat",
			id: "1",
			body: testData.GiveTrueDataForService,
		},
		{
			name: "shouldn't update flat",
			id: "12",
			body: testData.GiveTrueDataForService,
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("updateFlat")
	service := flats.NewService(repo)
	defer cleanup()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Update(ctx, tc.id, testData.ConvertToBytes(tc.body))
			if err != nil {
				assert.Errorf(t, err, "Incorrect result")
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		id string
	}{
		{
			name: "should delete flat",
			id: "2",
		},
		{
			name: "shouldn't delete flat",
			id: "12",
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("deleteFlat")
	service := flats.NewService(repo)
	defer cleanup()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Delete(ctx, tc.id)
			if err != nil {
				assert.Errorf(t, err, "Incorrect result")
			}
		})
	}
}


