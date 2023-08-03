package obstacle

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Obstacle struct {
	*canvas.Rectangle
}

const RECTANGLE_WIDTH = 55.0

func New(_window fyne.Window) (top, bottom *Obstacle) {
	posTopLeft, posBottomLeft := passBoxPosition(_window)

	topRectangle := canvas.NewRectangle(color.White)
	bottomRectangle := canvas.NewRectangle(color.White)

	topRectangle.Move(posTopLeft)
	topRectangle.Resize(fyne.NewSize(RECTANGLE_WIDTH, -_window.Canvas().Size().Height*0.5))

	bottomRectangle.Move(posBottomLeft)
	bottomRectangle.Resize(fyne.NewSize(RECTANGLE_WIDTH, _window.Canvas().Size().Height))

	top = &Obstacle{
		topRectangle,
	}

	bottom = &Obstacle{
		bottomRectangle,
	}

	return
}

func randomPoint(start float32, end float32) float32 {
	rand.Seed(time.Now().UnixNano())
	return end*rand.Float32() + start
}

func passBoxPosition(_window fyne.Window) (posTopLeft, posBottomLeft fyne.Position) {
	windowHalfWidth := float32(550.0 * 0.5)

	posTopLeft.X = randomPoint(windowHalfWidth, 250.0)
	posTopLeft.Y = randomPoint(0, 250.0)

	posBottomLeft.X = posTopLeft.X
	posBottomLeft.Y = posTopLeft.Y + RECTANGLE_WIDTH

	return
}
