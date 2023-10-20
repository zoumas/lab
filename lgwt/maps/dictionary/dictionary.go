package dictionary

type Dictionary map[string]string

// Search returns the definition for the associated word in the dictionary
// or an error if the word is not found.
func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDefinition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
