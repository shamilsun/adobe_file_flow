package phone

import (
	"../helpers"
	"github.com/andlabs/ui"
)

func UIMessengers(p *IPhone) ui.Control {
	group := ui.NewGroup("Мессенджеры")
	group.SetMargined(false)
	hbox := ui.NewHorizontalBox()

	var list = []AMessenger{
		EMessenger.Whatsapp,
		EMessenger.Telegram,
	}

	for _, v := range list {

		var chkBtn = func(v string) *ui.Checkbox {
			r := ui.NewCheckbox(v)
			r.SetChecked(helpers.ItemExists(p.Messengers, v))
			r.OnToggled(func(checkbox *ui.Checkbox) {
				if checkbox.Checked() {
					p.AddMessenger(v)
				} else {
					p.RemoveMessanger(v)
				}
			})
			return r
		}(v)

		hbox.Append(chkBtn, false)
	}

	group.SetChild(hbox)
	return group
}
