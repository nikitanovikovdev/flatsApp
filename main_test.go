package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getFlat(t *testing.T) {
	var err error

	url := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"123456",
		"localhost",
		"5432",
		"postgres")

	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	getFlat(rr, req)
}

func Test_createFlat(t *testing.T) {

}

func Test_getFlats(t *testing.T) {
	var err error

	url := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"123456",
		"localhost",
		"5432",
		"postgres")

	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	getFlats(rr, req)
}
