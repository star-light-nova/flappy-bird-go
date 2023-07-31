package play

import (
	"github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/bird"
	"github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/obstacle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

const SPACE_CODE = 49

func PlayScreen(_window fyne.Window) *fyne.Container {
	playScreen := container.NewWithoutLayout()
	playScreen.Resize(fyne.NewSize(500.0, 500.0))

	bird := bird.New()
	bird.Move(fyne.NewPos(playScreen.Size().Width/2.0-200.0, playScreen.Size().Height/2.0))

	if desktopCanvas, ok := _window.Canvas().(desktop.Canvas); ok {
		desktopCanvas.SetOnKeyUp(func(key *fyne.KeyEvent) {
			jumpOnSpace(key, bird)
		})
	}

	playScreen.AddObject(bird.Circle)
    playScreen.AddObject(obstacle.New().Rectangle)

    go func() {
        bird.Fall()
    }()

	return playScreen
}

func jumpOnSpace(key *fyne.KeyEvent, bird *bird.Bird) {
	if key.Physical.ScanCode == SPACE_CODE {
		bird.Jump()
	}
}
