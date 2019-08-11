package dao

import (
	. "github.com/marcio/neo-data/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
)

type CompaniesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Companies"
)

func (m *CompaniesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CompaniesDAO) GetAll() ([]Company, error) {
	var Companies []Company
	err := db.C(COLLECTION).Find(bson.M{}).All(&Companies)
	return Companies, err
}

func (m *CompaniesDAO) GetByID(id string) (Company, error) {
	var company Company
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&company)
	return company, err
}

func (m *CompaniesDAO) Create(company Company) error {
	//err := db.C(COLLECTION).Insert(&company)
	var err error
	var companies []Company
	companies, err = m.GetByName(company.Name)
	if len(companies)>0 {
		company.ID = companies[0].ID
		err = m.Update(company.ID, company)
	} else {
		err = db.C(COLLECTION).Insert(&company)
	}
	return err
}


func (m *CompaniesDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *CompaniesDAO) Update(id bson.ObjectId , company Company) error {
	//err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &company)
	err := db.C(COLLECTION).UpdateId(id, &company)
	return err
}

func (m *CompaniesDAO) GetByName(name string) ([]Company, error) {
	var Companies []Company
	err := db.C(COLLECTION).Find(bson.M{"name": bson.RegEx{Pattern:"^.*"+name+".*$", Options:"i"}}).All(&Companies)
	//^.*marcio.*$
	return Companies, err
}

func (m *CompaniesDAO) GetByNameZip(name string, zip string) ([]Company, error) {
	var Companies []Company
	err := db.C(COLLECTION).Find(bson.M{"name": bson.RegEx{Pattern:"^.*"+name+".*$", Options:"i"}, "zip": zip}).All(&Companies)
	return Companies, err
}


func (m *CompaniesDAO) CreateMany(companies *[]Company) error {
	var err error
	for _, company := range *companies {
		//err := db.C(COLLECTION).Insert(&company)
		//if err !=nil{
		//	return err
		//}
		m.Create(company)
	}
	return err
}
