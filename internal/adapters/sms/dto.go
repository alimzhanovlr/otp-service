package sms

type Payload struct {
	Recipient string
	Text      string
	Slug      string
	Source    string
	Type      int
}
