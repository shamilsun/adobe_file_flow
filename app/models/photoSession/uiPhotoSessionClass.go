package photoSession

import (
	"../../models/person"
	"../../models/preset"
	"../../models/scene"
	"github.com/andlabs/ui"
	"log"
	"time"
)

var (
	currentSession *IPhotosSession = getLastRow()
)

func uiCurrentDate() *ui.Group {
	group := ui.NewGroup("Дата")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	timeLabel := ui.NewEntry()
	timeLabel.SetReadOnly(true)

	hbox.Append(timeLabel, true)
	group.SetChild(hbox)

	scene.SetSessionPanelsDateTimeCounter(timeLabel)

	return group
}

func uiSessionDate() ui.Control {
	group := ui.NewGroup("Интервал сессии от, до")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()
	//hbox.Append(ui.NewEntry(), true)
	//hbox.Append(ui.NewVerticalSeparator(),false)
	cbox := ui.NewCombobox()
	cbox.Append("1 час")
	cbox.Append("2 часа")
	cbox.Append("3 часа")
	hbox.Append(cbox, true)
	group.SetChild(hbox)
	return group
}

func uiAdobeQueue(s *IPhotosSession) ui.Control {
	group := ui.NewGroup("На обработке")
	group.SetMargined(true)
	//s.uiLblAdobeQueue = ui.NewLabel("10")
	group.SetChild(ui.NewLabel(string(s.GetAdobeQueueCount())))
	return group
}

func uiFileFromCamera() ui.Control {
	group := ui.NewGroup("Получено с камеры")
	group.SetMargined(true)
	group.SetChild(ui.NewLabel("10"))
	return group
}

func uiFileToCC() ui.Control {
	group := ui.NewGroup("Передано на обработку")
	group.SetMargined(true)
	group.SetChild(ui.NewLabel("10"))
	return group
}

func uiFileFromCC() ui.Control {
	group := ui.NewGroup("Получено с обработки")
	group.SetMargined(true)
	group.SetChild(ui.NewLabel("10"))
	return group
}

func uiFileToCloud() ui.Control {
	group := ui.NewGroup("Передано в облако")
	group.SetMargined(true)
	group.SetChild(ui.NewLabel("10"))
	return group
}

func uiSessionStatus(s *IPhotosSession) ui.Control {
	group := ui.NewGroup("Статус сессии")
	group.SetMargined(true)
	hbox := ui.NewHorizontalBox()

	hbox.Append(uiAdobeQueue(s), false)
	hbox.Append(ui.NewHorizontalSeparator(), false)

	group.SetChild(hbox)
	return group
}

func (s *IPhotosSession) uiUpdatePresetsPanel() {
	list := preset.IPresetList{
		List: preset.GetProjectPresets(),
	}

	vbox := ui.NewVerticalBox()

	btn := ui.NewButton("Обновить")
	btn.OnClicked(func(button *ui.Button) {
		s.uiUpdatePresetsPanel()
	})

	vbox.Append(btn, false)
	vbox.Append(list.UIGetSnippetsList(&s.PresetsList, func() {
		s.Save()
	}), true)

	scene.GetSessionPresets().SetChild(vbox)
}

func (s *IPhotosSession) UISessionPanel() ui.Control {

	panel := ui.NewVerticalBox()

	sessionEntityPanel := ui.NewHorizontalBox()

	vbox := ui.NewVerticalBox()

	vbox.Append(uiCurrentDate(), false)

	//	vbox.Append(uiSessionDate(), false)

	vbox.Append(person.UIPerson(&s.Person), true)

	sessionEntityPanel.Append(vbox, true)

	scene.SetSessionPresets(ui.NewGroup("Выберите пресеты"))

	s.uiUpdatePresetsPanel()

	sessionEntityPanel.Append(scene.GetSessionPresets(), true)

	//scene.SetSessionStatusPanel(ui.NewGroup("Мониторинг сессий"))
	//scene.GetSessionStatusPanel().Hide()

	btnNewSession := ui.NewButton("Начать фото сессию")
	btnStopSession := ui.NewButton("Остановить фото сессию")

	closeStatusPane := func() {
		scene.GetSessionStartNewConfirmation().Hide()
		scene.GetSessionStopConfirmation().Hide()
		//statusPanel.Hide()
	}
	showStatusPane := func() {
		sessionEntityPanel.Enable()
		btnStopSession.Hide()
		btnNewSession.Show()
	}

	scene.SetSessionStartNewConfirmation(ui.NewHorizontalBox())
	scene.GetSessionStartNewConfirmation().Append(scene.SessionStartNewConfirmation(func() {
		log.Println("new cleared")
		closeStatusPane()
		currentSession.SetScenario(ESessionScenario.Finished)
		//currentSession.Stop()
		newSession := IPhotosSession{
			PresetsList: currentSession.PresetsList,
		}
		currentSession = &newSession
		currentSession.Save()
		showStatusPane()
		scene.GetSessionPanel().SetChild(currentSession.UISessionPanel())
	}, func() {
		log.Println("new but stay on same")
		closeStatusPane()
		//currentSession.Stop()
		currentSession.SetScenario(ESessionScenario.Finished)
		showStatusPane()
	}, func() {
		log.Println("cancel")
		scene.GetSessionStartNewConfirmation().Hide()
	},
	), false)
	scene.GetSessionStartNewConfirmation().Hide()
	scene.SetSessionStopConfirmation(ui.NewHorizontalBox())
	scene.GetSessionStopConfirmation().Append(scene.SessionStopConfirmation(func() {

		scene.GetSessionStartNewConfirmation().Show()

	}, func() {
		scene.GetSessionStartNewConfirmation().Hide()
		scene.GetSessionStopConfirmation().Hide()
	}), false)
	scene.GetSessionStopConfirmation().Hide()

	sessionActionBottomPanel := ui.NewHorizontalBox()
	sessionActionBottomPanel.SetPadded(true)

	fnNewSession := func() {
		currentSession.SetScenario(ESessionScenario.InProgress)
		//	currentSession.Save()
		sessionEntityPanel.Disable()
		btnNewSession.Hide()
		btnStopSession.Show()
		//	statusPanel.Show()

	}
	btnNewSession.OnClicked(func(button *ui.Button) {
		fnNewSession()
	})
	sessionActionBottomPanel.Append(btnNewSession, false)

	fnStopSession := func() {
		scene.GetSessionStopConfirmation().Show()
	}
	btnStopSession.Hide()
	btnStopSession.OnClicked(func(button *ui.Button) {
		fnStopSession()
	})

	sessionActionBottomPanel.Append(btnStopSession, false)

	panel.Append(sessionEntityPanel, true)

	panel.Append(sessionActionBottomPanel, false)
	panel.Append(scene.GetSessionStopConfirmation(), false)
	panel.Append(scene.GetSessionStartNewConfirmation(), false)

	if s.Scenario == ESessionScenario.InProgress {
		fnNewSession()
	}

	return panel
}

func (s *IPhotosSession) startScenario() {
	if s == nil {
		return
	}

	log.Println("start scenario")
	log.Println(s)
	log.Println(s.Scenario)

	scene.GetSessionPanel().SetChild(s.UISessionPanel())

}

func Render() {
	currentSession.startScenario()
}

func StartSessionsMonitor() {
	time.Sleep(5 * time.Second)
	for {

		if scene.IsUserClosedApp() {
			return
		}

		if scene.GetSessionStatusPanel() == nil {
			continue
		}

		var list []*IPhotosSession = nil

		for _, v := range sessionsList {
			if v.Scenario == ESessionScenario.Stopping || v.Scenario == ESessionScenario.InProgress || v.GetAdobeQueueCount() > 0 {
				list = append(list, v)
			}
		}

		if !scene.GetSessionStatusPanel().Visible() {
			scene.GetSessionStatusPanel().Show()
		}
		updateTable(list)
		for getUpdateTableTimeout() > 0 {
			time.Sleep(time.Second)
			decUpdateTableTimeout()

			if scene.IsUserClosedApp() {
				return
			}
		}

	}
}
