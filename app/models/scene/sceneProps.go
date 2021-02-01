package scene

import "github.com/andlabs/ui"

type IUIMainWindow struct {
	SessionPresets                   *ui.Group
	SessionStopConfirmationPanel     *ui.Box
	SessionStartNewConfirmationPanel *ui.Box
	SessionPanel                     *ui.Group
	SessionPanelsDateTimeCounter     *ui.Entry
	SessionStatusPanel               *ui.Group
	SessionStatusLabel               *ui.Label
	SessionsMonitoring               *ui.Group
}
