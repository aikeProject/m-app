package main

import (
	"log"
	"magick-app/demo"
	"magick-app/lib/config"
	"magick-app/lib/image"
	"magick-app/lib/stat"

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
		Height:    768,
		Title:     "magick-app",
		JS:        js,
		CSS:       css,
		Colour:    "#18181f",
		Resizable: true,
	})
	newConfig := config.NewConfig()
	newStat := stat.NewStat()
	fileManager := image.NewFileManager(newConfig, newStat)

	app.Bind(basic)
	app.Bind(myTodoList)
	app.Bind(newConfig)
	app.Bind(newStat)
	app.Bind(fileManager)
	_ = app.Run()
}
