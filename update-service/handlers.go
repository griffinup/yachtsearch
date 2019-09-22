package main

import (
	"bytes"
	"encoding/xml"
	//"fmt"
	//"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/griffinup/yachtsearch/db"
	"github.com/griffinup/yachtsearch/schema"
	"github.com/griffinup/yachtsearch/util"
)

func updateDBHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		status string `json:"status"`
	}

	ctx := r.Context()

	url := "http://ws.nausys.com/CBMS-external/rest/catalogue/v6/yachts/201"

	var jsonStr = []byte(`{"username":"` + cfg.NausysUser + `", "password":"` + cfg.NausysPassword + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var apires schema.YachtAllResponse

	if err := xml.Unmarshal(body, &apires); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to parse response")
		return
	}

	for _, yacht := range apires.Yachts {
		if err := db.InsertYacht(ctx, yacht); err != nil {
			log.Println(err)
			util.ResponseError(w, http.StatusInternalServerError, "Failed to insert yacht")
			return
		}
	}

	util.ResponseOk(w, response{status: "OK"})
}
