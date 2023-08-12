package flappybirdgo

import (
	"fyne.io/fyne/v2/app"

	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/content_manager"
	"github.com/star-light-nova/flappy-bird-go/internal/pkg/lib/game/constants"
)

func Start() {
	_app := app.New()
	// TODO: Implement settings for the app.
	// _settings := _app.Settings()
	_window := _app.NewWindow("Hello World")

	_window.Resize(constants.FULL_SIZE())

	_window.SetTitle("Flappy Bird Go")
	_window.RequestFocus()
	_window.CenterOnScreen()
	_window.SetFixedSize(true)

	contentmanager.New(_window)

	_window.ShowAndRun()
}
