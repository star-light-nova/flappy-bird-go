package contentmanager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"

	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/menu"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/play"
)

func New(_window fyne.Window) {
	gameEndChan := make(chan bool)
	mainMenu, playBtn, quitBtn := menu.Home()

	playBtn.OnTapped = func() {
        playScreen := play.PlayScreen(_window, gameEndChan)
		_window.SetContent(playScreen)
	}

	quitBtn.OnTapped = func() {
		_window.Close()
	}

	go func() {
		for {
            isGameEnded := <- gameEndChan

			if isGameEnded {
				dialog.ShowConfirm("You have lost", "Restart?", func(b bool) {
					if b {
						_window.SetContent(play.PlayScreen(_window, gameEndChan))
					} else {
						_window.SetContent(mainMenu)
					}
				}, _window)
			}
		}
	}()

    _window.SetContent(mainMenu)
}
