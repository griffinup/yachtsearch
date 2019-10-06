package main

import (
	//"context"
	//"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/griffinup/yachtsearch/db"
	"github.com/griffinup/yachtsearch/schema"
	//"github.com/griffinup/yachtsearch/search"
	"github.com/griffinup/yachtsearch/util"
)

func liveSearchHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	//чтение параметров
	vars := mux.Vars(r)
	query := vars["query"]

	//на случай реализации на фронтенде постраничного вывода
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	results, err := db.LiveSearch(ctx, query, skip, take)
	if err != nil {
		util.ResponseOk(w, err.Error())
		return
	}

	util.ResponseOk(w, results)
}

func infoYachtsByModelHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	vars := mux.Vars(r)

	model, err := strconv.Atoi(vars["id"])
	if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Missing model id parameter")
			return
	}

	results, err := db.InfoByModel(ctx, model)
	if err != nil {
		util.ResponseOk(w, err.Error())
		return
	}

	util.ResponseOk(w, results)
}

func infoYachtsByBuilderHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	vars := mux.Vars(r)

	builder, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "Missing builder id parameter")
		return
	}

	results, err := db.InfoByBuilder(ctx, builder)
	if err != nil {
		util.ResponseOk(w, err.Error())
		return
	}

	util.ResponseOk(w, results)
}

func infoYachtsByNameHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	vars := mux.Vars(r)

	inforesults := []schema.InfoResult{}

	results, err := db.LiveSearch(ctx, vars["query"], 0, 100)
	if err != nil {
		util.ResponseOk(w, err.Error())
		return
	}

	for _, result := range results {
		if result.Type == "model" {
			modelresults, err := db.InfoByModel(ctx, result.ID)
			if err != nil {
				util.ResponseOk(w, err.Error())
				return
			}
			inforesults = append(inforesults, modelresults...)
		} else if result.Type == "builder" {
			builderresults, err := db.InfoByBuilder(ctx, result.ID)
			if err != nil {
				util.ResponseOk(w, err.Error())
				return
			}
			inforesults = append(inforesults, builderresults...)
		}
	}

	util.ResponseOk(w, inforesults)
}