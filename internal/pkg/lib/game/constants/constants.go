package constants

import "fyne.io/fyne/v2"

const (
	FULL_HEIGHT = float32(550.0)
	FULL_WIDTH  = float32(550.0)
)

func FULL_SIZE() fyne.Size {
    return fyne.NewSize(FULL_WIDTH, FULL_HEIGHT)
}

func BIRD_INITIAL_POS() fyne.Position {
    return fyne.NewPos(FULL_WIDTH/2.0-200.0, FULL_HEIGHT/2.0)
}
