package main

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	)

// My databse struct
type Db struct {
	ip string
	Session *mgo.Session
}

// Connect the provided ip
func (db *Db) Connect(ip string) {
	session, err := mgo.Dial(ip)
	if err != nil {
		panic(err)
	}
	db.Session, db.ip = session, ip
}

/*// Return the whole collection from database
func (d *Db) GetCollection(database, collection string) *mgo.Collection {
	session := d.session.Copy()
	return session.DB(database).C(collection)
	// TODO : Check error
}*/
