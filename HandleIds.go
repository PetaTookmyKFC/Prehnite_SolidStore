package prinhitesolidstore

import "github.com/google/uuid"

func (s *Store) CheckID() {

}

func (s *Store) GenerateID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return id.String(), nil

}
