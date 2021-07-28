package flats

import (
	"context"
	"database/sql"
	"fmt"
	flat2 "github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/flat"
)

type RepositorySQL struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, f flat2.FlatWithUsername) (flat2.Flat, error)
	Read(ctx context.Context, u flat2.Username) ([]flat2.Flat, error)
	ReadAll(ctx context.Context) ([]flat2.Flat, error)
	Update(ctx context.Context, id string, f flat2.FlatWithUsername) error
	Delete(ctx context.Context, id string, usr flat2.Username) error
}

func NewRepository(db *sql.DB) *RepositorySQL {
	return &RepositorySQL{
		db: db,
	}
}

func (r *RepositorySQL) Create(ctx context.Context, f flat2.FlatWithUsername) (flat2.Flat, error) {
	createQuery := "INSERT INTO flats (street,house_number,room_number,description,city_id, username) VALUES ($1,$2,$3,$4,$5,$6) RETURNING street,house_number,room_number,description,city_id"

	var fl flat2.Flat

	stmt, err := r.db.PrepareContext(ctx, createQuery)
	if err != nil {
		return fl, err
	}

	if err := stmt.QueryRowContext(ctx, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID, f.Username.Username).Scan(&fl.Street, &fl.HouseNumber, &fl.RoomNumber, &fl.Description, &fl.City.ID); err != nil {
		return fl, err
	}

	return fl, nil
}

func (r *RepositorySQL) Read(ctx context.Context, u flat2.Username) ([]flat2.Flat, error) {
	readQuery := "SELECT flats.id, flats.street, flats.house_number, flats.room_number, " +
		"flats.description, cities.id, cities.country_name, cities.city_name " +
		"FROM flats LEFT JOIN cities ON flats.city_id = cities.id WHERE username= $1"

	var flats []flat2.Flat
	var f flat2.Flat

	stmt, err := r.db.PrepareContext(ctx, readQuery)
	if err != nil {
		return flats, err
	}

	rows, err := stmt.QueryContext(ctx, u.Username)

	for rows.Next() {
		err := rows.Scan(
			&f.ID,
			&f.Street,
			&f.HouseNumber,
			&f.RoomNumber,
			&f.Description,
			&f.City.ID,
			&f.City.Country,
			&f.City.Name)
		if err != nil {
			return flats, err
		}
		flats = append(flats, f)
	}
	return flats, nil
}

func (r *RepositorySQL) ReadAll(ctx context.Context) ([]flat2.Flat, error) {
	readAllFlatQuery := "SELECT flats.id,flats.street,flats.house_number,flats.room_number," +
		"flats.description,cities.id,cities.country_name,cities.city_name" +
		" FROM flats LEFT JOIN cities ON flats.city_id=cities.id"

	var flats []flat2.Flat
	var f flat2.Flat

	rows, err := r.db.QueryContext(ctx, readAllFlatQuery)

	if err != nil {
		return []flat2.Flat{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&f.ID,
			&f.Street,
			&f.HouseNumber,
			&f.RoomNumber,
			&f.Description,
			&f.City.ID,
			&f.City.Country,
			&f.City.Name)
		if err != nil {
			return flats, err
		}

		flats = append(flats, f)
	}

	return flats, nil
}

func (r *RepositorySQL) Update(ctx context.Context, id string, f flat2.FlatWithUsername) error {
	updateQuery := "UPDATE flats SET street = $2, house_number = $3, room_number = $4, description = $5, city_id = $6  WHERE id =$1 AND username=$7"

	stmt, err := r.db.PrepareContext(ctx, updateQuery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, id, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID, f.Username.Username); err != nil {
		return err
	}

	return nil
}

func (r *RepositorySQL) Delete(ctx context.Context, id string, usr flat2.Username) error {
	deleteQuery := "DELETE FROM flats WHERE id = $1 AND username = $2"

	stmt, err := r.db.PrepareContext(ctx, deleteQuery)

	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, id, usr.Username); err != nil {
		return fmt.Errorf("wasn't found id or username")
	}

	return nil
}