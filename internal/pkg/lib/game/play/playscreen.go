package play

import (
	"time"

	"github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/bird"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	SPACE_CODE = 49
	UP         = 25.0
	DOWN       = 0.1
)

// OPTIMIZE: Refactor
func PlayScreen(_window fyne.Window) *fyne.Container {
	playScreen := container.NewWithoutLayout()
	playScreen.Resize(fyne.NewSize(500.0, 500.0))

	bird := bird.NewBird()
	bird.Move(fyne.NewPos(playScreen.Size().Width/2.0-200.0, playScreen.Size().Height/2.0))

	if desktopCanvas, ok := _window.Canvas().(desktop.Canvas); ok {
		desktopCanvas.SetOnKeyUp(func(key *fyne.KeyEvent) {
			jumpOnSpace(key, bird)
		})
	}

	go func() {
		fall(bird)
	}()

	playScreen.AddObject(bird)

	return playScreen
}

func jump(bird *canvas.Circle) {
	move := canvas.NewPositionAnimation(bird.Position1, fyne.NewPos(bird.Position1.X, bird.Position1.Y-UP), time.Millisecond, bird.Move)
	move.Start()
}

func fall(bird *canvas.Circle) {
	for range time.Tick(time.Millisecond) {
		bird.Move(fyne.NewPos(bird.Position1.X, bird.Position1.Y+DOWN))
	}
}

func jumpOnSpace(key *fyne.KeyEvent, bird *canvas.Circle) {
	if key.Physical.ScanCode == SPACE_CODE {
		jump(bird)
	}
}
