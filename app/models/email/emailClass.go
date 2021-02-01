package email

import (
	"../dotenv"
	"fmt"
	"github.com/adityaxdiwakar/go-sendpulse"
	"github.com/quickemailverification/quickemailverification-go"
	"net"
	"regexp"
	"strings"
)

var qev *quickemailverification.Client

func getQuickEmailVerificationInstance() *quickemailverification.Client {
	if qev == nil {
		qev = quickemailverification.CreateClient(dotenv.GetEnv().QuickEmailVerification)
	}
	return qev
}
func GetQuickEmailVerification(v string) string {
	response, err := getQuickEmailVerificationInstance().Verify(v)
	if err != nil {
		panic(err)
	}
	print(response.Result)

	return response.Result
}
func SendYandexDiskMailNotify(toEmail string, toName string, publicUrl string) error {
	html := []byte(fmt.Sprintf("<strong>Ваша ссылка %s</strong>", publicUrl)) //fmt.Sprintf("foo: %s", bar)
	text := []byte(fmt.Sprintf("Ваша ссылка %s", publicUrl))

	recipients := []sendpulse.Recipient{
		sendpulse.Recipient{
			Name:  toName,
			Email: toEmail,
		},
	}
	subject := "Ссылка на материалы сессии SELF ПОРТРЕТ"

	sendpulse.Initialize(
		dotenv.GetEnv().ClientId,     //"51f13e76e28e20dbfbb57093e768cfb4",
		dotenv.GetEnv().ClientSecret, //"1057dbf1256cfc03db29150c51a16cd6",
		dotenv.GetEnv().FromName,     //"Self ПОРТРЕТ",
		dotenv.GetEnv().FromEmail,
		//"selfportrait.pro@yandex.ru",
		//"bot@selfportrait.pro",
	)
	err := sendpulse.SendEmail(
		html,
		text,
		subject,
		recipients,
	)
	return err
}

func IsEmailValid(e string, checkDomains bool) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	if checkDomains {
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			return false
		}
	}
	return true
}
