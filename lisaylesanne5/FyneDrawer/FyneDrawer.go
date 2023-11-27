package fynedrawer

import (
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

type FyneDrawer struct {
	myApp fyne.App
	w     fyne.Window
}

func NewFyneDrawer() *FyneDrawer {
	fd := FyneDrawer{}
	fd.myApp = app.New()
	fd.w = fd.myApp.NewWindow("Drawer")
	return &fd
}

func (fd *FyneDrawer) DrawPoints(points []Point.Point) {
	pnts := []fyne.CanvasObject{}
	for _, point := range points {
		pnt := canvas.NewCircle(color.RGBA{0, 0, 0, 255})
		pnt.Resize(fyne.NewSize(2, 2))
		pnt.Move(fyne.NewPos(float32(point.X), float32(point.Y)))
		pnts = append(pnts, pnt)
	}

	content := container.NewWithoutLayout(pnts...)
	scroll := container.NewVScroll(content)
	fd.w.SetContent(scroll)
	fd.w.Resize(fyne.NewSize(1000, 1000))
	fd.w.ShowAndRun()
}
