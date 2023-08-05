package obstacle

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
)

func Generate(_window fyne.Window) *fyne.Container {
    topObs, bottomObs := New(_window)

    obstacle := container.NewWithoutLayout(topObs.Rectangle, bottomObs.Rectangle)

    return obstacle
}
