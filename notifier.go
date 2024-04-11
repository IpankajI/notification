package main

import "fmt"

type Emailer interface {
	SendEmail(email string, subject, body string) error
}

type mockEmailer struct{}

func (me *mockEmailer) SendEmail(email string, subject, body string) error {

	fmt.Println("sending email to- ", email, "subject: ", subject, "body: ", body)

	return nil
}

func NewEmailService() Emailer {
	return &mockEmailer{}
}
