package flats

import (
	"context"
	"database/sql"
	"flatApp/pkg/platform/flat"
	"log"
)

type Repository struct {
	db          *sql.DB
	createQuery string
	readQuery   string
	updateQuery string
	deleteQuery string
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:          db,
		createQuery: "INSERT INTO flats (street,house_number,room_number,description,city_id) VALUES ($1,$2,$3,$4,$5) RETURNING house_number",
		readQuery:   "SELECT flats.id,flats.street,flats.house_number,flats.room_number,flats.description,cities.id,cities.country_name,cities.city_name FROM flats LEFT JOIN cities ON flats.city_id=cities.id",
		updateQuery: "UPDATE flats SET street = $1, house_number = $2, room_number = $3, description = $4, city_id = $5 WHERE id =$6",
		deleteQuery: "DELETE FROM flats WHERE id=$1",
	}
}

func (r *Repository) Create(ctx context.Context, f []flat.Flat) ([]string, error) {
	var houseNumbers []string

	for _, v := range f {
		var hn string

		stmt, err := r.db.PrepareContext(ctx, r.createQuery)
		if err != nil {
			log.Fatal(err)
		}

		if err := stmt.QueryRowContext(ctx, v.Street, v.HouseNumber, v.RoomNumber, v.Description, v.City).Scan(&hn); err != nil {
			log.Fatal(err)
		}

		houseNumbers = append(houseNumbers, hn)

		log.Fatal(stmt.Close())
	}

	return houseNumbers, nil
}

func (r *Repository) Read(ctx context.Context, id string) (flat.Flat, error) {
	stmt, err := r.db.PrepareContext(ctx, r.readQuery)
	if err != nil {
		log.Fatal(err)
	}
	var f flat.Flat

	if err := stmt.QueryRowContext(ctx, id).Scan(&f.Street, &f.HouseNumber, &f.RoomNumber, &f.Description, &f.City); err != nil {
		return flat.Flat{}, nil
	}
	return f, nil
}

func (r *Repository) Update(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, r.updateQuery)
	if err != nil {
		log.Fatal(err)
	}

	var f flat.Flat

	if _, err := stmt.ExecContext(ctx, id, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, r.deleteQuery)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		log.Fatal(err)
	}

	return nil
}
