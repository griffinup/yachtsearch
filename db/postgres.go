package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/griffinup/yachtsearch/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertYacht(ctx context.Context, yacht schema.Yacht) error {
	_, err := r.db.Exec("INSERT INTO yachts(id, name, company) VALUES($1, $2, $3)", yacht.ID, yacht.Name, yacht.Company)
	return err
}

func (r *PostgresRepository) ListYachts(ctx context.Context, skip uint64, take uint64) ([]schema.Yacht, error) {
	rows, err := r.db.Query("SELECT * FROM yachts ORDER BY name DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Yachts
	yachts := []schema.Yacht{}
	for rows.Next() {
		yacht := schema.Yacht{}
		if err = rows.Scan(&yacht.ID, &yacht.Name, &meow.Company); err == nil {
			yacht = append(yachts, yacht)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return yachts, nil
}
