package bird

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"image/color"
	"time"
)

const (
	UP   = 25.0
	DOWN = 0.1
)

type Bird struct {
	*canvas.Circle
}

func New() *Bird {
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 1.0
	circle.StrokeColor = color.Black
	circle.Resize(fyne.NewSize(20.0, 20.0))

	return &Bird{
		circle,
	}
}

func (b Bird) Jump() {
	move := canvas.NewPositionAnimation(b.Position1, fyne.NewPos(b.Position1.X, b.Position1.Y-UP), time.Millisecond, b.Move)
	move.Start()
}

func (b Bird) Fall() {
	for range time.Tick(time.Millisecond) {
		b.Move(fyne.NewPos(b.Position1.X, b.Position1.Y+DOWN))
	}
}
