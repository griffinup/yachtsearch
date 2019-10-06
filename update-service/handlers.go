package main

import (
	"encoding/json"
	"strconv"

	//"fmt"
	//"html/template"
	"log"
	"net/http"

	"github.com/griffinup/yachtsearch/db"
	"github.com/griffinup/yachtsearch/schema"
	"github.com/griffinup/yachtsearch/util"
	"github.com/kelseyhightower/envconfig"
)

func updateDBHandler(w http.ResponseWriter, r *http.Request) {

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	type response struct {
		status string `json:"status"`
	}

	ctx := r.Context()

	var body []byte

	//Get all charter companies and their yachts
	body = util.ApiRequest("http://ws.nausys.com/CBMS-external/rest/catalogue/v6/charterCompanies", `{"username":"` + cfg.NausysUser + `", "password":"` + cfg.NausysPassword + `"}`)

	var allccompanies schema.CompaniesAllResponse

	if err := json.Unmarshal(body, &allccompanies); err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to parse companies response" + err.Error())
		return
	}

	for _, company := range allccompanies.Companies {
		if err := db.InsertCompany(ctx, company); err != nil {
			util.ResponseError(w, http.StatusInternalServerError, "Failed to insert company" + err.Error())
			return
		}

		body = util.ApiRequest("http://ws.nausys.com/CBMS-external/rest/catalogue/v6/yachts/" + strconv.Itoa(company.ID), `{"username":"` + cfg.NausysUser + `", "password":"` + cfg.NausysPassword + `"}`)
		var allyachts schema.YachtAllResponse

		if err := json.Unmarshal(body, &allyachts); err != nil {
			util.ResponseError(w, http.StatusInternalServerError, "Failed to parse response" + err.Error())
			return
		}

		for _, yacht := range allyachts.Yachts {
			if err := db.InsertYacht(ctx, yacht); err != nil {
				util.ResponseError(w, http.StatusInternalServerError, "Failed to insert yacht" + err.Error())
				return
			}
		}
	}

	//Get all yacht models
	body = util.ApiRequest("http://ws.nausys.com/CBMS-external/rest/catalogue/v6/yachtModels", `{"username":"` + cfg.NausysUser + `", "password":"` + cfg.NausysPassword + `"}`)

	var allmodels schema.ModelsAllResponse

	if err := json.Unmarshal(body, &allmodels); err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to parse companies response" + err.Error())
		return
	}

	for _, model := range allmodels.Models {
		if err := db.InsertModel(ctx, model); err != nil {
			util.ResponseError(w, http.StatusInternalServerError, "Failed to insert model" + err.Error())
			return
		}
	}

	//Get all yacht builders
	body = util.ApiRequest("http://ws.nausys.com/CBMS-external/rest/catalogue/v6/yachtBuilders", `{"username":"` + cfg.NausysUser + `", "password":"` + cfg.NausysPassword + `"}`)

	var allbuilders schema.BuildersAllResponse

	if err := json.Unmarshal(body, &allbuilders); err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to parse companies response" + err.Error())
		return
	}

	for _, builder := range allbuilders.Builders {
		if err := db.InsertBuilder(ctx, builder); err != nil {
			util.ResponseError(w, http.StatusInternalServerError, "Failed to insert builder" + err.Error())
			return
		}
	}
	util.ResponseOk(w, response{status: "OK"})
}
