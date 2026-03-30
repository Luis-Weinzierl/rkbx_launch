package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func NewMultiIpEntry(label string, bind binding.StringList) *fyne.Container {
	ipBind := binding.NewString()
	ipEntry := widget.NewEntryWithData(ipBind)
	ipEntry.Validator = validation.NewRegexp(`^(\d{1,3}\.){3}\d{1,3}$`, "Invalid IP address format")
	ipEntry.TextStyle.Monospace = true

	ips := container.NewVBox()
	ipsScroll := container.NewVScroll(ips)
	ipsScroll.SetMinSize(fyne.Size{Width: 200, Height: 150})

	fillIpList(ips, bind)

	bind.AddListener(binding.NewDataListener(func() {
		fillIpList(ips, bind)
	}))

	return container.NewVBox(
		widget.NewLabel(label),
		ipsScroll,
		container.NewBorder(nil, nil, nil,
			widget.NewButton("Add IP", func() {
				ip, _ := ipBind.Get()
				if ipEntry.Validate() == nil {
					bind.Append(ip)
					ips.Add(newIpListItem(ip, bind))
					ipBind.Set("")
				}
			}),
			ipEntry,
		),
	)
}

func newIpListItem(ip string, bind binding.StringList) *fyne.Container {
	label := widget.NewLabel(ip)
	label.TextStyle.Monospace = true

	return container.NewBorder(nil, nil, nil,
		widget.NewButton("-", func() {
			bind.Remove(ip)
		}),
		label,
	)
}

func fillIpList(ips *fyne.Container, bind binding.StringList) {
	currentIps, _ := bind.Get()
	ips.RemoveAll() // Clear the list before repopulating to avoid duplicates
	for ip := range currentIps {
		ips.Add(newIpListItem(currentIps[ip], bind))
	}
}
