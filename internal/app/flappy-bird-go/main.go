package flappybirdgo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/menu"
)

func Start() {
	_app := app.New()
	_window := _app.NewWindow("Hello World")

    _window.SetTitle("Flappy Bird Go")
    _window.Resize(fyne.NewSize(550.0, 550.0))

    // _window.SetMainMenu(menu.Home(_window))
	_window.SetContent(menu.Home(_window))

    _window.ShowAndRun()
}
