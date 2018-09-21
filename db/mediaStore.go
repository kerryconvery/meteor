package db

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
)

//MediaRecord represents the media data in storage
type MediaRecord struct {
	Position int
	Duration int
}

//MediaStore represents persistent storage for media records
type MediaStore struct {
	Store
}

//NewMediaStore returns a new instance of a MediaStore
func NewMediaStore(dbLocation string) MediaStore {
	store := MediaStore{}

	store.open(dbLocation)

	return store
}

// ReadPosition reads the position recorded against the key
func (s MediaStore) Read(key string) (MediaRecord, error) {
	mediaRecord := MediaRecord{}

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}

		return s.readObject(item, &mediaRecord)
	})

	return mediaRecord, err
}

// UpdatePosition sets the current position against the key
func (s MediaStore) UpdatePosition(key string, position int) error {
	data, _ := s.Read(key)

	data.Position = position

	return s.writeObject(key, data)
}

// UpdateDuration sets the current duration against the key
func (s MediaStore) UpdateDuration(key string, duration int) error {
	data, _ := s.Read(key)

	data.Duration = duration

	return s.writeObject(key, data)
}

func (s MediaStore) readObject(item *badger.Item, out interface{}) error {
	bytes, err := item.Value()

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, out)
}
