package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewOscOptions(config *helpers.BoundRkbxConfig) *fyne.Container {
	options := binding.BindStringList(&[]string{"string", "int", "float"})

	oscOptions := container.NewVBox(
		NewIpPortEntry("Source IP Address", config.Osc_source),
		NewIpPortEntry("Destination IP Address", config.Osc_destination),
		NewEntrySlider("Send Every n-th Message", 1, 4, config.Osc_sendEveryNth),
		NewSelectEntry("Phrase Output Format", config.Osc_phraseOutputFormat, options),
	)
	oscOptions.Hide()

	return oscOptions
}
