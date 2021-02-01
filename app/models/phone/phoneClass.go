package phone

import "log"
import "../helpers"

func (p *IPhone) AddMessenger(f AMessenger) {
	if !helpers.ItemExists(p.Messengers, f) {
		p.Messengers = append(p.Messengers, f)
	}
	log.Println(p.Messengers)
}

func (p *IPhone) RemoveMessanger(f AMessenger) {
	if helpers.ItemExists(p.Messengers, f) {
		p.Messengers = helpers.RemoveItem(p.Messengers, f)
	}
}
