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
	_, err := r.db.Exec("INSERT INTO models(id, name, builder) VALUES($1, $2, $3) ON CONFLICT (id) DO UPDATE SET name=$2, builder=$3", model.ID, model.Name, model.Builder)
	return err
}

func (r *PostgresRepository) InsertBuilder(ctx context.Context, model schema.Builder) error {
	_, err := r.db.Exec("INSERT INTO builders(id, name) VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET name=$2", model.ID, model.Name)
	return err
}

func (r *PostgresRepository) LiveSearch(ctx context.Context, query string, skip uint64, take uint64) ([]schema.LiveResult, error) {
	//rows, err := r.db.Query("SELECT * FROM (SELECT yachts.id, yachts.name, companies.name AS cname, models.name AS mname FROM yachts LEFT OUTER JOIN companies ON yachts.company = companies.id LEFT OUTER JOIN models ON yachts.model = models.id) AS t1 WHERE LOWER(name) LIKE '" + strings.ToLower(query) + "%' OR LOWER(mname) LIKE '" + strings.ToLower(query) + "%' ORDER BY name DESC OFFSET $1 LIMIT $2", skip, take)
	rows, err := r.db.Query("SELECT * FROM (SELECT models.id as id, models.name as name, 'model' as type FROM models WHERE LOWER(name) LIKE '" + strings.ToLower(query) + "%' UNION SELECT builders.id as id, builders.name as name, 'builder' as type FROM builders WHERE LOWER(name) LIKE '" + strings.ToLower(query) + "%') as t1 ORDER BY name DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Results
	results := []schema.LiveResult{}
	for rows.Next() {
		result := schema.LiveResult{}
		if err = rows.Scan(&result.ID, &result.Name, &result.Type); err == nil {
			results = append(results, result)
		} else {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *PostgresRepository) InfoByModel(ctx context.Context, model int) ([]schema.InfoResult, error) {

	var modelname string
	var builderid int
	var buildername string

	err := r.db.QueryRow("SELECT name, builder FROM models WHERE id = $1", model).Scan(&modelname, &builderid)

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow("SELECT name FROM builders WHERE id = $1", builderid).Scan(&buildername)

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT yachts.id, yachts.name, companies.name FROM yachts LEFT JOIN companies ON yachts.company = companies.id WHERE model = $1", model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []schema.InfoResult{}
	for rows.Next() {
		result := schema.InfoResult{}
		if err = rows.Scan(&result.ID, &result.Name, &result.CompanyName); err == nil {
			result.ModelName = modelname
			result.BuilderName = buildername
			result.Avail = true
			results = append(results, result)
		} else {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *PostgresRepository) InfoByBuilder(ctx context.Context, builder int) ([]schema.InfoResult, error) {

	var buildername string
	err := r.db.QueryRow("SELECT name FROM builders WHERE id = $1", builder).Scan(&buildername)

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT yachts.id, yachts.name, models.name, companies.name FROM yachts INNER JOIN (SELECT * FROM models WHERE builder = $1) as models ON yachts.model = models.id LEFT JOIN companies ON yachts.company = companies.id", builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []schema.InfoResult{}
	for rows.Next() {
		result := schema.InfoResult{}
		if err = rows.Scan(&result.ID, &result.Name, &result.ModelName, &result.CompanyName); err == nil {
			result.BuilderName = buildername
			result.Avail = true
			results = append(results, result)
		} else {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}