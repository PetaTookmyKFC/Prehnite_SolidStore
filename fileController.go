package prinhitesolidstore

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (s *Store) CheckID(id string) (bool, error) {

	// Check the folder for the id
	fullpath := s.Fullpath(id)

	_, err := os.Stat(fullpath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Store) GenerateID() (string, error) {

	found := true
	var id uuid.UUID
	var err error
	for found {
		found = false
		id, err = uuid.NewV7()

		if err != nil {
			return "", err
		}

		if val, err := s.CheckID(id.String()); err != nil {
			return "", err
			// Throw error if it occours
		} else if val {
			// The item was found - create a new one... Just backup incase of improbability
			found = true
		}

	}
	return id.String(), nil

}

// This writes a file, if the file doesn't exist a file is created.
// This should be internal use only. ( Just to ensure that keys are checked correcly )
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

func (s *Store) DeleteItem(id string) (err error) {

	// Create the full path for the item
	fullpath := filepath.Join(s.folder, id+".bin")
	// Check if that file exists
	_, err = os.Stat(fullpath)
	if os.IsNotExist(err) {
		return err
	}
	// Attempt to remove the file.
	err = os.Remove(fullpath)
	return err
}

func (s *Store) readItem(id string) (data []byte, err error) {

	fullpath := s.Fullpath(id)
	// Check that the path is correct
	if _, err = os.Stat(fullpath); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(fullpath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buff := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(buff)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buff, nil
}

func (s *Store) FindItem(check Search) ([]Result, error) {
	// Create an empty result array
	result := make([]Result, 0)

	// Read all files within the folder

	records, err := filepath.Glob(s.folder + "/*.bin")
	if err != nil {
		return nil, err
	}

	var buff []byte
	var info fs.FileInfo
	for _, recName := range records {
		// Open the file
		file, err := os.Open(recName)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		// Get the length of the file
		info, err = file.Stat()
		if err != nil {
			return nil, err
		}

		// Create the buffer for the file size, then read the file
		buff = make([]byte, info.Size())
		_, err = bufio.NewReader(file).Read(buff)
		if err != nil {
			return nil, err
		}

		// Run the check function < Check if the record is what was being searched for
		wanted, err := check(buff)
		if err != nil {
			return nil, err
		}
		// Check if the record is wanted
		if wanted {
			NR := Result{
				Value: buff,
				ID:    strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
			}
			result = append(result, NR)
		}
	}

	return result, nil
}
