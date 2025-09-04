package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne counter example")

	count := 0

	// Use a label instead for better text alignment control
	txtNumber := widget.NewLabel("0")
	txtNumber.Alignment = fyne.TextAlignTrailing // This works for Label!
	txtNumber.Resize(fyne.NewSize(200, txtNumber.MinSize().Height))

	updateText := func() {
		txtNumber.SetText(strconv.Itoa(count))
	}

	// Create the minus button
	minusButton := widget.NewButton("-", func() {
		count--
		updateText()
	})

	// Create the plus button
	plusButton := widget.NewButton("+", func() {
		count++
		updateText()
	})

	// Create a centered row
	content := container.NewHBox(
		minusButton,
		txtNumber,
		plusButton,
	)

	// Center the content in the window
	centered := container.New(layout.NewCenterLayout(), content)

	myWindow.SetContent(centered)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
