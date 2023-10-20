package dictionary

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("could not add definition because word already exists")
	ErrWordDoesNotExist = DictionaryErr("could not update definition because word does not exist")
)
