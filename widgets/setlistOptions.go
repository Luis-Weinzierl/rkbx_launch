package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewSetlistOptions(config *helpers.BoundRkbxConfig) *fyne.Container {
	setlistOptions := container.NewVBox(
		NewFormEntry("Output Filename", config.Setlist_filename),
		NewFormEntry("Track Seperator", config.Setlist_seperator),
	)
	setlistOptions.Hide()

	return setlistOptions
}
