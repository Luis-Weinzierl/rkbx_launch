package widgets

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewSetlistOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	setlistOptions := container.NewVBox(
		NewFormEntry(globalisation.OutputFilenameLabel, config.Setlist_filename),
		NewFormEntry(globalisation.TrackSepatorLabel, config.Setlist_seperator),
	)
	setlistOptions.Hide()

	return setlistOptions
}
