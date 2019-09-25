package search

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/griffinup/yachtsearch/schema"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertYacht(ctx context.Context, yacht schema.Yacht) error {
	_, err := r.client.Index().
		Index("yachts").
		Type("yacht").
		Id(strconv.Itoa(yacht.ID)).
		BodyJson(yacht).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchYachts(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Yacht, error) {
	result, err := r.client.Search().
		Index("yachts").
		Query(
			elastic.NewMultiMatchQuery(query, "name").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	yachts := []schema.Yacht{}
	for _, hit := range result.Hits.Hits {
		var yacht schema.Yacht
		if err = json.Unmarshal(*hit.Source, &yacht); err != nil {
			log.Println(err)
		}
		yachts = append(yachts, yacht)
	}
	return yachts, nil
}
