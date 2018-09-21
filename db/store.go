package db

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
)

// Store is a structure that represents a database
type Store struct {
	db *badger.DB
}

// New returns a new instance of Store
func (s *Store) open(dbLocation string) {
	options := badger.DefaultOptions
	options.Dir = dbLocation
	options.ValueDir = dbLocation
	db, err := badger.Open(options)

	if err != nil {
		panic(err)
	}

	s.db = db
}

// Close closes the database
func (s Store) Close() {
	s.db.Close()
}

// Delete removes the record from the database
func (s Store) Delete(key string) {

}

func (s Store) writeObject(key string, data interface{}) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), bytes)
	})
}
