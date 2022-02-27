package mydict

import "errors"

/*
type에도 method를 연결할 수 있다.
*/

// Dictionary type
type Dictionary map[string]string

var ErrNotFound = errors.New("not found word")
var ErrDuplicateWord = errors.New("already exist word")

func (d *Dictionary) Search(word string) (string, error) {
	dict := *d
	value, exists := dict[word]

	if exists {
		return value, nil
	}

	return "", ErrNotFound
}

// Add a word to the dictionary
func (d *Dictionary) Add(word, description string) error {
	dict := *d
	_, err := dict.Search(word)

	if err == nil {
		return ErrDuplicateWord
	}

	dict[word] = description
	return nil
}

func (d *Dictionary) Update(word, newDescription string) error {
	dict := *d
	_, err := dict.Search(word)

	if err != nil {
		return err
	}

	dict[word] = newDescription
	return nil
}

func (d *Dictionary) Delete(word string) error {
	dict := *d
	_, err := dict.Search(word)

	if err != nil {
		return err
	}

	delete(dict, word)
	return nil
}
