package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/straight-to-the-code-service/descriptor"
)

func findHandler(db *mgo.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values, _ := url.ParseQuery(r.URL.RawQuery)
		results := find(db.C("descriptors"), values.Get("s"))
		byteResult, _ := json.Marshal(results)

		w.Header().Set("Content-Type", "application/json")
		w.Write(byteResult)
	})
}

func find(c *mgo.Collection, queryStr string) []descriptor.Descriptor {
	results := make([]descriptor.Descriptor, 0)
	query := bson.M{"tags": bson.M{"$regex": fmt.Sprint(".*", queryStr, ".*")}}
	err := c.Find(query).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func allTagHandler(db *mgo.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results := allTags(db.C("descriptors"))
		byteResult, _ := json.Marshal(results)

		w.Header().Set("Content-Type", "application/json")
		w.Write(byteResult)
	})
}

func allTags(c *mgo.Collection) []string {
	results := make([]string, 0)

	err := c.Find(nil).Select(bson.M{"tags": 1}).Distinct("tags",&results)
	if err != nil {
		panic(err)
	}
	return results
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("straightothecode")
	defer session.Close()

	http.Handle("/find/", findHandler(db))
	http.Handle("/alltags/", allTagHandler(db))
	http.ListenAndServe(":8080", nil)
}
