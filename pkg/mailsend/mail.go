package mailsend

import (
	"fmt"
	"strings"
)

type Mail struct {
	sender 		string
	to 			[]string
	subject		string
	body		string
}

func buildMessage(mail Mail) string {
	msg := "Content-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.to, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.body)

	return msg
}


