package router

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	. "neo-data/config/dao"
	. "neo-data/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var dao = CompaniesDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	companies, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, companies)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	company, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}
	respondWithJson(w, http.StatusOK, company)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var company Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if company.ID == "" {
		company.ID = bson.NewObjectId()
	}
	company.Name = strings.ToUpper(company.Name)

	zip_, _ := strconv.Atoi(company.Zip)

	company.Zip = fmt.Sprintf("%05d",zip_)

	if err := dao.Create(company); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, company)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var company Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	company.Name = strings.ToUpper(company.Name)
	zip_, _ := strconv.Atoi(company.Zip)

	company.Zip = fmt.Sprintf("%05d",zip_)

	if err := dao.Update(bson.ObjectIdHex(params["id"]), company); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": company.Name + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}


func GetByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companies, err := dao.GetByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}
	respondWithJson(w, http.StatusOK, companies)
}

func GetByNameZip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companies, err := dao.GetByNameZip(params["name"], params["zip"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}
	respondWithJson(w, http.StatusOK, companies)
}


func BatchData(w http.ResponseWriter, r *http.Request) {

	file, err := os.Create("filename")
	_, err = io.Copy(file, r.Body)
	if err!=nil{
		respondWithError(w, http.StatusBadRequest, "Error while Posting data")
		return
	}

	f, err := os.Open("filename")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';' //delimiter
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()

	companies := []Company{}
	for _, row := range records {

		zip_, _ := strconv.Atoi(row[1])

		c := Company{ID: bson.NewObjectId(), Name: strings.ToUpper(row[0]), Zip: fmt.Sprintf("%05d",zip_), Website: ""}
		if  len(row) > 2 {
			c.Website = row[2]
		}

		companies=append(companies, c)

	}

	if err := dao.CreateMany(&companies); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
