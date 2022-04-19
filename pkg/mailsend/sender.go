package mailsend

import (
	"fmt"
	"mailSender/pkg/parser"
	"mailSender/сmd/config"
	"net/smtp"
)

func SendError(config config.Config,  to []string, body string) error {
	sender := config.From.Email
	password := config.From.Password
	host := config.SMTP.Host
	port := config.SMTP.Port
	hostAndPort := host + ":" + port


	subject := "Произошла ошибка!"

	request := Mail{
		sender:  sender,
		to:      to,
		subject: subject,
		body:    body,
	}

	message := buildMessage(request)

	auth := smtp.PlainAuth("", sender, password, host)

	err := smtp.SendMail(hostAndPort, auth, sender, to, []byte(message))
	return err
}

func SendCourse(config config.Config,  to []string,  usd parser.Course, eur parser.Course) error{
	sender := config.From.Email
	password := config.From.Password
	host := config.SMTP.Host
	port := config.SMTP.Port
	hostAndPort := host + ":" + port

	subject := "Курс валют"
	body :=  makeCourseBody(usd, eur)

	request := Mail{
		sender:  sender,
		to:      to,
		subject: subject,
		body:    body,
	}

	message := buildMessage(request)

	// Authentication.
	auth := smtp.PlainAuth("", sender, password, host)

	// Sending email.
	err := smtp.SendMail(hostAndPort, auth, sender, to, []byte(message))

	return err
}

//Формирует курс в html формате для красивого вывода в сообщение
func makeCourseBody(usd parser.Course, eur parser.Course) string {
	body := "<p><b>USD КУРС:</b> </p>\r"
	body += fmt.Sprintf("<p><b>Банк:</b> %s\r </p>", usd.Bankname)
	body += fmt.Sprintf("<p><b>USD продажа:</b> %0.2f\r </p>", usd.Sell)
	body += fmt.Sprintf("<p><b>USD покупка:</b> %0.2f\r </p>", usd.Buy)
	body += fmt.Sprintf("<p><b>Дата:</b> %s\r </p>", usd.Date)

	body +=fmt.Sprintf("<p> </p>")

	body += fmt.Sprintf("<p><b>EUR КУРС:</b></p>")
	body += fmt.Sprintf("<p><b>Банк:</b> %s\r</p>", eur.Bankname)
	body += fmt.Sprintf("<p><b>EUR продажа:</b> %0.2f\r</p>", eur.Sell)
	body += fmt.Sprintf("<p><b>EUR покупка:</b> %0.2f\r</p>", eur.Buy)
	body += fmt.Sprintf("<p><b>Дата:</b> %s\r</p>", eur.Date)

	return body
}
