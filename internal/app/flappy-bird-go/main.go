package flappybirdgo

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Start() {
	_app := app.New()
	_window := _app.NewWindow("Hello World")

	_window.SetContent(
		container.New(layout.NewMaxLayout(),
			widget.NewLabel("Hello World"),
        ),
    )

    _window.ShowAndRun()
}
