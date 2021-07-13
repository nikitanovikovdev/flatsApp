package flats

import (
	"context"
	"database/sql"
	"flatApp/pkg/platform/flat"
	"github.com/pkg/errors"
)

type RepositorySQL struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, f flat.Flat) (flat.Flat, error)
	Read(ctx context.Context, id string) (flat.Flat, error)
	ReadAll(ctx context.Context) ([]flat.Flat, error)
	Update(ctx context.Context, id string, f flat.Flat) error
	Delete(ctx context.Context, id string) error
}

func NewRepository(db *sql.DB) *RepositorySQL {
	return &RepositorySQL{
		db: db,
	}
}

func (r *RepositorySQL) Create(ctx context.Context, f flat.Flat) (flat.Flat, error) {
	createQuery := "INSERT INTO flats (street,house_number,room_number,description,city_id) VALUES ($1,$2,$3,$4,$5) RETURNING street,house_number,room_number,description,city_id"

	var fl flat.Flat

	stmt, err := r.db.PrepareContext(ctx, createQuery)
	if err != nil {
		return fl, errors.Wrap(err, "creating stmt for Repository.Create failed")
	}

	if err := stmt.QueryRowContext(ctx, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID).Scan(&fl.Street, &fl.HouseNumber, &fl.RoomNumber, &fl.Description, &fl.City.ID); err != nil {
		return fl, errors.Wrap(err, "failed to add data from Repository.Create to structure")
	}

	return fl, nil
}

func (r *RepositorySQL) Read(ctx context.Context, id string) (flat.Flat, error) {
	readQuery := "SELECT flats.id, flats.street, flats.house_number, flats.room_number, " +
		"flats.description, cities.id, cities.country_name, cities.city_name " +
		"FROM flats LEFT JOIN cities ON flats.city_id = cities.id WHERE flats.id = $1"

	stmt, err := r.db.PrepareContext(ctx, readQuery)
	if err != nil {
		return flat.Flat{}, errors.Wrap(err, "creating stmt for Repository.Read failed")
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
		return f, errors.Wrap(err, "user with this id was not found")
	}
	return f, nil
}

func (r *RepositorySQL) ReadAll(ctx context.Context) ([]flat.Flat, error) {
	readAllFlatQuery := "SELECT flats.id,flats.street,flats.house_number,flats.room_number," +
		"flats.description,cities.id,cities.country_name,cities.city_name" +
		" FROM flats LEFT JOIN cities ON flats.city_id=cities.id"

	var flats []flat.Flat
	var f flat.Flat

	rows, err := r.db.QueryContext(ctx, readAllFlatQuery)
	if err != nil {
		return []flat.Flat{}, errors.Wrap(err, "creating rows for Repository.ReadAll failed")
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
			return flats, errors.Wrap(err, "user with this id was not found")
		}

		flats = append(flats, f)
	}

	return flats, nil
}

func (r *RepositorySQL) Update(ctx context.Context, id string, f flat.Flat) error {
	updateQuery := "UPDATE flats SET street = $2, house_number = $3, room_number = $4, description = $5, city_id = $6  WHERE id =$1"

	stmt, err := r.db.PrepareContext(ctx, updateQuery)
	if err != nil {
		return errors.Wrap(err, "creating stmt for Repository.Update failed")
	}

	if _, err := stmt.ExecContext(ctx, id, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.ID); err != nil {
		return errors.Wrap(err, "failed to update data from Repository.Update")
	}

	return nil
}

func (r *RepositorySQL) Delete(ctx context.Context, id string) error {
	deleteQuery := "DELETE FROM flats WHERE id=$1"

	stmt, err := r.db.PrepareContext(ctx, deleteQuery)
	if err != nil {
		return errors.Wrap(err, "creating stmt for Repository.Delete failed")
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete data from Repository.Delete")
	}

	return nil
}
