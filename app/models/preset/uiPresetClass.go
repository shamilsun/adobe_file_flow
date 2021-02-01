package preset

import (
	"../helpers"
	"../imageFormat"
	//"../photoSession"
	"github.com/andlabs/ui"
	"log"
	"strings"
)

func (p *IPreset) UIGetSnippet() *ui.Group {
	group := ui.NewGroup("" + p.SubPathName)
	group.SetMargined(false)
	vbox := ui.NewVerticalBox()
	//chk := ui.NewCheckbox("")
	//vbox.Append(chk, false)
	//vbox.Append(ui.NewLabel(p.SubPathName), false)
	//	hbox := ui.NewHorizontalBox()
	//	hbox.Append(ui.NewLabel(""+strings.Join(p.Formats,",")), false)
	//	hbox.Append(ui.NewHorizontalSeparator(),false)
	//	hbox.Append(ui.NewLabel(""+strings.Join(p.PresetSizes,",")), false)
	//vbox.Append(hbox,false)
	vbox.Append(ui.NewLabel(""+strings.Join(p.Formats, ",")), false)
	vbox.Append(ui.NewLabel(""+strings.Join(p.PresetSizes, ",")), false)
	group.SetChild(vbox)
	return group
}

func (pl *IPresetList) UIGetSnippetsList(sessionPresets *IPresetList, onChange func()) *ui.Box {
	vbox := ui.NewVerticalBox()

	if len(pl.List) > 0 {

		for i := 0; i < len(pl.List); i++ {

			p := pl.List[i]

			var panel = func(p *IPreset) *ui.Box {
				c := ui.NewCheckbox("")
				c.SetChecked(PresetExists(sessionPresets, p))
				c.OnToggled(func(checkbox *ui.Checkbox) {
					if checkbox.Checked() {
						sessionPresets.AddPreset(p)
					} else {
						if len(sessionPresets.List) == 1 && PresetExists(sessionPresets, p) {
							checkbox.SetChecked(true)
							return
						}
						log.Println("len", len(sessionPresets.List))
						log.Println(PresetExists(sessionPresets, p))
						sessionPresets.RemovePreset(p)
					}
					onChange()
				})

				hbox := ui.NewHorizontalBox()
				hbox.Append(c, false)
				hbox.Append(p.UIGetSnippet(), true)

				return hbox

			}(p)

			vbox.Append(panel, false)
			//vbox.Append(ui.NewVerticalSeparator(),false)
		}
	}
	return vbox
}

var (
	panelOfPresets *ui.Box
)

func uiPresetBlock(p *IPreset) ui.Control {

	panel := ui.NewHorizontalBox()

	group := ui.NewGroup("Пресет")
	group.SetMargined(true)

	vbox := ui.NewVerticalBox()
	//vbox.Append(entryFormPaths, false)
	//btnDeletePreset := ui.NewButton("Удалить пресет")

	//vbox.Append(btnDeletePreset, false)
	vbox.Append(uiPresetComment(p), false)
	vbox.Append(uiChoosePresetFolder(p), false)
	vbox.Append(uiPresetActionSet(p), false)

	hbox := ui.NewHorizontalBox()

	hbox.Append(uiFileFormat(p), false)
	hbox.Append(uiSize(p), false)

	vbox.Append(hbox, false)

	group.SetChild(vbox)

	//btnDeletePreset.OnClicked(func(button *ui.Button) {
	//	panel.Hide()
	//})

	panel.Append(group, true)

	operationPanel := helpers.UIPanelEVCD(panel,
		func() {
			log.Println("update: ", p)
			p.Save()
		},
		func() {

		},
		func() {

		},
		func() {
			//	operationPanel.Hide()
		},
	)

	//btnDeletePreset.OnClicked(func(button *ui.Button) {
	//	operationPanel.Hide()
	//})

	return operationPanel
}

func uiFileFormat(p *IPreset) ui.Control {
	group := ui.NewGroup("Формат файла")
	group.SetMargined(true)
	hbox := ui.NewHorizontalBox()

	var list []imageFormat.AImageFormat = []imageFormat.AImageFormat{
		imageFormat.EImageFormat.JPG,
		imageFormat.EImageFormat.PNG,
		imageFormat.EImageFormat.PSD,
	}

	radio := ui.NewRadioButtons()
	if len(p.Formats) > 1 {
		p.Formats = nil
	}
	for i, v := range list {
		radio.Append(v)
		if helpers.ItemExists(p.Formats, v) {
			radio.SetSelected(i)
		}
	}
	radio.OnSelected(func(buttons *ui.RadioButtons) {
		p.Formats = nil
		p.AddFormat(list[buttons.Selected()])
	})
	hbox.Append(radio, true)
	//for _, v := range list {
	//
	//	var chkBtn = func(v string) *ui.Checkbox {
	//		r := ui.NewCheckbox(v)
	//		r.SetChecked(helpers.ItemExists(p.Formats, v))
	//		r.OnToggled(func(checkbox *ui.Checkbox) {
	//			log.Println("toggled")
	//			if checkbox.Checked() {
	//				p.AddFormat(v)
	//			} else {
	//				p.RemoveFormat(v)
	//			}
	//		})
	//		return r
	//	}(v)
	//
	//	hbox.Append(chkBtn, false)
	//}

	group.SetChild(hbox)
	return group
}

func uiSize(p *IPreset) ui.Control {
	group := ui.NewGroup("Размер")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	var list = []APresetSize{
		EPresetSize.Big,
		EPresetSize.Small,
	}

	for _, v := range list {

		var chkBtn = func(v string) *ui.Checkbox {
			r := ui.NewCheckbox(v)
			r.SetChecked(helpers.ItemExists(p.PresetSizes, v))
			r.OnToggled(func(checkbox *ui.Checkbox) {
				if checkbox.Checked() {
					p.AddSize(v)
				} else {
					p.RemoveSize(v)
				}
			})
			return r
		}(v)

		hbox.Append(chkBtn, false)
	}

	group.SetChild(hbox)
	return group
}

func uiChoosePresetFolder(p *IPreset) ui.Control {
	group := ui.NewGroup("Выбор пресета")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	input := ui.NewEntry()
	input.SetText(p.Filename)

	input.OnChanged(func(entry *ui.Entry) {
		p.Filename = entry.Text()
	})

	hbox.Append(input, true)
	//hbox.Append(ui.NewButton("..."), false)
	group.SetChild(hbox)
	return group
}

func uiPresetComment(p *IPreset) ui.Control {
	group := ui.NewGroup("Наименование подпапки пресета")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	input := ui.NewEntry()
	input.SetText(p.SubPathName)
	input.OnChanged(func(entry *ui.Entry) {
		p.SubPathName = entry.Text()
	})
	hbox.Append(input, true)
	group.SetChild(hbox)
	return group
}

func uiPresetActionSet(p *IPreset) ui.Control {
	group := ui.NewGroup("Экшены")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	actionName := ui.NewEntry()
	actionName.SetText(p.ActionName)
	actionName.OnChanged(func(entry *ui.Entry) {
		p.ActionName = entry.Text()
	})
	actionSet := ui.NewEntry()
	actionSet.SetText(p.ActionSet)
	actionSet.OnChanged(func(entry *ui.Entry) {
		p.ActionSet = entry.Text()
	})
	hbox.Append(actionName, true)
	hbox.Append(actionSet, true)
	group.SetChild(hbox)
	return group
}
func uiPresetAction() ui.Control {
	hbox := ui.NewHorizontalBox()

	btnAddNewPreset := ui.NewButton("Добавить пресет")
	btnAddNewPreset.OnClicked(func(button *ui.Button) {
		p := IPreset{
			Id: 0,
		}
		getPanelOfPresets().Append(uiPresetBlock(&p), false)
	})

	hbox.Append(btnAddNewPreset, false)
	return hbox
}

func setPanelOfPresets(panel *ui.Box) {
	panelOfPresets = panel
}
func getPanelOfPresets() *ui.Box {
	return panelOfPresets
}
func UIPresetsList() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	setPanelOfPresets(ui.NewVerticalBox())

	presets := GetProjectPresets()

	for k, v := range presets {
		log.Println(k, v)
		getPanelOfPresets().Append(uiPresetBlock(v), false)
	}

	vbox.Append(uiPresetAction(), false)
	vbox.Append(getPanelOfPresets(), false)

	return vbox
}
