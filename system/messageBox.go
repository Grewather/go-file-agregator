package system

import (
	"github.com/ncruces/zenity"
)

func ShowMessageBox(err string) {
	zenity.Error(err,
		zenity.Title("Error"),
		zenity.ErrorIcon)

		
}
