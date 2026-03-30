package widgets

import (
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func NewSacnOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	options := binding.BindStringList(&[]string{"multicast", "unicast"})
	sacnOptions := container.NewVBox(
		NewIpEntry("Local Address to bind or sACN", config.Sacn_source),
		NewEntrySlider("Packet Priority", 1, 200, config.Sacn_priority),
		NewEntrySlider("Universe", 1, 10, config.Sacn_universe),
		NewEntrySlider("Start Channel", 1, 255, config.Sacn_startChannel),
		NewSelectEntry("Transmission mode", config.Sacn_mode, options),
		NewFormEntry("Source Name for packets", config.Sacn_sourceName),
		NewMultiIpEntry("sACN Unicast Targets", config.Sacn_targets),
	)
	sacnOptions.Hide()

	return sacnOptions
}
