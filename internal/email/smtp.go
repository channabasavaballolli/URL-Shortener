package email

import (
	"bytes" //because after rendering and writing into memory buffer it becomes html string
	"html/template"
	"net/smtp" //for implementing smtp functionality
	"os"
)

type SMTPConfig struct { //helpful for how will we connect to SMTP
	Host     string
	Port     string
	Username string
	Password string
}

func NewSMTPConfig() SMTPConfig {
	return SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"), //loads envi variable in go struct
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASS"),
	}
}

type Email struct { //represents what data should be included in mail
	To          string
	Subject     string
	UserName    string
	OriginalURL string
	ShortURL    string
	Alias       string
	CreatedAt   string
	ExpiresAt   string
}

func RenderTemplate(email Email) (string, error) { //converts template+data into html
	tmpl, err := template.ParseFiles(
		"internal/email/templates/url_created.html",
	) //reads html file
	if err != nil {
		return "", err
	}
	var body bytes.Buffer            //template engine writee html into memory buffer
	err = tmpl.Execute(&body, email) //injects into template placeholders
	return body.String(), nil

}
func BuildMessage(from string, email Email) (string, error) { //to construct Raw smtp message

	htmlBody, err := RenderTemplate(email)

	if err != nil {
		return "", err
	}

	message := "From: " + from + "\r\n" +
		"To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"MIME-Version: 1.0\r\n" + //tells email client that this uses MIME formatting and tells render body as html
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" //header sction and \r\n\r\n means headers end here

	message += htmlBody
	// message += "Hello, " + email.UserName + "\r\n\r\n" //CRLF (Carriage Return + Line Feed) beacause SMTP protocol officially requires CRLF line endings.
	// message += "Your short URL was successfully created.\r\n\r\n"

	// message += "Original URL:\r\n"
	// message += email.OriginalURL + "\r\n\r\n"

	// message += "Short URL:\r\n"
	// message += email.ShortURL + "\r\n\r\n"

	// if email.Alias != "" {
	// 	message += "Custom Alias:\r\n"
	// 	message += email.Alias + "\r\n\r\n"
	// }

	// message += "Created At:\r\n"
	// message += email.CreatedAt + "\r\n\r\n"

	// message += "Expires At:\r\n"
	// message += email.ExpiresAt + "\r\n\r\n"

	// message += "Thank you for using our URL Shortener."

	// return message
	return message, nil
}
func SendEmail(config SMTPConfig, email Email) error { //this function sends the already ready mail (actual SMTP communication)
	auth := smtp.PlainAuth( //prepares SMTP credentials for Gmail like username and password
		"",
		config.Username,
		config.Password,
		config.Host,
	)
	message, err := BuildMessage(config.Username, email) //generates headers a string
	if err != nil {
		return err
	}
	smtpAddr := config.Host + ":" + config.Port //gets smtp.gmail.com:587
	err = smtp.SendMail(
		smtpAddr,           //target smtp server address
		auth,               //authentication credentials
		config.Username,    //sender address
		[]string{email.To}, //reciever slice if multiple recepients
		[]byte(message),    //smtp transmits bytes over TCP sockets not full strings
	)

	return err

}
