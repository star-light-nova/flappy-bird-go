package play

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/bird"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/obstacle"
)

func NewInfiniteWallObstacles(bird *bird.Bird, _window fyne.Window) *fyne.Container {
	wall := container.NewWithoutLayout()
	wall.Resize(fyne.NewSize(_window.Canvas().Size().Width, _window.Canvas().Size().Height))

	top, bot := obstacle.Generate(_window)

	targetX := -obstacle.RECTANGLE_WIDTH

	wall.Add(top.Rectangle)
	wall.Add(bot.Rectangle)

	go func() {
		for range time.Tick(time.Millisecond * 6) {
			nextPosTop := top.Position().SubtractXY(1.0, 0)
			nextPosBot := bot.Position().SubtractXY(1.0, 0)

			top.Move(nextPosTop)
			bot.Move(nextPosBot)

			if nextPosTop.X == targetX {
				wall.Remove(top.Rectangle)
				wall.Remove(bot.Rectangle)

				top, bot = obstacle.Generate(_window)

				wall.Add(top.Rectangle)
				wall.Add(bot.Rectangle)
			}

            topPos := top.Position()
            botPos := bot.Position()
            bPosX := bird.Position1.X
            bPosY := bird.Position1.Y

            if topPos.X < bPosX && topPos.X + obstacle.RECTANGLE_WIDTH > bPosX {
                if (bPosY > 0 && bPosY < topPos.Y)  || (bPosY > botPos.Y && bPosY < _window.Canvas().Size().Height) {
                    break
                }
            }
		}
	}()

	return wall
}
