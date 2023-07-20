package play

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PlayScreen() *fyne.Container {
    playScreen := container.NewPadded(widget.NewLabel("Flying Bird"))

    return playScreen
}
