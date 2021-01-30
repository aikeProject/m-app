package main

import (
	"log"
	"magick-app/demo"
	"magick-app/lib/config"
	"magick-app/lib/image"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	myTodoList, err := demo.NewTodos()
	log.Println("myTodoList", myTodoList)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    576,
		Title:     "magick-app",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	newConfig := config.NewConfig()
	fileManager := image.NewFileManager(newConfig)

	app.Bind(basic)
	app.Bind(myTodoList)
	app.Bind(newConfig)
	app.Bind(fileManager)
	_ = app.Run()
}
