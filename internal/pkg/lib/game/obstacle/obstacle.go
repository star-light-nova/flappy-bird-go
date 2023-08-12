package obstacle

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/constants"
)

type Obstacle struct {
	*canvas.Rectangle
}

const RECTANGLE_WIDTH = float32(55.0)

func New(_window fyne.Window) (top, bottom *Obstacle) {
	posTopLeft, posBottomLeft := passBoxPosition(_window)

	topRectangle := canvas.NewRectangle(color.White)
	bottomRectangle := canvas.NewRectangle(color.White)

	topRectangle.Move(posTopLeft)
	topRectangle.Resize(fyne.NewSize(RECTANGLE_WIDTH, -constants.FULL_HEIGHT))

	bottomRectangle.Move(posBottomLeft)
	bottomRectangle.Resize(fyne.NewSize(RECTANGLE_WIDTH, constants.FULL_HEIGHT))

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
	windowWidth := constants.FULL_WIDTH
	windowHeight := constants.FULL_HEIGHT

	posTopLeft.X = windowWidth
	posTopLeft.Y = randomPoint(RECTANGLE_WIDTH, windowHeight-RECTANGLE_WIDTH*2)

	posBottomLeft.X = posTopLeft.X
	posBottomLeft.Y = posTopLeft.Y + RECTANGLE_WIDTH

	return
}
