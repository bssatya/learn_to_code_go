package main

type Dictionary map[string]string

const (
	ErrorNotFound          = DictionaryErr("couldnt find the word you are looking for")
	ErrorWordExists        = DictionaryErr("can not add the word because it already exists")
	ErrorWordDoesNotExists = DictionaryErr("can not update the definition because the word does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return (string(e))
}

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
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
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	return nil
}
