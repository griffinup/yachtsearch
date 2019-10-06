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
	InsertBuilder(ctx context.Context, builder schema.Builder) error
	LiveSearch(ctx context.Context, query string, skip uint64, take uint64) ([]schema.LiveResult, error)
	InfoByModel(ctx context.Context, model int) ([]schema.InfoResult, error)
	InfoByBuilder(ctx context.Context, builder int) ([]schema.InfoResult, error)
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

func InsertBuilder(ctx context.Context, builder schema.Builder) error {
	return impl.InsertBuilder(ctx, builder)
}


func LiveSearch(ctx context.Context, query string, skip uint64, take uint64) ([]schema.LiveResult, error) {
	return impl.LiveSearch(ctx, query, skip, take)
}

func InfoByModel(ctx context.Context, model int) ([]schema.InfoResult, error) {
	return impl.InfoByModel(ctx, model)
}

func InfoByBuilder(ctx context.Context, builder int) ([]schema.InfoResult, error) {
	return impl.InfoByBuilder(ctx, builder)
}
