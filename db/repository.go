package db

import (
	"context"

	"github.com/griffinup/yachtsearch/schema"
)

type Repository interface {
	Close()
	InsertYacht(ctx context.Context, yacht schema.Yacht) error
	InsertCompany(ctx context.Context, company schema.Company) error
	InsertModel(ctx context.Context, model schema.Model) error
	SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.YachtFull, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertYacht(ctx context.Context, yacht schema.Yacht) error {
	return impl.InsertYacht(ctx, yacht)
}

func InsertCompany(ctx context.Context, company schema.Company) error {
	return impl.InsertCompany(ctx, company)
}

func InsertModel(ctx context.Context, model schema.Model) error {
	return impl.InsertModel(ctx, model)
}

func SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.YachtFull, error) {
	return impl.SearchYachts(ctx, query, skip, take)
}
