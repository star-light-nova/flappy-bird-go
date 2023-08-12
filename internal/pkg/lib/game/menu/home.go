package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// "github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/play"
)

func Home() (*fyne.Container, *widget.Button, *widget.Button) {
	homeBtnsContainer, playBtn, quitBtn := homeButtonsContainer()

    homeContainer := container.NewCenter(homeBtnsContainer)

	return homeContainer, playBtn, quitBtn
}

func homeButtonsContainer() (*fyne.Container, *widget.Button, *widget.Button) {
	playButton := widget.NewButton("Play", func() {})

	quitButton := widget.NewButton("Quit", func() {})

	vbox := container.NewVBox(playButton, quitButton)

	return vbox, playButton, quitButton
}
