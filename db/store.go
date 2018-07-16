package db

import (
	"strconv"

	"github.com/dgraph-io/badger"
)

// Store is a structure that represents a database
type Store struct {
	db *badger.DB
}

// New returns a new instance of Store
func New(dbLocation string) Store {
	options := badger.DefaultOptions
	options.Dir = dbLocation
	options.ValueDir = dbLocation
	db, err := badger.Open(options)

	if err != nil {
		panic(err)
	}

	return Store{db}
}

// Close closes the database
func (s Store) Close() {
	s.db.Close()
}

// Delete removes the record from the database
func (s Store) Delete(key string) {

}

// UpdatePosition sets the current position against the key
func (s Store) UpdatePosition(key string, position int) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(strconv.Itoa(position)))
	})
}

func (s Store) readIntegerValue(item *badger.Item) (int, error) {
	val, err := item.Value()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(val))
}

// ReadPosition reads the position recorded against the key
func (s Store) ReadPosition(key string) (int, error) {
	position := 0

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}

		intValue, err := s.readIntegerValue(item)

		position = intValue

		return err
	})

	return position, err
}
