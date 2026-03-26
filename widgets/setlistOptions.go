package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewSetlistOptions(config *helpers.RkbxConfig) *fyne.Container {
	setlistOptions := container.NewVBox(
		NewFormEntry("Track Seperator", &config.Setlist_seperator),
		NewFormEntry("Output Filename", &config.Setlist_filename),
	)
	setlistOptions.Hide()

	return setlistOptions
}
