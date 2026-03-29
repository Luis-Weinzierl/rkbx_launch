package widgets

import (
	"com/rkbx_launch/helpers"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewIPEntry(label string, configEntry *helpers.IPAddress) *fyne.Container {
	layer1Binding := binding.BindInt(&configEntry.Layer1)
	layer1BindingS := binding.IntToString(layer1Binding)
	layer1Entry := widget.NewEntryWithData(layer1BindingS)

	layer2Binding := binding.BindInt(&configEntry.Layer2)
	layer2BindingS := binding.IntToString(layer2Binding)
	layer2Entry := widget.NewEntryWithData(layer2BindingS)

	layer3Binding := binding.BindInt(&configEntry.Layer3)
	layer3BindingS := binding.IntToString(layer3Binding)
	layer3Entry := widget.NewEntryWithData(layer3BindingS)

	layer4Binding := binding.BindInt(&configEntry.Layer4)
	layer4BindingS := binding.IntToString(layer4Binding)
	layer4Entry := widget.NewEntryWithData(layer4BindingS)

	portBinding := binding.BindInt(&configEntry.Port)
	portBindingS := binding.IntToString(portBinding)
	portEntry := widget.NewEntryWithData(portBindingS)

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
	layer1Binding := binding.BindInt(&configEntry.Layer1)
	layer1BindingS := binding.IntToString(layer1Binding)
	layer1Entry := widget.NewEntryWithData(layer1BindingS)

	layer2Binding := binding.BindInt(&configEntry.Layer2)
	layer2BindingS := binding.IntToString(layer2Binding)
	layer2Entry := widget.NewEntryWithData(layer2BindingS)

	layer3Binding := binding.BindInt(&configEntry.Layer3)
	layer3BindingS := binding.IntToString(layer3Binding)
	layer3Entry := widget.NewEntryWithData(layer3BindingS)

	layer4Binding := binding.BindInt(&configEntry.Layer4)
	layer4BindingS := binding.IntToString(layer4Binding)
	layer4Entry := widget.NewEntryWithData(layer4BindingS)

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

func changedCallback(
	layer1Entry *widget.Entry,
	layer2Entry *widget.Entry,
	layer3Entry *widget.Entry,
	layer4Entry *widget.Entry,
	portEntry *widget.Entry,
	invalidityLabel *canvas.Text,
	ip *helpers.IPAddress,
) {
	layer1, err1 := strconv.Atoi(layer1Entry.Text)
	layer2, err2 := strconv.Atoi(layer2Entry.Text)
	layer3, err3 := strconv.Atoi(layer3Entry.Text)
	layer4, err4 := strconv.Atoi(layer4Entry.Text)
	port, err5 := strconv.Atoi(portEntry.Text)

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
	layer1, err1 := strconv.Atoi(layer1Entry.Text)
	layer2, err2 := strconv.Atoi(layer2Entry.Text)
	layer3, err3 := strconv.Atoi(layer3Entry.Text)
	layer4, err4 := strconv.Atoi(layer4Entry.Text)

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
