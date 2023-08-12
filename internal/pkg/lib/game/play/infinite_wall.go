package play

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/bird"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/constants"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/obstacle"
)

func NewInfiniteWallObstacles(bird *bird.Bird, _window fyne.Window, gameEndChan chan<- bool) *fyne.Container {
	wall := container.NewWithoutLayout()
	wall.Resize(constants.FULL_SIZE())

	top, bot := obstacle.Generate(_window)

	wall.Add(top.Rectangle)
	wall.Add(bot.Rectangle)

	go wallMovement(wall, top, bot, bird, _window, gameEndChan)

	return wall
}

func wallMovement(wall *fyne.Container, top, bot *obstacle.Obstacle, bird *bird.Bird, _window fyne.Window, gameEndChan chan<- bool) {
	targetX := -obstacle.RECTANGLE_WIDTH

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
        birdPosTop := bird.Position1
        birdPosBot := bird.Position2

		if isCollidedWithObstacle(topPos, botPos, birdPosTop, birdPosBot) {
			gameEndChan <- true
			break
		}
	}
}

func isCollidedWithObstacle(topPos, botPos, birdPosTop, birdPosBot fyne.Position) bool {
    bPosX := birdPosTop.X + bird.RADIUS / 2.0
    
    bPosYTop := birdPosTop.Y
    bPosYBot := birdPosBot.Y

	if topPos.X < bPosX && topPos.X+obstacle.RECTANGLE_WIDTH > bPosX {
		if (bPosYTop > 0 && bPosYTop < topPos.Y) || (bPosYBot > botPos.Y && bPosYBot < constants.FULL_HEIGHT) {
            return true
		}
	}

    if bPosYTop <= 0 || bPosYBot >= constants.FULL_HEIGHT {
        return true
    }

    return false
}
