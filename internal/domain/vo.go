package domain

import (
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

const (
	LanguageRu Language = "ru"
	LanguageKk Language = "kk"
	LanguageEn Language = "en"

	ChannelSMS   Channel = "sms"
	ChannelEmail Channel = "email"
	ChannelPush  Channel = "push"
)

var (
	_channels  = map[Channel]struct{}{ChannelSMS: {}, ChannelEmail: {}, ChannelPush: {}}
	_languages = map[Language]struct{}{LanguageRu: {}, LanguageKk: {}, LanguageEn: {}}
)

type TTLSeconds int
type Language string

func (l Language) Valid() bool {
	_, ok := _languages[l]
	return ok
}

func NewLanguage(s string) Language {
	lang := Language(s)
	if !lang.Valid() {
		return LanguageRu
	}
	return lang
}

type Status string

const (
	StatusSent      Status = "sent"
	StatusVerified  Status = "verified"
	StatusExpired   Status = "expired"
	StatusFailed    Status = "failed"
	StatusThrottled Status = "throttled"
	StatusCancelled Status = "cancelled"
)

type Channel string

func (c Channel) IsValid() bool {
	_, ok := _channels[c]
	return ok
}

type Target interface {
	Channel() Channel
	Value() string
}

func NewTarget(ch Channel, raw string) (Target, error) {
	switch ch {
	case ChannelSMS:
		return NewPhoneTarget(raw)
	case ChannelEmail:
		return NewEmailTarget(raw)
	case ChannelPush:
		id, err := strconv.Atoi(raw)
		if err != nil {
			return nil, ErrInvalidCustomerID
		}
		return NewCustomerTarget(id)
	default:
		return nil, ErrUnknownChannel
	}
}

type phoneTarget string

func (t phoneTarget) Channel() Channel { return ChannelSMS }
func (t phoneTarget) Value() string    { return string(t) }

func NewPhoneTarget(raw string) (Target, error) {
	// примитивная валидация
	re := regexp.MustCompile(`^\+?\d{10,15}$`)
	if !re.MatchString(raw) {
		return nil, ErrInvalidPhone
	}
	return phoneTarget(raw), nil
}

type emailTarget string

func (t emailTarget) Channel() Channel { return ChannelEmail }
func (t emailTarget) Value() string    { return string(t) }

func NewEmailTarget(raw string) (Target, error) {
	if _, err := mail.ParseAddress(raw); err != nil {
		return nil, ErrInvalidEmail
	}
	return emailTarget(raw), nil
}

type customerTarget string

func (t customerTarget) Channel() Channel { return ChannelPush }
func (t customerTarget) Value() string    { return string(t) }

func NewCustomerTarget(id int) (Target, error) {
	if id <= 0 {
		return nil, ErrInvalidCustomerID
	}
	return customerTarget(strconv.Itoa(id)), nil
}

type OTPContext struct {
	Action string
	Event  string
}

func (o OTPContext) Build() string {
	return fmt.Sprintf("%s.%s", o.Action, o.Event)
}

func NewOTPContext(raw string) (OTPContext, error) {
	var snake = regexp.MustCompile(`^[a-z]+(_[a-z]+)*$`)
	parts := strings.Split(raw, ".")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return OTPContext{}, ErrInvalidOTPContext
	}
	if !snake.MatchString(parts[0]) || !snake.MatchString(parts[1]) {
		return OTPContext{}, ErrInvalidOTPContext
	}
	return OTPContext{Action: parts[0], Event: parts[1]}, nil
}
