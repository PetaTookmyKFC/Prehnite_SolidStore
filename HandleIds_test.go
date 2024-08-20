package prinhitesolidstore

import (
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

func Test_WriteBuffer(t *testing.T) {

	store, err := CreateStore("./test_out")
	if err != nil {
		t.Error(err)
	}

	buff := [][]byte{
		{0x11, 0x52, 0x12},
		{
			0x01, 0x02, 0x03,
		},
		{
			0x22, 0x11, 0x67, 0x82, 0x72,
		},
	}

	for _, v := range buff {

		id, err := store.CreateNew(v)
		if err != nil {
			t.Error(err)
		}

		t.Log(id)

	}

}

func Test_DeleteItem(t *testing.T) {
	store, err := CreateStore("./test_out")
	if err != nil {
		t.Error(err)
	}

	res, err := store.FindItem(func(b []byte) (Wanted bool, err error) {
		if b[0] != byte(3) {
			return false, nil
		}
		return true, nil
	})

	if err != nil {
		t.Error(err)
	}

	for _, v := range res {
		err = store.DeleteItem(v.ID)
		if err != nil {
			t.Error(err)
		}
	}

	t.Log("Done!")

}

func Test_FindItem(t *testing.T) {
	store, err := CreateStore("./test_out")
	if err != nil {
		t.Error(err)
	}

	res, err := store.FindItem(func(b []byte) (Wanted bool, err error) {

		if b[0] == byte(3) {
			return false, nil
		}

		return true, nil
	})

	if err != nil {
		t.Error(err)
	}

	for _, r := range res {

		t.Logf("%v \n", r.Value)

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
	// var out string
	// var ch bool

	for i := range amount {
		id, err = store.GenerateID()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v : - Internal Code Struct", i)

		b := byte(i)

		store.writeFile(id, []byte{b})

		// // Check if ID EXISTS
		// ch, err = store.CheckID(id)
		// if err != nil {
		// 	t.Error(err)
		// }

		// if ch {
		// 	out += fmt.Sprintf("Found : %v \n ", id)
		// } else {
		// 	out += " correct id not found \n "
		// }

		// ch, err = store.CheckID(id + "BAD ID")
		// if err != nil {
		// 	t.Error(err)
		// }

		// if ch {
		// 	out += " found the incorrect id \n"
		// } else {
		// 	out += " bad id not found \n"
		// }

		// store.writeFile(id, []byte(out))

	}

}
