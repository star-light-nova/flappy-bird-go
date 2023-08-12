package flappybirdgo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/menu"
    "github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/content_manager"
)

func Start() {
	_app := app.New()
    // TODO: Implement settings for the app.
    // _settings := _app.Settings()
	_window := _app.NewWindow("Hello World")

    _window.Resize(fyne.NewSize(550.0, 550.0))

    _window.SetTitle("Flappy Bird Go")
    _window.RequestFocus()
    _window.CenterOnScreen()

    contentmanager.New(_window)

	// _window.SetContent(menu.Home(_window))

    _window.ShowAndRun()
}
