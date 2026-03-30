package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewFileOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	fileOptions := container.NewVBox(
		NewFormEntry("Output Filename", config.File_fileName),
	)
	fileOptions.Hide()

	return fileOptions
}
