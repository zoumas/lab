package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TodoItem struct {
	ID        int    `json:"-"`
	UserID    int    `json:"-"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func unmarshal() {
	url := "https://jsonplaceholder.typicode.com/todos/2"
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	todoItem := TodoItem{}
	err = json.Unmarshal(body, &todoItem)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", todoItem)

	data, err := json.MarshalIndent(todoItem, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func decode() {
	url := "https://jsonplaceholder.typicode.com/todos/2"
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	todoItem := TodoItem{}

	decoder := json.NewDecoder(resp.Body)
	// decoder.DisallowUnknownFields()

	if err := decoder.Decode(&todoItem); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", todoItem)
}

func main() {
	todoItem := &TodoItem{
		ID:        1,
		UserID:    1,
		Title:     "",
		Completed: false,
	}

	data, err := json.MarshalIndent(todoItem, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
