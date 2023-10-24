package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	directory := Directory{"contact": "address", "contact2": "address2", "contact3": "address3"}
	t.Run("Normal case", func(t *testing.T) {

		got, _ := directory.Search("contact")
		want := "address"

		checkString(t, got, want)
	})
	t.Run("Not in the directory", func(t *testing.T) {

		_, err := directory.Search("contact6")
		want := "not in the directory, please check your input"

		if err == nil {
			t.Fatal("expected an error")
		}

		checkString(t, err.Error(), want)
	})
	t.Run("Add new contact", func(t *testing.T) {
		name := "NewContact"
		address := "NewAddress"
		directory.AddContact(name, address)

		got, _ := directory.Search(name)
		want := address

		checkString(t, got, want)
	})
	t.Run("Add existing contact", func(t *testing.T) {
		name := "contact2"
		address := "address2"
		err := directory.AddContact(name, address)

		want := "the contact already exists"

		if err == nil {
			t.Fatal("error expected but got nothing")
		}

		checkString(t, err.Error(), want)
	})
	t.Run("Update existing contact", func(t *testing.T) {
		name := "contact2"
		address := "address2New"
		directory.UpdateContact(name, address)

		got, _ := directory.Search(name)
		want := address

		checkString(t, got, want)
	})
	t.Run("Delete existing contact", func(t *testing.T) {
		name := "contact2"
		directory.DeleteContact(name)

		got, _ := directory.Search(name)
		want := ""

		checkString(t, got, want)
	})
	t.Run("Delete contact that does not exist", func(t *testing.T) {
		name := "contact2"
		err := directory.DeleteContact(name)
		want := ErrContactNotThere.Error()

		checkString(t, err.Error(), want)
	})
}

func checkString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
