package maps

type Directory map[string]string
type DirectoryErr string

func (e DirectoryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound        = DirectoryErr("not in the directory, please check your input")
	ErrContactExists   = DirectoryErr("the contact already exists")
	ErrContactNotThere = DirectoryErr("the contact dose not exists")
)

func (d Directory) Search(word string) (string, error) {
	if d[word] == "" {
		return "", ErrNotFound
	}
	return d[word], nil
}

func (d Directory) AddContact(name, address string) error {
	searchOut, _ := d.Search(name)
	if searchOut != "" {
		return ErrContactExists
	} else {
		d[name] = address
	}
	return nil
}

func (d Directory) UpdateContact(name, address string) error {
	searchOut, _ := d.Search(name)
	if searchOut == "" {
		return ErrContactNotThere
	} else {
		d[name] = address
	}
	return nil
}

func (d Directory) DeleteContact(name string) error {
	searchOut, _ := d.Search(name)
	if searchOut == "" {
		return ErrContactNotThere
	} else {
		delete(d, name)
	}
	return nil
}
