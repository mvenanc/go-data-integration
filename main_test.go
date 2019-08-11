package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/marcio/neo-data/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"

	companyrouter "github.com/marcio/neo-data/router"
	"testing"
	//companyrouter "github.com/marcio/neo-data/router"
)


func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/company", companyrouter.Create).Methods("POST")
	r.HandleFunc("/api/company/{id}", companyrouter.Update).Methods("PUT")
	r.HandleFunc("/api/company/{id}", companyrouter.Delete).Methods("DELETE")

	r.HandleFunc("/api/company/{id}", companyrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/company/name/{name}", companyrouter.GetByName).Methods("GET")
	r.HandleFunc("/api/company/name/{name}/zip/{zip}", companyrouter.GetByNameZip).Methods("GET")
	r.HandleFunc("/api/company", companyrouter.GetAll).Methods("GET")

	r.HandleFunc("/api/company/batch", companyrouter.BatchData).Methods("POST")
	return r
}


func TestCreate_01(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d110", "name": "TESTCREATE_01", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"id":"5d4f2986cad99073f5a8d110","name":"TESTCREATE_01","zip":"12345","website":"company@teste.com"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreate_02(t *testing.T) {
	var jsonStr = []byte(`{"name": "TestCreate_02", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	//assert.Equal(t,200,rr.Code, nil)
	//expected := `{"id":"5d4f2986cad99073f5a8d444","name":"TestGetById_01","zip":"12345","website":"company@teste.com"}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}

func TestCreate_03(t *testing.T) {
	var jsonStr = []byte(`{"name": "TestCreate_03", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"name":"TESTCREATE_03","zip":"12345","website":"company@teste.com"}`

	var company Company

	json.NewDecoder(rr.Body).Decode(&company)
	if !(company.Name == "TESTCREATE_03" &&  company.Zip == "12345" && company.Website == "company@teste.com"){
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestUpdate_01(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d111", "name": "TESTUPDATE_01", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d111", "name": "TESTUPDATE_01", "zip" : "99999" , "website": "company@teste.com"}`)
	req, err = http.NewRequest("PUT", "/api/company/5d4f2986cad99073f5a8d111", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"result":"TESTUPDATE_01 atualizado com sucesso!"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//5d4f2986cad99073f5a8d110  5d4f2986cad99073f5a8d111
func TestUpdate_02(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d112", "name": "TESTUPDATE_02", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("PUT", "/api/company/5d4f2986cad99073f5a8d112", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"error":"not found"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDelete_01(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d113", "name": "TESTDELETE_01", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	jsonStr = []byte(`{}`)

	req, err = http.NewRequest("DELETE", "/api/company/5d4f2986cad99073f5a8d113", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"result":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestDelete_02(t *testing.T) {
	var jsonStr = []byte(`{}`)

	req, err := http.NewRequest("DELETE", "/api/company/5d4f2986cad99073f5a8d114", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"error":"not found"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetById_01(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d115", "name": "TESTGETBYID_01", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/api/company/5d4f2986cad99073f5a8d115", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"id":"5d4f2986cad99073f5a8d115","name":"TESTGETBYID_01","zip":"12345","website":"company@teste.com"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetById_02(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/company/5d4f2986cad99073f5a8d116", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"error":"Invalid company ID"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetByName_01(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d119", "name": "TESTGETBYNAME_01", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/api/company/name/TESTGETBYNAME_01", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `[{"id":"5d4f2986cad99073f5a8d119","name":"TESTGETBYNAME_01","zip":"12345","website":"company@teste.com"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetByName_02(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/company/name/TEST", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	var companies []Company
	json.NewDecoder(rr.Body).Decode(&companies)

	if (len(companies) == 0){
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "")
	}
}

func TestGetAll_01(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/company", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	assert.Equal(t,200,rr.Code, nil)
}


func TestGetAll_02(t *testing.T) {

	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d118", "name": "TESTGETALL_02", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/api/company", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	var companies []Company
	json.NewDecoder(rr.Body).Decode(&companies)

	if (len(companies) == 0){
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "")
	}

}


//func TestEXAMPLE(t *testing.T) {
//
//	//var jsonStr = []byte(`{"name": "company 1", "zip" : "12345" , "website": "company@teste.com"}`)
//	//req, err := http.NewRequest("GET", "/api/company", bytes.NewBuffer(jsonStr))
//	req, err := http.NewRequest("GET", "/api/company", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	Router().ServeHTTP(rr, req)
//
//	//fmt.Println("Server running in port:", port)
//	//http.ListenAndServe(port, r)
//	//fmt.Println(rr.Body)
//	assert.Equal(t,200,rr.Code, nil)
//	//if status := rr.Code; status != http.StatusOK {
//	//	t.Errorf("handler returned wrong status code: got %v want %v",
//	//		status, http.StatusOK)
//	//}
//	//
//	//expected := `{"name": "company 1", "zip" : "12345" , "website": "company@teste.com"}`
//	////if rr.Body.String() != expected {
//	////	t.Errorf("handler returned unexpected body: got %v want %v",
//	////		rr.Body.String(), expected)
//	////}
//	//
//	//var company Company
//	//json.NewDecoder(rr.Body).Decode(&company)
//	//if !(company.AdressZip == "12345" && company.Name == "company 1" && company.Website == "company@teste.com"){
//	//	t.Errorf("handler returned unexpected body: got %v want %v",
//	//		rr.Body.String(), expected)
//	//}
//}

func TestDelete_All(t *testing.T) {
	var jsonStr = []byte(`{ "id":"5d4f2986cad99073f5a8d383", "name": "TESTDELETE_02", "zip" : "12345" , "website": "company@teste.com"}`)
	req, err := http.NewRequest("POST", "/api/company", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	jsonStr = []byte(`{}`)

	req, err = http.NewRequest("DELETE", "/api/company/5d4f2986cad99073f5a8d383", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	expected := `{"result":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}




