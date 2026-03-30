package widgets

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewInfoLabel(text string) *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(globalisation.Get(text)),
		NewInfoHover(text),
	)
}
