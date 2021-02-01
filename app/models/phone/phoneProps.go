package phone

type IPhone struct {
	PhoneNumber string
	Messengers  []AMessenger
}

type AMessenger = string

type LMessenger struct {
	Whatsapp AMessenger
	Telegram AMessenger
}

// Enum for public use
var EMessenger = &LMessenger{
	Whatsapp: "Whatsapp",
	Telegram: "Telegram",
}
