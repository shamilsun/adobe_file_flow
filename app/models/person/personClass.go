package person

func (p *IPerson) GetFIO() string {
	return p.LastName + " " + p.FirstName + " " + p.PatronymicName
}
