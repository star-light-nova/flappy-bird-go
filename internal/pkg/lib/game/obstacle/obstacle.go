package obstacle

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Obstacle struct {
    *canvas.Rectangle
}

func New() *Obstacle {
    rectangle := canvas.NewRectangle(color.White)

    rectangle.Resize(fyne.NewSize(50.0, 100.0))

    return &Obstacle{
        rectangle,
    }
}
