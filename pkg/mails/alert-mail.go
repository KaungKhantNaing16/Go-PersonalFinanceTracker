package mail

import (
	"log"
)

const (
	MailSenderName     = "Personal Finance Tracker"
	MailSenderAddress  = "scm.kaungkhantnaing@gmail.com"
	MailSenderPassword = "pusmtugnczhasovv"
)

func SendAlertMail() {
	sender := NewMailSender(MailSenderName, MailSenderAddress, MailSenderPassword)
	subject := "သားကြီးမင်းတအားသုံးနေတယ်နော် "
	content := `ပိုက်ဆံသိပ်မကျန်တော့ဘူးနော် ဟျောင့်ရေ`
	to := []string{"kaungkhantnaing168@gmail.com"}

	err := sender.SendMail(subject, content, to, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}
