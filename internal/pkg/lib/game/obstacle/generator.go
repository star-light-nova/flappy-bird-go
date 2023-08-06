package obstacle

import "fyne.io/fyne/v2"

func Generate(_window fyne.Window) (topObs, botObs *Obstacle) {
	topObs, botObs = New(_window)

	return
}
