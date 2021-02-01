package person

import (
	"../email"
	"../phone"
	"github.com/andlabs/ui"
)

func uiFIO(p *IPerson) ui.Control {
	entryForm := ui.NewForm()
	entryForm.SetPadded(true)

	inputFirstName := ui.NewEntry()
	inputLastName := ui.NewEntry()
	inputPatronymic := ui.NewEntry()

	inputFirstName.SetText(p.FirstName)
	inputLastName.SetText(p.LastName)
	inputPatronymic.SetText(p.PatronymicName)

	inputFirstName.OnChanged(func(entry *ui.Entry) {
		p.FirstName = entry.Text()
	})

	inputLastName.OnChanged(func(entry *ui.Entry) {
		p.LastName = entry.Text()
	})

	inputPatronymic.OnChanged(func(entry *ui.Entry) {
		p.PatronymicName = entry.Text()
	})

	entryForm.Append("Фамилия", inputLastName, false)
	entryForm.Append("Имя", inputFirstName, false)
	entryForm.Append("Отчество", inputPatronymic, false)
	return entryForm
}

func uiContacts(p *IPerson) ui.Control {
	vbox := ui.NewVerticalBox()
	entryForm := ui.NewForm()
	entryForm.SetPadded(true)

	inputPhone := ui.NewEntry()

	inputPhone.SetText(p.Phone.PhoneNumber)
	inputPhone.OnChanged(func(entry *ui.Entry) {
		p.Phone.PhoneNumber = entry.Text()
	})

	//
	entryForm.Append("Телефон", inputPhone, false)
	entryForm.Append("", phone.UIMessengers(&p.Phone), false)

	email.AddEmailControlBox(p.Email, vbox, func(v string) {
		p.Email = v
	})

	entryForm.Append("E-mail", vbox, false)

	return entryForm
}

func uiPassport(p *IPerson) ui.Control {
	entryForm := ui.NewForm()
	entryForm.SetPadded(true)

	inputPassport := ui.NewMultilineEntry()

	inputPassport.SetText(p.Passport)

	inputPassport.OnChanged(func(entry *ui.MultilineEntry) {
		p.Passport = entry.Text()
	})

	entryForm.Append("Паспортные данные", inputPassport, true)

	return entryForm
}

func UIPerson(p *IPerson) ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(uiFIO(p), false)
	vbox.Append(ui.NewVerticalSeparator(), false)
	vbox.Append(uiContacts(p), false)
	vbox.Append(ui.NewVerticalSeparator(), false)
	vbox.Append(uiPassport(p), true)

	return vbox
}
