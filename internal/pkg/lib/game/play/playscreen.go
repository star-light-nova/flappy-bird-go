package play

import (
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/bird"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func PlayScreen(_window fyne.Window, gameEndChan chan<- bool) *fyne.Container {
	playScreen := container.NewWithoutLayout()
	playScreen.Resize(constants.FULL_SIZE())
	playScreen.Move(fyne.NewPos(0, 0))

	bird := bird.New(_window)
	bird.Move(constants.BIRD_INITIAL_POS())

	infWall := NewInfiniteWallObstacles(bird, _window, gameEndChan)

	playScreen.Add(bird.Circle)
	playScreen.Add(infWall)

	return playScreen
}
