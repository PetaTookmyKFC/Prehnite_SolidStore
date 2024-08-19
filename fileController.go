package prinhitesolidstore

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Store) CheckID(id string) (bool, error) {

	// Check the folder for the id
	fullpath := filepath.Join(s.folder, id+".bin")

	_, err := os.Stat(fullpath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Store) GenerateID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return id.String(), nil

}

func (s *Store) writeFile(id string, data []byte) error {
	fullpath := filepath.Join(s.folder, id)

	// file, err := os.Open(fullpath)
	// if err != nil {
	// 	return err
	// }

	file, err := os.Create(fullpath + ".bin")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
