package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"log"
	"magick-app/demo"
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
		Width:  1024,
		Height: 768,
		Title:  "magick-app",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(myTodoList)
	app.Run()
}
