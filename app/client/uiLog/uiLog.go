// 12 august 2018

// build OMIT

package uiLog

import (
	"../../models/log"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func UILogPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	group := ui.NewGroup("Журнал событий")
	group.SetMargined(true)

	var logLines = ui.NewMultilineEntry()
	group.SetChild(logLines)

	logLines.SetReadOnly(true)
	log.SetUiLogContainer(logLines)

	vbox.Append(group, true)
	return vbox
}
