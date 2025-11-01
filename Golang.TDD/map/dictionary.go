package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform opearation on word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Map syntax: map[key-type]value-type
// Getting a value out of a Map is the same as getting a value out of Array `map[key]`
func (d Dictionary) Search(word string) (string, error) {
	value, err := d[word]
	if !err {
		return "", ErrNotFound
	}
	return value, nil
}

// d[word] = definition => adding value to map
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
