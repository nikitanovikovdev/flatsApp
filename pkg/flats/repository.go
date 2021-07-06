package flats

import (
	"context"
	"database/sql"
	"flatApp/pkg/platform/flat"
	"fmt"
)

type Repository struct {
	db *sql.DB
	RepositoryMethods
}

type RepositoryMethods interface {
	Create()
	Read()
	Update()
	Delete()
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, f *flat.Flat) (string, error) {
	createQuery := "INSERT INTO flats (street,house_number,room_number,description,city_id) VALUES ($1,$2,$3,$4,$5) RETURNING street"

	var street string

	stmt, err := r.db.PrepareContext(ctx, createQuery)
	if err != nil {
		street = "failed to create"
		return street, err
	}

	if err := stmt.QueryRowContext(ctx, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID).Scan(&street); err != nil {
		fmt.Println(err.Error())
	}

	return street, nil
}

func (r *Repository) Read(ctx context.Context, id string) (flat.Flat, error) {
	readQuery := "SELECT flats.id, flats.street, flats.house_number, flats.room_number, " +
		"flats.description, cities.id, cities.country_name, cities.city_name " +
		"FROM flats LEFT JOIN cities ON flats.city_id = cities.id WHERE flats.id = $1"

	stmt, err := r.db.PrepareContext(ctx, readQuery)
	if err != nil {
		return flat.Flat{}, err
	}
	var f flat.Flat

	if err := stmt.QueryRowContext(ctx, id).Scan(
		&f.ID,
		&f.Street,
		&f.HouseNumber,
		&f.RoomNumber,
		&f.Description,
		&f.City.ID,
		&f.City.Country,
		&f.City.Name); err != nil {
		return flat.Flat{}, nil
	}
	return f, nil
}

func (r *Repository) Update(ctx context.Context, id string, f *flat.Flat) error {
	updateQuery := "UPDATE flats SET street = $2, house_number = $3, room_number = $4, description = $5, city_id = $6  WHERE id =$1"

	stmt, err := r.db.PrepareContext(ctx, updateQuery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, id, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	deleteQuery := "DELETE FROM flats WHERE id=$1"

	stmt, err := r.db.PrepareContext(ctx, deleteQuery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
