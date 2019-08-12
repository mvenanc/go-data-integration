package main

import (
	"fmt"
	"github.com/gorilla/mux"
	. "neo-data/config"
	. "neo-data/config/dao"
	companyrouter "neo-data/router"
	"log"
	"net/http"
)

var dao =  CompaniesDAO{}

var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/company", companyrouter.Create).Methods("POST")
	r.HandleFunc("/api/company/{id}", companyrouter.Update).Methods("PUT")
	r.HandleFunc("/api/company/{id}", companyrouter.Delete).Methods("DELETE")

	r.HandleFunc("/api/company/{id}", companyrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/company/name/{name}", companyrouter.GetByName).Methods("GET")
	r.HandleFunc("/api/company/name/{name}/zip/{zip}", companyrouter.GetByNameZip).Methods("GET")
	r.HandleFunc("/api/company", companyrouter.GetAll).Methods("GET")

	r.HandleFunc("/api/company/batch", companyrouter.BatchData).Methods("POST")

	var port= ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
