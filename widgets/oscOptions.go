package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewOscOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	options := binding.BindStringList(&[]string{"string", "int", "float"})

	oscOptions := container.NewVBox(
		NewIpPortEntry("Source IP Address", config.Osc_source),
		NewIpPortEntry("Destination IP Address", config.Osc_destination),
		NewEntrySlider("Send Every n-th Message", 1, 4, config.Osc_sendEveryNth),
		NewSelectEntry("Phrase Output Format", config.Osc_phraseOutputFormat, options),

		NewBoolConfig("Master Time", config.Osc_msg_masterTime),
		NewBoolConfig("Master Phrase", config.Osc_msg_masterPhrase),
		NewBoolConfig("Non-Master Time", config.Osc_msg_nTime),
		NewBoolConfig("Non-Master Phrase", config.Osc_msg_nPhrase),

		NewOptionalEntrySliderF("Master Beat Subdivision", 0.25, 4, 0.25, config.Osc_msg_masterBeatSubdiv, config.Osc_msg_masterBeatSubdivEnabled),
		NewOptionalEntrySliderF("Master Beat Trigger", 0.25, 4, 0.25, config.Osc_msg_masterBeatTrigger, config.Osc_msg_masterBeatTriggerEnabled),
		NewOptionalEntrySliderF("Non-Master Beat Subdivision", 0.25, 4, 0.25, config.Osc_msg_nBeatSubdiv, config.Osc_msg_nBeatSubdivEnabled),
		NewOptionalEntrySliderF("Non-Master Beat Trigger", 0.25, 4, 0.25, config.Osc_msg_nBeatTrigger, config.Osc_msg_nBeatTriggerEnabled),
	)
	oscOptions.Hide()

	return oscOptions
}
