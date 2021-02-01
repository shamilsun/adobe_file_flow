package email

import (
	"github.com/andlabs/ui"
	"log"
)

func AddEmailControlBox(v string, box *ui.Box, onChange func(v string)) {
	inputEmail := ui.NewEntry()

	inputEmail.SetText(v)

	textArea := ui.NewLabel("")
	textArea.SetText("Не указан")
	textArea.Hide()

	checkByQEV := ui.NewButton("Проверить через QEV")
	checkByQEV.Hide()

	var isValid bool = false

	var triggerIsValid = func(newValue bool) {

		if newValue == isValid {
			return
		}
		log.Println("xxx1", newValue)
		isValid = newValue

		if isValid {
			checkByQEV.Show()
			textArea.Hide()
		} else {
			textArea.SetText("Редактируется")
			checkByQEV.Hide()
			textArea.Show()
		}
	}

	inputEmail.OnChanged(func(entry *ui.Entry) {
		//p.Email = entry.Text()
		v := entry.Text()
		go onChange(v)
		go triggerIsValid(IsEmailValid(v, false))
	})

	checkByQEV.OnClicked(func(button *ui.Button) {
		result := GetQuickEmailVerification(inputEmail.Text())
		if result == "valid" {
			textArea.SetText("Валидный емайл")
			textArea.Show()
			checkByQEV.Hide()
			return
		}
		if result == "invalid" {
			textArea.SetText("Емайл не сущетвует")
			textArea.Show()
			checkByQEV.Hide()
			return
		}
		textArea.SetText("Проверить не удалось")
		textArea.Show()
		checkByQEV.Hide()
	})

	box.Append(inputEmail, false)
	box.Append(textArea, false)
	box.Append(checkByQEV, false)
}
