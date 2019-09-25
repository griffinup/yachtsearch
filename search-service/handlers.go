package main

import (
	//"context"
	//"log"
	"net/http"
	"strconv"

	//"strconv"

	"github.com/griffinup/yachtsearch/db"
	//"github.com/griffinup/yachtsearch/schema"
	//"github.com/griffinup/yachtsearch/search"
	"github.com/griffinup/yachtsearch/util"
)

/*
func listYachtsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	// Read parameters
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

	// Fetch yachts
	yachts, err := db.ListYachts(ctx, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Could not fetch yachts")
		return
	}

	util.ResponseOk(w, yachts)
}
*/

func searchYachtsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	// Read parameters
	query := r.FormValue("query")
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

	// Search yachts
	yachts, err := db.SearchYachts(ctx, query, skip, take)
	if err != nil {
		//log.Println(err)
		//util.ResponseOk(w, []schema.Yacht{})
		util.ResponseOk(w, err.Error())
		return
	}

	util.ResponseOk(w, yachts)
}
