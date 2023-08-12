package bird

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"

	"image/color"
	"time"
)

const (
	UP          = float32(0.25)
	DOWN        = float32(0.1)
    TOLLERANCE  = float32(0.001)
	JUMP_HEIGHT = float32(18.0)
	SPACE_CODE  = 49
)

type Bird struct {
	*canvas.Circle
}

func New(_window fyne.Window) *Bird {
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 1.0
	circle.StrokeColor = color.Black
	circle.Resize(fyne.NewSize(20.0, 20.0))

	bird := &Bird{circle}

	if desktopCanvas, ok := _window.Canvas().(desktop.Canvas); ok {
		desktopCanvas.SetOnKeyUp(func(key *fyne.KeyEvent) {
			bird.JumpOnSpace(key)
		})
	}

	go func() {
		bird.Fall()
	}()

	return bird
}

func (b Bird) Jump() {
	target := b.Position().AddXY(0, -JUMP_HEIGHT)
    nextPosition := b.Position1.SubtractXY(0, UP)

	for range time.Tick(time.Millisecond) {
        diff := abs32(nextPosition.Y - target.Y)

        if diff <= TOLLERANCE || diff >= JUMP_HEIGHT {
            break
        }

		b.Move(nextPosition)

        nextPosition = b.Position1.SubtractXY(0, UP)
	}
}

func (b Bird) Fall() {
	for range time.Tick(time.Millisecond) {
		b.Move(fyne.NewPos(b.Position1.X, b.Position1.Y+DOWN))
	}
}

func (b Bird) JumpOnSpace(key *fyne.KeyEvent) {
	if key.Physical.ScanCode == SPACE_CODE {
		b.Jump()
	}
}

func abs32(f float32) float32 {
    if f > 0 {
        return f
    }

    return -f
}
