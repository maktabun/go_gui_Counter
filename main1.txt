package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	counterWindow := app.New()

	window := counterWindow.NewWindow("Counter App")

	counter := 0
	mylabel := widget.NewLabel("Counter: 0")

	mybutton := widget.NewButton("Increment", func() {
		counter++
		mylabel.SetText("Counter: " + strconv.Itoa(counter))
	})

	content := container.NewVBox(mylabel, mybutton)
	window.SetContent(content)
	window.Resize(fyne.NewSize(300, 200))

	window.ShowAndRun()
}
