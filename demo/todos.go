package demo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Todos struct {
	filename string
}

// NewTodos attempts to create a new Todo list
func NewTodos() (*Todos, error) {
	// Create new Todos instance
	result := &Todos{}
	// Try and get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Join the cwd with our todos filename
	filename := path.Join(cwd, "myList.json")
	// Set the filename member of our new Todo list
	result.filename = filename
	// Return it
	return result, nil
}

func (t *Todos) LoadList() (string, error) {
	bytes, err := ioutil.ReadFile(t.filename)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s ", t.filename)
	}
	return string(bytes), err
}

func (t *Todos) SaveList(todos string) error {
	return ioutil.WriteFile(t.filename, []byte(todos), 0600)
}
