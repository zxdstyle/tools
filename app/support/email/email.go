package email

import "gopkg.in/gomail.v2"

func Send() {
	message := gomail.NewMessage()
	message.SetHeader("From", "cxytools@yeah.net")
	message.SetHeader("To", "bob@example.com", "cora@example.com")
	message.SetAddressHeader("Cc", "dan@example.com", "Dan")
	message.SetHeader("Subject", "Hello!")
	message.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//message.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.yeah.net", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
}
