// 12 august 2018

// +build OMIT

package home

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

func uiSettingsPage() ui.Control {

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(uiSettingsPathsGroup(), true)
	vbox.Append(uiSettingsCloudGroup(), true)

	return vbox
}

func UiSettingsPathsGroup() {
	return uiSettingsPathsGroup()
}
func uiSettingsPathsGroup() ui.Control {
	group := ui.NewGroup("Настройки папок")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()
	entryFormPaths.Append("Источник (папка с исходниками камеры)", ui.NewEntry(), false)
	entryFormPaths.Append("Хранилище исходных фото (папка) ", ui.NewEntry(), false)
	entryFormPaths.Append("Хранилище обработанных фото (папка)", ui.NewEntry(), false)
	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}

func uiSettingsCloudGroup() ui.Control {
	group := ui.NewGroup("Настройки облака (яндекс.диск)")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()
	entryFormPaths.Append("Логин", ui.NewEntry(), false)
	entryFormPaths.Append("Пароль", ui.NewEntry(), false)
	entryFormPaths.Append("Токен", ui.NewEntry(), false)
	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}

func uiStartNewSessionPage() ui.Control {

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	//vbox.Append(ui.NewLabel("Автопортрет"), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Новая сессия")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Фамилия", ui.NewEntry(), false)
	entryForm.Append("Имя", ui.NewEntry(), false)
	entryForm.Append("Отчество", ui.NewEntry(), false)
	entryForm.Append("Комментарии", ui.NewMultilineEntry(), true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	hbox.Append(ui.NewButton("Начать сессию"), false)

	vbox.Append(hbox, false)

	return vbox
}

func setupUI() {
	mainwin = ui.NewWindow("Авто портрет", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	//mainwin.SetChild( uiStartNewSessionPage())

	tab := ui.NewTab()

	tab.Append("Сессия", uiStartNewSessionPage())
	tab.SetMargined(0, true)

	tab.Append("Настройки", uiSettingsPage())
	tab.SetMargined(1, true)

	//tab.Append("Data Choosers", makeDataChoosersPage())
	//tab.SetMargined(2, true)

	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
func launch() {
	ui.Main(setupUI)
}
