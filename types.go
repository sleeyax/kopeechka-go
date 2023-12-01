package kopeechka

type ResponseType string

const (
	Json ResponseType = "json"
	// Text is not supported yet.
	// Deprecated: text is not supported yet.
	Text ResponseType = "text" // TODO: maybe support text
)

type ResponseStatus string

const (
	Success ResponseStatus = "OK"
	Error   ResponseStatus = "ERROR"
)

type Response struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`

	// Error message in case of an error.
	// Optional extra info in case of success.
	Message string `json:"value"`
}

type BalanceResponse struct {
	Response

	// The balance of your account.
	Balance float64 `json:"balance"`
}

type MailType string

const (
	Yandex  MailType = "YANDEX"
	Outlook MailType = "OUTLOOK"
	MailCom MailType = "MAILCOM"
	MailRu  MailType = "MAILRU"
	Rambler MailType = "RAMBLER"
	Gmx     MailType = "GMX"
	// Mine are all your linked domains.
	Mine MailType = "MINE"
	// All are random mails from all domains.
	All MailType = "ALL"
	// Real are mails from popular sites.
	Real MailType = "REAL"
)

type OrderMailRequest struct {
	// The site you need mail for (required).
	Site string `url:"site"`

	// Specify the type of mail (mail.ru, for example) you want to receive.
	MailType MailType `url:"mail_type,omitempty"`

	// Set to 1 if you need in IMAP access.
	Password int `url:"password,omitempty"`

	// Specify a regular expression to parse information if you don't want the whole email.
	Regex string `url:"regex,omitempty"`

	// Specify if you want to receive a message with a specific subject (one word is enough). It is a priority. If the specified word is not in the subject, the message WILL NOT be issued!
	Subject string `url:"subject,omitempty"`

	// Specify 1 if you need the mails to be taken not from the common base, but from your own bays (via Telegram bot.
	Investor int `url:"investor,omitempty"`

	// Used by developers in the referral program.
	SoftId string `url:"soft,omitempty"`
}

type OrderMailResponse struct {
	Response

	// The order id.
	Id string `json:"id"`

	// The email address.
	Mail string `json:"mail"`

	// The password for the email address.
	// Only returned if you set the password field to `1` in OrderMailRequest.
	Password string `json:"password"`
}

type MessageRequest struct {
	// ID of the activation for which we should receive the message (required).
	OrderId string `url:"id"`

	// Specify 1 if you want to return the full message, not just the link (sometimes needed to save bandwidth).
	Full int `url:"full,omitempty"`
}

type MessageResponse struct {
	Response
	FullMessage string `json:"fullmessage"`
}

type CancelMailRequest struct {
	// ID of the activation which we should cancel (required).
	OrderId string `url:"id"`
}

type CancelMailResponse struct {
	Response
}
