package play

import (
	"github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/bird"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func PlayScreen(_window fyne.Window) *fyne.Container {
	playScreen := container.NewWithoutLayout()
	playScreen.Resize(fyne.NewSize(550.0, 550.0))
    playScreen.Move(fyne.NewPos(0, 0))

	bird := bird.New(_window)
	bird.Move(fyne.NewPos(playScreen.Size().Width/2.0-200.0, playScreen.Size().Height/2.0))

    infWall := NewInfiniteWallObstacles(bird, _window)

	playScreen.Add(bird.Circle)
    playScreen.Add(infWall)

	return playScreen
}

