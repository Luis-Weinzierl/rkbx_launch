package widgets

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2/theme"
	ttwidget "github.com/dweymouth/fyne-tooltip/widget"
)

func NewInfoHover(translationId string) *ttwidget.Icon {
	icon := ttwidget.NewIcon(theme.InfoIcon())

	text := globalisation.Get(translationId + "/tooltip")
	icon.SetToolTip(text)

	if text == "" {
		icon.Hide()
	}

	return icon
}
