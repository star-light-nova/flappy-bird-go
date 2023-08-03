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

    topObs, bottomObs := obstacle.New(_window)

    wall.Add(topObs.Rectangle)
    wall.Add(bottomObs.Rectangle)

    initial := wall.Position()
    target := _window.Canvas().Content().Position().X - 650.0

    go func() {
        for range time.Tick(time.Millisecond * 3) {
            nextPos := wall.Position().SubtractXY(1.0, 0)

            wall.Move(nextPos)

            if nextPos.X == target {
                wall.Move(initial)
            }
        }
    }()

    return wall
}
