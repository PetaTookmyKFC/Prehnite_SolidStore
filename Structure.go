package prinhitesolidstore

import "path/filepath"

type Store struct {
	folder string // No External editing after creation
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

	s = &Store{
		folder: location,
	}
	return s, nil
}
