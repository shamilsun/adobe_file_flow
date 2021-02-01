package person

import "../phone"

type IPerson struct {
	//Uid            uint64
	FirstName      string
	LastName       string
	PatronymicName string
	Passport       string

	Email string
	Phone phone.IPhone
}
