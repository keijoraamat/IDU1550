package fynedrawer

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

type FyneDrawer struct {
	app *app
	w   *app.Window
}

func NewFyneDrawer() *FyneDrawer {
	fd := FyneDrawer{}
	fd.app = app.NewApp()
	fd.w = fd.app.NewWindow("Drawer")
	return &fd
}

func (fd *FyneDrawer) DrawPoints(points []Point.Point) {
	pnts := []fyne.CanvasObject{}
	for _, point := range points {
		pnt := canvas.NewCircle(color.RGBA{0, 0, 0, 255})
		pnt.Resize(fyne.NewSize(2, 2))
		pnt.Move(fyne.NewPos(point.X, point.Y))
		pnts = append(pnts, pnt)
	}

	content := container.NewWithoutLayout(points...)
	scroll := container.NewVScroll(content)
	fd.w.SetContent(container.NewBorder(nil, nil, nil, nil,
		scroll,
	))
	fd.w.Resize(fyne.NewSize(1000, 1000))
	fd.w.ShowAndRun()
}
