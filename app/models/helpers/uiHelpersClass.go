// 12 august 2018

// build OMIT

package helpers

import "C"
import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)
import "../settings"

func btnDisable(b *ui.Button) {
	if b != nil {
		b.Disable()
	}
}

func btnEnable(b *ui.Button) {
	if b != nil {
		b.Enable()
	}
}

func UIPanelEVCD(input *ui.Box, onConfirmed func(), onEdit func(), onCancel func(), onDelete func()) ui.Control {

	hbox := ui.NewHorizontalBox()

	input.Disable()

	evcPanel := ui.NewVerticalBox()
	evcHorizontal := ui.NewHorizontalBox()

	btnEdit := ui.NewButton("E")
	btnEnable(btnEdit)

	var btnDelete *ui.Button = nil
	if onDelete != nil {
		btnDelete = ui.NewButton("D")
		btnDisable(btnDelete)
		btnDelete.OnClicked(func(button *ui.Button) {
			onDelete()
		})
	}

	btnCancel := ui.NewButton("X")
	btnDisable(btnCancel)

	btnConfirm := ui.NewButton("V")
	btnDisable(btnConfirm)

	btnConfirm.OnClicked(func(button *ui.Button) {
		input.Disable()
		onConfirmed()

		btnEnable(btnEdit)

		btnDisable(btnDelete)
		btnDisable(btnConfirm)
		btnDisable(btnCancel)
	})

	btnCancel.OnClicked(func(button *ui.Button) {
		input.Disable()
		onCancel()

		btnEnable(btnEdit)

		btnDisable(btnDelete)
		btnDisable(btnConfirm)
		btnDisable(btnCancel)
	})

	btnEdit.OnClicked(func(button *ui.Button) {
		input.Enable()
		onEdit()

		btnDisable(btnEdit)

		btnEnable(btnConfirm)
		btnEnable(btnCancel)
		btnEnable(btnDelete)

	})

	hbox.Append(input, true)
	evcHorizontal.Append(btnEdit, false)
	evcHorizontal.Append(btnConfirm, false)
	if btnDelete != nil {
		evcHorizontal.Append(btnDelete, false)
	}
	evcHorizontal.Append(btnCancel, false)

	evcPanel.Append(evcHorizontal, false)
	hbox.Append(evcPanel, false)

	return hbox
}

func UICustomInputField(input *ui.Box, onUpdateModel func(newValue string), getValueModel func() string, getValueFromInput func() string, setValueToInput func(v string)) ui.Control {

	value := getValueModel()
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	input.Disable()

	btnEdit := ui.NewButton("E")
	btnEdit.Show()
	btnCancel := ui.NewButton("X")

	btnConfirm := ui.NewButton("V")
	btnConfirm.Hide()
	btnConfirm.OnClicked(func(button *ui.Button) {
		input.Disable()
		onUpdateModel(getValueFromInput())
		value = getValueFromInput()
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnCancel.Hide()
	btnCancel.OnClicked(func(button *ui.Button) {
		input.Disable()
		setValueToInput(value)
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnEdit.OnClicked(func(button *ui.Button) {
		input.Enable()
		btnEdit.Hide()
		btnConfirm.Show()
		btnCancel.Show()
	})

	hbox.Append(input, true)
	hbox.Append(btnEdit, false)
	hbox.Append(btnConfirm, false)
	hbox.Append(btnCancel, false)

	return hbox
}

func UIGetInputField(v string, onUpdate func(newValue string)) ui.Control {

	value := v
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	input := ui.NewEntry()
	input.SetText(value)
	input.SetReadOnly(true)

	btnEdit := ui.NewButton("E")
	btnEdit.Show()
	btnCancel := ui.NewButton("X")

	btnConfirm := ui.NewButton("V")
	btnConfirm.Hide()
	btnConfirm.OnClicked(func(button *ui.Button) {
		input.SetReadOnly(true)
		onUpdate(input.Text())
		value = input.Text()
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnCancel.Hide()
	btnCancel.OnClicked(func(button *ui.Button) {
		input.SetReadOnly(true)
		input.SetText(value)
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnEdit.OnClicked(func(button *ui.Button) {
		input.SetReadOnly(false)
		btnEdit.Hide()
		btnConfirm.Show()
		btnCancel.Show()
	})

	hbox.Append(input, true)
	hbox.Append(btnEdit, false)
	hbox.Append(btnConfirm, false)
	hbox.Append(btnCancel, false)

	return hbox
}

func UIGetSpinField(v int, min int, max int, onUpdate func(newValue int)) ui.Control {

	value := v
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	spinBox := ui.NewSpinbox(min, max)
	spinBox.SetValue(v)
	spinBox.Disable()

	//spinRange := ui.NewSlider(min, max)
	//spinRange.SetValue(v)
	//spinRange.Disable()

	btnEdit := ui.NewButton("E")
	btnEdit.Show()
	btnCancel := ui.NewButton("X")

	btnConfirm := ui.NewButton("V")
	btnConfirm.Hide()
	btnConfirm.OnClicked(func(button *ui.Button) {
		spinBox.Disable()
		onUpdate(spinBox.Value())
		value = spinBox.Value()
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnCancel.Hide()
	btnCancel.OnClicked(func(button *ui.Button) {
		spinBox.Disable()
		spinBox.SetValue(value)
		btnEdit.Show()
		btnConfirm.Hide()
		btnCancel.Hide()
	})

	btnEdit.OnClicked(func(button *ui.Button) {
		spinBox.Enable()
		btnEdit.Hide()
		btnConfirm.Show()
		btnCancel.Show()
	})

	hbox.Append(spinBox, true)
	hbox.Append(btnEdit, false)
	hbox.Append(btnConfirm, false)
	hbox.Append(btnCancel, false)

	return hbox
}

func UIGetInputFieldForSetting(v settings.ISettings) ui.Control {

	return UIGetInputField(v.GetValue(), func(newValue string) {
		v.UpdateValue(newValue)
	})
}

func UIGetSpinFieldForSetting(v settings.ISettings) ui.Control {
	value, _ := strconv.Atoi(v.GetValue())
	return UIGetSpinField(value, v.MinInt, v.MaxInt, func(newValue int) {
		v.UpdateValue(newValue)
	})
}
