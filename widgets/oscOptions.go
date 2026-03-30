package widgets

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewOscOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	options := binding.BindStringList(&[]string{"string", "int", "float"})

	oscOptions := container.NewVBox(
		NewIpPortEntry(globalisation.SourceIpAddressLabel, config.Osc_source),
		NewIpPortEntry(globalisation.DestinationIpAddressLabel, config.Osc_destination),
		NewEntrySlider(globalisation.SendEveryNthMessage, 1, 4, config.Osc_sendEveryNth),
		NewSelectEntry(globalisation.PhraseOutputFormatLabel, config.Osc_phraseOutputFormat, options),

		NewBoolConfig(globalisation.MasterTimeLabel, config.Osc_msg_masterTime),
		NewBoolConfig(globalisation.MasterPhraseLabel, config.Osc_msg_masterPhrase),
		NewBoolConfig(globalisation.NonMasterTimeLabel, config.Osc_msg_nTime),
		NewBoolConfig(globalisation.NonMasterPhraseLabel, config.Osc_msg_nPhrase),

		NewOptionalEntrySliderF(globalisation.MasterBeatSubdivisionLabel, 0.25, 4, 0.25, config.Osc_msg_masterBeatSubdiv, config.Osc_msg_masterBeatSubdivEnabled),
		NewOptionalEntrySliderF(globalisation.MasterBeatTriggerLabel, 0.25, 4, 0.25, config.Osc_msg_masterBeatTrigger, config.Osc_msg_masterBeatTriggerEnabled),
		NewOptionalEntrySliderF(globalisation.NonMasterBeatSubdivisionLabel, 0.25, 4, 0.25, config.Osc_msg_nBeatSubdiv, config.Osc_msg_nBeatSubdivEnabled),
		NewOptionalEntrySliderF(globalisation.NonMasterBeatTriggerLabel, 0.25, 4, 0.25, config.Osc_msg_nBeatTrigger, config.Osc_msg_nBeatTriggerEnabled),
	)
	oscOptions.Hide()

	return oscOptions
}
