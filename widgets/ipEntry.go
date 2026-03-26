package widgets

import (
	"com/rkbx_launch/helpers"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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

	invalidityLabel := canvas.NewText("IP Address is invalid", fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForegroundOnError, 0))
	invalidityLabel.Hide()

	callback := func(s string) {
		changedCallback(
			layer1Entry,
			layer2Entry,
			layer3Entry,
			layer4Entry,
			portEntry,
			invalidityLabel,
			configEntry)
	}

	layer1Entry.OnChanged = callback
	layer2Entry.OnChanged = callback
	layer3Entry.OnChanged = callback
	layer4Entry.OnChanged = callback
	portEntry.OnChanged = callback

	return container.NewVBox(
		widget.NewLabel(label),
		container.NewGridWithColumns(
			5,
			container.NewBorder(nil, nil, nil, widget.NewLabel("."), layer1Entry),
			container.NewBorder(nil, nil, nil, widget.NewLabel("."), layer2Entry),
			container.NewBorder(nil, nil, nil, widget.NewLabel("."), layer3Entry),
			container.NewBorder(nil, nil, nil, widget.NewLabel(":"), layer4Entry),
			portEntry),
		invalidityLabel,
	)
}

func NewIPOnlyEntry(label string, configEntry *helpers.IPAddress) *fyne.Container {
	layer1Entry := widget.NewEntry()
	layer1Entry.SetText(fmt.Sprintf("%d", configEntry.Layer1))

	layer2Entry := widget.NewEntry()
	layer2Entry.SetText(fmt.Sprintf("%d", configEntry.Layer2))

	layer3Entry := widget.NewEntry()
	layer3Entry.SetText(fmt.Sprintf("%d", configEntry.Layer3))

	layer4Entry := widget.NewEntry()
	layer4Entry.SetText(fmt.Sprintf("%d", configEntry.Layer4))

	invalidityLabel := canvas.NewText("IP Address is invalid", fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForegroundOnError, 0))
	invalidityLabel.Hide()

	callback := func(s string) {
		changedCallbackIpOnly(
			layer1Entry,
			layer2Entry,
			layer3Entry,
			layer4Entry,
			invalidityLabel,
			configEntry)
	}

	layer1Entry.OnChanged = callback
	layer2Entry.OnChanged = callback
	layer3Entry.OnChanged = callback
	layer4Entry.OnChanged = callback

	return container.NewVBox(
		widget.NewLabel(label),
		container.NewGridWithColumns(
			7,
			layer1Entry,
			widget.NewLabel("."),
			layer2Entry,
			widget.NewLabel("."),
			layer3Entry,
			widget.NewLabel("."),
			layer4Entry),
		invalidityLabel,
	)
}

func tryParseUint8(s string) (uint8, error) {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	} else if val > 255 || val < 0 {
		return 0, fmt.Errorf("Value %d is outside of [0;255]", val)
	}

	return uint8(val), nil
}

func tryParseUint16(s string) (uint16, error) {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	} else if val > 65535 || val < 0 {
		return 0, fmt.Errorf("Value %d is outside of [0;65535]", val)
	}

	return uint16(val), nil
}

func changedCallback(
	layer1Entry *widget.Entry,
	layer2Entry *widget.Entry,
	layer3Entry *widget.Entry,
	layer4Entry *widget.Entry,
	portEntry *widget.Entry,
	invalidityLabel *canvas.Text,
	ip *helpers.IPAddress,
) {
	layer1, err1 := tryParseUint8(layer1Entry.Text)
	layer2, err2 := tryParseUint8(layer2Entry.Text)
	layer3, err3 := tryParseUint8(layer3Entry.Text)
	layer4, err4 := tryParseUint8(layer4Entry.Text)
	port, err5 := tryParseUint16(portEntry.Text)

	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil ||
		err5 != nil {
		invalidityLabel.Show()
	} else {
		invalidityLabel.Hide()

		ip.Layer1 = layer1
		ip.Layer2 = layer2
		ip.Layer3 = layer3
		ip.Layer4 = layer4
		ip.Port = port
	}
}

func changedCallbackIpOnly(
	layer1Entry *widget.Entry,
	layer2Entry *widget.Entry,
	layer3Entry *widget.Entry,
	layer4Entry *widget.Entry,
	invalidityLabel *canvas.Text,
	ip *helpers.IPAddress,
) {
	layer1, err1 := tryParseUint8(layer1Entry.Text)
	layer2, err2 := tryParseUint8(layer2Entry.Text)
	layer3, err3 := tryParseUint8(layer3Entry.Text)
	layer4, err4 := tryParseUint8(layer4Entry.Text)

	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil {
		invalidityLabel.Show()
	} else {
		invalidityLabel.Hide()

		ip.Layer1 = layer1
		ip.Layer2 = layer2
		ip.Layer3 = layer3
		ip.Layer4 = layer4
	}
}
