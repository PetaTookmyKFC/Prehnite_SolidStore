package prinhitesolidstore

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	amount := 5

	store, err := CreateStore("./test_out")
	if err != nil {
		t.Error(err)
	}

	var id string

	for i := range amount {
		id, err = store.GenerateID()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v :- %v", i, id)

	}

}
func Test_WriteFile(t *testing.T) {

	// write file to
	amount := 5
	store, err := CreateStore("./test_out")
	if err != nil {
		t.Error(err)
	}
	var id string
	var out string
	var ch bool

	for i := range amount {
		id, err = store.GenerateID()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v :- %v", i, id)

		store.writeFile(id, []byte(id))

		// Check if ID EXISTS
		ch, err = store.CheckID(id)
		if err != nil {
			t.Error(err)
		}

		if ch {
			out += fmt.Sprintf("Found : %v \n ", id)
		} else {
			out += " correct id not found \n "
		}

		ch, err = store.CheckID(id + "BAD ID")
		if err != nil {
			t.Error(err)
		}

		if ch {
			out += " found the incorrect id \n"
		} else {
			out += " bad id not found \n"
		}

		store.writeFile(id, []byte(out))

	}

}
