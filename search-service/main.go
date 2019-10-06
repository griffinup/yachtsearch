package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/griffinup/yachtsearch/db"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	//Live search
	router.HandleFunc("/search/{query}", liveSearchHandler).
		Methods("GET")
	//Get yacht list by model
	router.HandleFunc("/info/model/{id:[0-9]+}", infoYachtsByModelHandler).
		Methods("GET")
	//Get yacht list by builder
	router.HandleFunc("/info/builder/{id:[0-9]+}", infoYachtsByBuilderHandler).
		Methods("GET")
	//Get yacht list by keyword (models + builders)
	router.HandleFunc("/info/name/{query}", infoYachtsByNameHandler).
		Methods("GET")
	return
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	// Run HTTP server
	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
