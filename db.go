package main

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	)

// My databse struct
type Db struct {
	ip string
	session *mgo.Session
}

// Connect the provided ip
func (db *Db) Connect(ip string) {
	session, err := mgo.Dial(ip)
	if err != nil {
		panic(err)
	}
	db.session, db.ip = session, ip
}

// Return the whole collection from database
func GetCollection(database, collection string) *mgo.Collection {
	session := db.session.Copy()
	return session.DB(database).C(collection)
	// TODO : Check error
}
