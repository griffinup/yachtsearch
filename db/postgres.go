package db

import (
	"context"
	"database/sql"
	"strings"
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
	_, err := r.db.Exec("INSERT INTO yachts(id, name, company, model) VALUES($1, $2, $3, $4) ON CONFLICT (id) DO UPDATE SET name = $2, company = $3, model = $4", yacht.ID, yacht.Name, yacht.Company, yacht.Model)
	return err
}

func (r *PostgresRepository) InsertCompany(ctx context.Context, company schema.Company) error {
	_, err := r.db.Exec("INSERT INTO companies(id, name) VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET name=$2", company.ID, company.Name)
	return err
}

func (r *PostgresRepository) InsertModel(ctx context.Context, model schema.Model) error {
	_, err := r.db.Exec("INSERT INTO models(id, name) VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET name=$2", model.ID, model.Name)
	return err
}

func (r *PostgresRepository) SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.YachtFull, error) {
	rows, err := r.db.Query("SELECT * FROM (SELECT yachts.id, yachts.name, companies.name AS cname, models.name AS mname FROM yachts LEFT OUTER JOIN companies ON yachts.company = companies.id LEFT OUTER JOIN models ON yachts.model = models.id) AS t1 WHERE LOWER(name) LIKE '" + strings.ToLower(query) + "%' OR LOWER(mname) LIKE '" + strings.ToLower(query) + "%' ORDER BY name DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Yachts
	yachts := []schema.YachtFull{}
	for rows.Next() {
		yacht := schema.YachtFull{}
		if err = rows.Scan(&yacht.ID, &yacht.Name, &yacht.CompanyName, &yacht.ModelName); err == nil {
			yachts = append(yachts, yacht)
		} else {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return yachts, nil
}
