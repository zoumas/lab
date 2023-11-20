package main

import (
	"encoding/json"
	"os"
	"sync"
)

type DB struct {
	path string
	mu   *sync.Mutex

	Chirps    map[int]Chirp `json:"chirps"`
	Users     map[int]User  `json:"users"`
	Structure map[string]any
}

var databaseStructure = `
{
  "chirps": {
    "1": {
      "id": 1,
      "body": "This is the first chirp ever!"
    },
    "2": {
      "id": 2,
      "body": "Hello, world!"
    }
  }
}
`

func NewDB(path string) (*DB, error) {
	_, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return &DB{
		path:      path,
		mu:        &sync.Mutex{},
		Chirps:    map[int]Chirp{},
		Users:     map[int]User{},
		Structure: map[string]any{},
	}, nil
}

// Persist writes the changes to the Chrips map into disk
func (db *DB) persist() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.Structure["chirps"] = db.Chirps
	db.Structure["users"] = db.Users

	data, err := json.MarshalIndent(db.Structure, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(db.path, data, 0644)
}

func (db *DB) CreateChirp(body string) (Chirp, error) {
	id := len(db.Chirps) + 1
	chirp := Chirp{
		Body: body,
		ID:   id,
	}

	db.Chirps[id] = chirp
	err := db.persist()

	return chirp, err
}

func (db *DB) GetChirps() []Chirp {
	chirps := make([]Chirp, 0, len(db.Chirps))
	for _, chrip := range db.Chirps {
		chirps = append(chirps, chrip)
	}
	return chirps
}

func (db *DB) CreateUser(email string) (User, error) {
	id := len(db.Users) + 1
	user := User{
		Email: email,
		ID:    id,
	}

	db.Users[id] = user
	err := db.persist()
	return user, err
}
