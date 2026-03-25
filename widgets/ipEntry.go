package widgets

import (
	"com/rkbx_launch/helpers"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewIPEntry(label string, configEntry *helpers.IPAddress) *fyne.Container {
	layer1Entry := widget.NewEntry()
	layer1Entry.SetText(fmt.Sprintf("%d", configEntry.Layer1))

	layer2Entry := widget.NewEntry()
	layer2Entry.SetText(fmt.Sprintf("%d", configEntry.Layer2))

	layer3Entry := widget.NewEntry()
	layer3Entry.SetText(fmt.Sprintf("%d", configEntry.Layer3))

	layer4Entry := widget.NewEntry()
	layer4Entry.SetText(fmt.Sprintf("%d", configEntry.Layer4))

	portEntry := widget.NewEntry()
	portEntry.SetText(fmt.Sprintf("%d", configEntry.Port))

	layer1Entry.OnChanged = func(s string) {
		if val, err := strconv.Atoi(s); err == nil {
			configEntry.Layer1 = uint8(val)
		}
	}

	layer2Entry.OnChanged = func(s string) {
		if val, err := strconv.Atoi(s); err == nil {
			configEntry.Layer2 = uint8(val)
		}
	}

	layer3Entry.OnChanged = func(s string) {
		if val, err := strconv.Atoi(s); err == nil {
			configEntry.Layer3 = uint8(val)
		}
	}

	layer4Entry.OnChanged = func(s string) {
		if val, err := strconv.Atoi(s); err == nil {
			configEntry.Layer4 = uint8(val)
		}
	}

	portEntry.OnChanged = func(s string) {
		if val, err := strconv.Atoi(s); err == nil {
			configEntry.Port = uint16(val)
		}
	}

	return container.NewVBox(
		widget.NewLabel(label),
		container.NewHBox(
			layer1Entry,
			widget.NewLabel("."),
			layer2Entry,
			widget.NewLabel("."),
			layer3Entry,
			widget.NewLabel("."),
			layer4Entry,
			widget.NewLabel(":"),
			portEntry),
	)
}
