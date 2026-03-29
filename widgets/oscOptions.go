package widgets

import (
	"com/rkbx_launch/helpers"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewOscOptions(config *helpers.BoundRkbxConfig) *fyne.Container {
	cBind := binding.BindStruct(config)
	fmt.Println(cBind.GetValue("Osc_source"))

	oscOptions := container.NewVBox(
		NewIPEntry("Source IP Address", config.Osc_source),
		NewIPEntry("Destination IP Address", config.Osc_destination),
		NewEntrySlider("Send Every n-th Message", 1, 4, config.Osc_sendEveryNth),
		NewSelectEntry("Phrase Output Format", config.Osc_phraseOutputFormat, []string{"string", "int", "float"}),
	)
	oscOptions.Hide()

	return oscOptions
}
