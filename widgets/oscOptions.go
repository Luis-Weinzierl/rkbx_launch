package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewOscOptions(config *helpers.RkbxConfig) *fyne.Container {
	oscOptions := container.NewVBox(
		NewIPEntry("Destination IP Address", &config.Osc_destination),
	)
	oscOptions.Hide()

	return oscOptions
}
