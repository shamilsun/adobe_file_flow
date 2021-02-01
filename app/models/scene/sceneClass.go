package scene

import (
	"github.com/andlabs/ui"
)

var isClosed = false

func CloseApp() {
	isClosed = true
}
func IsUserClosedApp() bool {
	return isClosed
}

var w IUIMainWindow = IUIMainWindow{}

func SetSessionPresets(v *ui.Group) {
	w.SessionPresets = v
}

func GetSessionPresets() *ui.Group {
	return w.SessionPresets
}

func SetSessionPanel(v *ui.Group) {
	w.SessionPanel = v
}

func GetSessionPanel() *ui.Group {
	return w.SessionPanel
}

func SetSessionStopConfirmation(v *ui.Box) {
	w.SessionStopConfirmationPanel = v
}

func GetSessionStopConfirmation() *ui.Box {
	return w.SessionStopConfirmationPanel
}

func SetSessionStatusPanel(v *ui.Group) {
	w.SessionStatusPanel = v
}

func GetSessionStatusPanel() *ui.Group {
	return w.SessionStatusPanel
}

func SetSessionStatusLabel(v *ui.Label) {
	w.SessionStatusLabel = v
}

func GetSessionStatusLabel() *ui.Label {
	return w.SessionStatusLabel
}

func SetSessionStartNewConfirmation(v *ui.Box) {
	w.SessionStartNewConfirmationPanel = v
}

func GetSessionStartNewConfirmation() *ui.Box {
	return w.SessionStartNewConfirmationPanel
}

func SetSessionPanelsDateTimeCounter(v *ui.Entry) {
	w.SessionPanelsDateTimeCounter = v
}

func GetSessionPanelsDateTimeCounter() *ui.Entry {
	return w.SessionPanelsDateTimeCounter
}

func SessionStopConfirmation(onStop func(), onCancel func()) ui.Control {

	hbox := ui.NewHorizontalBox()
	btnStop := ui.NewButton("Закончить сессию")
	btnCancel := ui.NewButton("Отменить")
	btnStop.OnClicked(func(button *ui.Button) {
		onStop()
	})
	btnCancel.OnClicked(func(button *ui.Button) {
		onCancel()
	})
	hbox.Append(btnStop, false)
	hbox.Append(btnCancel, false)
	return hbox
}

func SessionStartNewConfirmation(onClearNew func(), onStayCurrent func(), onCancel func()) ui.Control {

	hbox := ui.NewHorizontalBox()
	btnClearNew := ui.NewButton("Открыть новую сесиию")
	btnStayCurrent := ui.NewButton("Оставаться в этой сессии (ID, Настройки)")
	btnCancel := ui.NewButton("Отменить")
	btnClearNew.OnClicked(func(button *ui.Button) {
		onClearNew()
	})
	btnStayCurrent.OnClicked(func(button *ui.Button) {
		onStayCurrent()
	})
	btnCancel.OnClicked(func(button *ui.Button) {
		onCancel()
	})
	hbox.Append(btnClearNew, false)
	hbox.Append(btnStayCurrent, false)
	hbox.Append(btnCancel, false)
	return hbox
}
