package play

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/obstacle"
)

func NewInfiniteWallObstacles(_window fyne.Window) *fyne.Container {
	wall := container.NewWithoutLayout()
	wall.Resize(fyne.NewSize(_window.Canvas().Size().Width, _window.Canvas().Size().Height))

    obs := obstacle.Generate(_window)

	target := -wall.Size().Width - obstacle.RECTANGLE_WIDTH

	wall.Add(obs)

	go func() {
		for range time.Tick(time.Millisecond * 6) {
			nextPos := obs.Position().SubtractXY(1.0, 0)

			obs.Move(nextPos)

			if nextPos.X == target {
				wall.Remove(obs)

				obs = obstacle.Generate(_window)

				wall.Add(obs)
			}
		}
	}()

	return wall
}
