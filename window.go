package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

// CreateTypingWindow creates and configures the main application window
func CreateTypingWindow() (fyne.Window, *canvas.Text) {
	myApp := app.New()

	var window fyne.Window
	if drv, ok := myApp.Driver().(desktop.Driver); ok {
		window = drv.CreateSplashWindow()
	} else {
		window = myApp.NewWindow("Prototypage")
	}

	window.SetFixedSize(true)
	window.SetPadded(false)

	text := canvas.NewText("Waiting...", color.White)
	text.TextSize = 14
	content := container.NewCenter(text)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 100))
	window.CenterOnScreen()

	return window, text
}
