package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

    "github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/play"
)

func Home(_window fyne.Window) *fyne.Container {
    homeContainer := container.NewCenter(homeButtonsContainer(_window))

    return homeContainer
}

func homeButtonsContainer(_window fyne.Window) *fyne.Container {
    playButton := widget.NewButton("Play", func(){
        _window.SetContent(play.PlayScreen(_window))
    })

    quitButton := widget.NewButton("Quit", func() {
        _window.Close()
    })

    vbox := container.NewVBox(playButton, quitButton)

    // FIX: Does not work for now.
    vbox.Resize(fyne.NewSize(550.0, 550.0))

    return vbox
}

