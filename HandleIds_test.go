package prinhitesolidstore

import "testing"

func TestGenerate(t *testing.T) {
	amount := 5

	store, err := CreateStore("./")
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
