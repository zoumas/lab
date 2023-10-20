package dictionary_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/zoumas/lab/lgwt/maps/dictionary"
)

func TestDictionary_Search(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dict := dictionary.Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dict, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown word")
		assertError(t, err, dictionary.ErrNotFound)
	})
}

func ExampleDictionary_Search() {
	dict := dictionary.Dictionary{"woman": "an adult human female"}

	definition, err := dict.Search("woman")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(definition)
	// Output: an adult human female
}

func TestDictionary_Add(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := dictionary.Dictionary{}
		word := "test"
		definition := "this is a test"

		dict.Add(word, definition)

		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := dictionary.Dictionary{word: definition}

		err := dict.Add(word, "new definition")

		assertError(t, err, dictionary.ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func ExampleDictionary_Add() {
	dict := dictionary.Dictionary{}

	dict.Add("woman", "an adult human female")
	definition, err := dict.Search("woman")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(definition)
	// Output: an adult human female
}

func TestDictionary_Update(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := dictionary.Dictionary{word: definition}

		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)
		assertNoError(t, err)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dict := dictionary.Dictionary{}
		word := "test"
		definition := "this is a test"

		err := dict.Update(word, definition)
		assertError(t, err, dictionary.ErrWordDoesNotExist)
	})
}

func ExampleDictionary_Update() {
	dict := dictionary.Dictionary{
		"woman": "an adult who lives and identifies as female though they may have been said to have a different sex at birth",
	}
	definition, err := dict.Search("woman")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(definition)
	// Output: an adult who lives and identifies as female though they may have been said to have a different sex at birth
}

func TestDictionary_Delete(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dict := dictionary.Dictionary{word: definition}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != dictionary.ErrNotFound {
		t.Errorf("\nExpected %q to be deleted", word)
	}
}

func ExampleDictionary_Delete() {
	dict := dictionary.Dictionary{"man": "an adult human male"}

	dict.Delete("man")
	fmt.Println("Dictionary is empty:", len(dict) == 0)
	// Output: Dictionary is empty: true
}

// assertDefinition should be used only when you know that the Search won't fail.
func assertDefinition(t testing.TB,
	dict dictionary.Dictionary,
	word,
	wantDefinition string,
) {
	t.Helper()

	got, err := dict.Search(word)
	assertNoError(t, err)

	if got != wantDefinition {
		t.Errorf("\ngot:\n%q\nwant:\n%q\nwhen:\nsearching for %q in %#v",
			got, wantDefinition, word, dict)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("\nExpected to get an error")
	}

	if got != want {
		t.Errorf("\ngot error:\n%q\nwant error:\n%q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("\nUnexpected error:\n%q", err)
	}
}
