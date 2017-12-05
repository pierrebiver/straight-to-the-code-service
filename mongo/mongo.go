package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/straight-to-the-code-service/model"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

var DBName = "straight-to-the-code"

func GetSession() (*mgo.Session, error) {
	if session == nil {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return session, err
	}

	newSession := session.Copy()

	return newSession, nil
}

func Descriptors() ([]model.Descriptor, error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	descriptors := make([]model.Descriptor, 0)
	err := db.C("descriptors").Find(bson.M{}).All(&descriptors)
	logError(err)

	return descriptors, err
}

func Add(d *model.DescriptorInput) (error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	err := db.C("descriptors").Insert(&d)
	logError(err)

	return err
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
