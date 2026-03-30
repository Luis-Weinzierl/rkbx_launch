package widgets

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewAppOptions(config *helpers.RkbxLinkConfig) *fyne.Container {
	return container.NewVBox(
		NewSelectEntry(globalisation.RekordboxVersionLabel, config.Keeper_rekordboxVersion, config.AvailableRekordboxVersions),
		NewBoolConfig(globalisation.KeepWarmLabel, config.Keeper_keepWarm),
		NewEntrySlider(globalisation.UpdateRateLabel, 10, 500, config.Keeper_updateRate),
		NewEntrySlider(globalisation.SlowUpdateEveryNthLabel, 5, 20, config.Keeper_slowUpdateEveryNth),
		NewEntrySlider(globalisation.DelayCompensationLabel, -5, 5, config.Keeper_delayCompensation),
		NewEntrySlider(globalisation.ActiveDecksLabel, 2, 4, config.Keeper_decks),
	)
}
