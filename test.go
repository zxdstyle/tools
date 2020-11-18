package main

import "tools/app/support/email"

func main() {
	letter := &email.Letter{
		From:     "cxytools@foxmail.com",
		FromName: "",
		To:       "cxytools@yeah.net",
		ToName:   "",
		Cc:       "",
		CcName:   "",
		Subject:  "test",
		Body:     "test",
	}
	err := letter.Send()
	if err != nil {
		panic(err)
	}
}
