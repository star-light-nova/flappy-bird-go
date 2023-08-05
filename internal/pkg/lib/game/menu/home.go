package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

    "github.com/StarLightNova/flappy-bird-go/internal/pkg/lib/game/play"

    "log"
)

func Home(_window fyne.Window) *fyne.Container {
    homeContainer := container.NewCenter(homeButtonsContainer(_window))

    return homeContainer
}

func homeButtonsContainer(_window fyne.Window) *fyne.Container {
    playButton := widget.NewButton("Play", func(){
        log.Println("Play button tapped.")

        _window.SetContent(play.PlayScreen(_window))
    })

    quitButton := widget.NewButton("Quit", func() {
        log.Println("Quit button tapped.")

        _window.Close()
    })

    vbox := container.NewVBox(playButton, quitButton)

    // FIX: Does not work for now.
    vbox.Resize(fyne.NewSize(550.0, 550.0))

    return vbox
}

