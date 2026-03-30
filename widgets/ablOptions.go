package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewAblOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	ablOptions := container.NewVBox(
		NewEntrySliderF("Cumulative Error Tolerance", 0.01, 0.1, config.Link_cumulativeErrorTolerance),
	)
	ablOptions.Hide()

	return ablOptions
}
