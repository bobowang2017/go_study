package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	//err := os.Setenv("FYNE_FONT", "D:\\GoWorkSpace\\go_study\\jd\\alibaba.ttf")
	err := os.Setenv("FYNE_FONT", "alibaba.ttf")
	fmt.Println(err)
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("JD Login")
	//img := canvas.NewImageFromFile("D:\\GoWorkSpace\\go_study\\jd\\code.png")
	img := canvas.NewImageFromFile("code.png")
	img.FillMode = canvas.ImageFillOriginal
	img.Resize(fyne.NewSize(300, 300))

	leftLabel := widget.NewLabel("")
	leftLabel.Resize(fyne.NewSize(50, 300))

	rightLabel := widget.NewLabel("")
	rightLabel.Resize(fyne.NewSize(50, 300))

	bottomLabel := widget.NewLabelWithStyle("请打开京东APP扫描上方二维码进行登陆", fyne.TextAlignCenter, fyne.TextStyle{})
	content := container.NewBorder(nil, bottomLabel, leftLabel, rightLabel, img)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
