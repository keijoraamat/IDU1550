package pointdrawer

import (
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

type PointDrawer struct {
	myApp fyne.App
	w     fyne.Window
}

func NewPointDrawer() *PointDrawer {
	fd := PointDrawer{}
	fd.myApp = app.New()
	fd.w = fd.myApp.NewWindow("Drawer")
	return &fd
}

func (fd *PointDrawer) DrawPoints(points []Point.Point) {
	pnts := []fyne.CanvasObject{}
	for _, point := range points {
		pnt := canvas.NewCircle(color.RGBA{0, 255, 0, 255})
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
