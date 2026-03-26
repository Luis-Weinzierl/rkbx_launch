package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewOscOptions(config *helpers.RkbxConfig) *fyne.Container {
	oscOptions := container.NewVBox(
		NewIPEntry("Source IP Address", &config.Osc_source),
		NewIPEntry("Destination IP Address", &config.Osc_destination),
		NewEntrySlider("Send Every n-th Message", 1, 4, &config.Osc_sendEveryNth),
		NewSelectEntry("Phrase Output Format", &config.Osc_phraseOutputFormat, []string{"string", "int", "float"}),
	)
	oscOptions.Hide()

	return oscOptions
}
