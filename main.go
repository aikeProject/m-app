package main

import (
	"fmt"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"io/ioutil"
	"os"
	"path"
)

func basic() string {
	return "Hello World!"
}

func saveList(todos string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := path.Join(cwd, "myList.json")
	return ioutil.WriteFile(filename, []byte(todos), 0600)
}

func loadList() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filename := path.Join(cwd, "my.json")
	result, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s ", filename)
	}
	return string(result), err
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "magick-app",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(saveList)
	app.Bind(loadList)
	app.Run()
}
