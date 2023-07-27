package bird

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func NewBird() *canvas.Circle {
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 1.0
	circle.StrokeColor = color.Black
	circle.Resize(fyne.NewSize(20.0, 20.0))

	return circle
}
