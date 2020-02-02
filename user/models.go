package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//User represents a sample database entity.
type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstname"`
	LastName string    `json:"lastname"`
	EmailID string    `json:"email"`
}

var db *mgo.Database

// Connection to MoongoDB
func init() {
	session, err := mgo.Dial("localhost/api_db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("api_db")
}

// Get Database Object With Collection Name
func collection() *mgo.Collection {
	return db.C("api")
}

// GetAll returns all users from the database.
func GetAll() ([]User, error) {
	res := []User{}
	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetOne returns a single user from the database.
func GetOne(id string) (*User, error) {
	res := User{}
	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Save inserts an user to the database.
func Save(user User) error {
	return collection().Insert(user)
}

// Remove deletes an user from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}

// Update user an user from the database
func Update(user User) error {
	_, err := collection().UpsertId(user.ID, user)
	if err != nil {
		return err
	}
	return nil
}