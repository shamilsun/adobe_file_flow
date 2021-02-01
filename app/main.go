// 12 august 2018

// build OMIT

package main

import "C"
import (
	"./client/setup"
	"./client/uiLog"
	"./models/log"
	"./models/photoSession"
	"./models/scene"	//"./client/preset"
	"./models/preset"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

func mainUI() {
	mainwin = ui.NewWindow("SELF ПОРТРЕТ", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		scene.CloseApp()
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	_uiLog := uiLog.UILogPage()
	tab := ui.NewTab()

	scene.SetSessionPanel(ui.NewGroup("Сессия"))
	photoSession.Render()
	tab.Append("Сессия", scene.GetSessionPanel())
	tab.SetMargined(0, true)

	tab.Append("Опции", setup.UISettingsPage())
	tab.SetMargined(1, true)

	tab.Append("Пресеты", preset.UIPresetsList())
	tab.SetMargined(1, true)

	tab.Append("Мониторинг сессий", photoSession.UIGetPhotoSessionMonitoring())
	tab.SetMargined(1, true)

	hbox.Append(tab, true)
	hbox.Append(_uiLog, false)


	vbox := ui.NewVerticalBox()
	vbox.Append(hbox, true)

	mainwin.SetChild(vbox)
	mainwin.SetMargined(true)

	log.AddLog("Started app")

	mainwin.Show()
}

func main() {
	ui.Main(mainUI)
}
