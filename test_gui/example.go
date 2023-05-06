package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.Black)
	left := canvas.NewText("left", color.Gray{
		Y: 0x99,
	})
	right := canvas.NewText("right", color.Gray{
		Y: 0x99,
	})
	bottom := canvas.NewText("left", color.Gray{
		Y: 0x88,
	})
	middle := canvas.NewText("content", color.Gray16{Y: 0x98})
	content := container.NewBorder(top, bottom, left, right, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
