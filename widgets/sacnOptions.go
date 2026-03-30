package widgets

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewSacnOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	options := binding.BindStringList(&[]string{"multicast", "unicast"})
	sacnOptions := container.NewVBox(
		NewIpEntry(globalisation.SacnLocalAddressLabel, config.Sacn_source),
		NewEntrySlider(globalisation.PacketPriorityLabel, 1, 200, config.Sacn_priority),
		NewEntrySlider(globalisation.UniverseLabel, 1, 10, config.Sacn_universe),
		NewEntrySlider(globalisation.StartChannelLabel, 1, 255, config.Sacn_startChannel),
		NewSelectEntry(globalisation.TransmissionModeLabel, config.Sacn_mode, options),
		NewFormEntry(globalisation.SourceNameLabel, config.Sacn_sourceName),
		NewMultiIpEntry(globalisation.SacnUnicastTargetsLabel, config.Sacn_targets),
	)
	sacnOptions.Hide()

	return sacnOptions
}
