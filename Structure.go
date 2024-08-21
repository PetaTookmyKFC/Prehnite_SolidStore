package SolidStore

import (
	"errors"
	"os"
	"path/filepath"
)

type Store struct {
	folder        string // No External editing after creation
	routineNumber int    // The number of goroutines for the searching
}

// Result Construct
type Result struct {
	ID    string
	Value []byte
}

// Search Function
type Search func([]byte) (Wanted bool, err error)

// Small function to make fullpath
func (s *Store) Fullpath(id string) (fullpath string) {
	return filepath.Join(s.folder, id+".bin")
}

func (s *Store) Folder() string {
	return s.folder
}

// Main Function
func CreateStore(location string) (s *Store, err error) {
	// Check if the path is absolute
	if filepath.IsLocal(location) {
		// Convert the path to absolute
		location, err = filepath.Abs(location)
		if err != nil {
			return nil, err
		}
	}

	st, err := os.Stat(location)
	if os.IsNotExist(err) || !st.IsDir() {
		os.MkdirAll(location, os.ModePerm)
	}

	s = &Store{
		folder:        location,
		routineNumber: 3,
	}
	return s, nil
}

// Writes the data to a new file! <- Doesn't update the item
func (s *Store) CreateNew(data []byte) (id string, err error) {
	id, err = s.GenerateID()
	if err != nil {
		return "", err
	}
	// Will return with id even if writing failed!
	err = s.writeFile(id, data)
	return id, err
}

// Updates an item... This is used to prevent a record from being created.
func (s *Store) UpdateItem(id string, data []byte) (written bool, err error) {
	// Check that the item exists or return error
	if v, err := s.CheckID(id); err != nil {
		return false, err
	} else if !v {
		return false, errors.New("item not found " + id)
	}
	// Update the item
	err = s.writeFile(id, data)
	if err != nil {
		return false, err
	}
	return true, nil
}

// // Delete an item
// func (s *Store) DeleteItem(id string) (err error) {
// 	// Check that the item exists
// 	if v, err := s.CheckID(id); err != nil {
// 		return err
// 	} else if !v {
// 		return errors.New("item not found " + id)
// 	}

// 	return nil
// }
