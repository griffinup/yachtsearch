package search

import (
	"context"

	"github.com/griffinup/yachtsearch/schema"
)

type Repository interface {
	Close()
	InsertYacht(ctx context.Context, yacht schema.Yacht) error
	SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Yacht, error)
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

func SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Yacht, error) {
	return impl.SearchYacht(ctx, query)
}
