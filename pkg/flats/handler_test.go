package flats_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/flats"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/flat"
	"github.com/nikitanovikovdev/flatsApp-flats/tests/database"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		name    string
		body    flat.Flat
		status  int
		message []byte
	}{
		{
			name: "should create",
			body: flat.Flat{
				Street:      "Pirogova",
				HouseNumber: "33",
				RoomNumber:  69,
				Description: "test description",
				City: flat.City{
					ID: 1,
				},
			},
			status:  http.StatusCreated,
			message: []byte(`{"id":0,"street":"Pirogova","house_number":"33","room_number":69,"description":"test description","city":{"id":1,"country":"","name":""}}`),
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("createFlat")
	defer cleanup()
	s := flats.NewService(repo)
	h := flats.NewHandler(s)
	r := mux.NewRouter()
	r.Handle("/flats", h.Create())

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			b, err := json.Marshal(tc.body)
			require.NoError(tt, err)

			body := ioutil.NopCloser(bytes.NewBuffer(b))

			req, err := http.NewRequest(http.MethodPost, "/flats", body)
			require.NoError(tt, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			message, err := ioutil.ReadAll(rec.Body)

			require.NoError(tt, err, "response cannot be empty")

			require.Equal(tt, tc.status, rec.Code)
			require.Equal(tt, tc.message, message)
		})
	}
}

func TestHandler_ReadAll(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		message []byte
	}{
		{
			name:    "Should return info about flat",
			status:  http.StatusOK,
			message: []byte(`[{"id":1,"street":"Lenina","house_number":"77A","room_number":33,"description":"good flat","city":{"id":1,"country":"Belarus","name":"Minsk"}},{"id":2,"street":"Tolstogo","house_number":"13","room_number":71,"description":"","city":{"id":2,"country":"Belarus","name":"Brest"}},{"id":3,"street":"Dimitrova","house_number":"13","room_number":6,"description":"bad flat","city":{"id":3,"country":"Belarus","name":"Gomel"}}]`),
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("readAllFlats")
	defer cleanup()

	s := flats.NewService(repo)
	h := flats.NewHandler(s)
	r := mux.NewRouter()
	r.Handle("/{id}", h.ReadAll())

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/flats", nil)
			require.NoError(tt, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			message, err := ioutil.ReadAll(rec.Body)

			require.NoError(tt, err, "response cannot be empty")

			require.Equal(tt, tc.status, rec.Code)
			require.Equal(tt, tc.message, message)
		})
	}
}

func TestHandler_Read(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		status  int
		message []byte
	}{
		{
			name:    "Should return info about flat",
			id:      "/1",
			status:  http.StatusOK,
			message: []byte(`{"id":1,"street":"Lenina","house_number":"77A","room_number":33,"description":"good flat","city":{"id":1,"country":"Belarus","name":"Minsk"}}`),
		},
		{
			name:    "Shouldn't return info about flat",
			id:      "/22",
			status:  http.StatusBadRequest,
			message: []byte(`{"error":"sql: no rows in result set"}`),
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("readFlat")
	defer cleanup()

	s := flats.NewService(repo)
	h := flats.NewHandler(s)
	r := mux.NewRouter()
	r.Handle("/{id}", h.Read())

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.id, nil)
			require.NoError(tt, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			message, err := ioutil.ReadAll(rec.Body)
			log.Println(string(message))
			require.NoError(tt, err, "response cannot be empty")

			require.Equal(tt, tc.status, rec.Code)
			require.Equal(tt, tc.message, message)
		})
	}
}

func TestHandler_Update(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		body    flat.Flat
		status  int
		message []byte
	}{
		{
			name: "Should update info about flat",
			id:   "1",
			body: flat.Flat{
				Street:      "Pirogova",
				HouseNumber: "33",
				RoomNumber:  69,
				Description: "test description",
				City: flat.City{
					ID: 1,
				},
			},
			status:  http.StatusOK,
			message: []byte(""),
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("updateFlat")
	defer cleanup()

	s := flats.NewService(repo)
	h := flats.NewHandler(s)
	r := mux.NewRouter()
	r.Handle("/{id}", h.Update())

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			b, err := json.Marshal(tc.body)
			require.NoError(tt, err)

			body := ioutil.NopCloser(bytes.NewBuffer(b))

			req, err := http.NewRequest(http.MethodPut, tc.id, body)
			require.NoError(tt, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			message, err := ioutil.ReadAll(rec.Body)
			require.NoError(tt, err, "response cannot be empty")

			require.Equal(tt, tc.message, message)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		status  int
		message []byte
	}{
		{
			name:    "Should delete info about flat",
			id:      "1",
			status:  http.StatusOK,
			message: []byte(""),
		},
	}

	repo, cleanup := database.CreateTestFlatsRepository("deleteFlat")
	defer cleanup()

	s := flats.NewService(repo)
	h := flats.NewHandler(s)
	r := flats.NewRouter(h)

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, tc.id, nil)
			require.NoError(tt, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			message, err := ioutil.ReadAll(rec.Body)
			require.NoError(tt, err, "response cannot be empty")

			require.Equal(tt, tc.message, message)
		})
	}
}