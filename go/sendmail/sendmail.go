package main

import (
	"encoding/base64"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

const (
	SmtpUsername     = ""
	SmtpPassword     = ""
	SmtpHostName     = "smtp.gmail.com"
	SmtpHostProtocol = "587"
)

type SmtpData struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Body    string
}

func main() {
	app := cli.NewApp()
	app.Name = "sendmail"
	app.Usage = "Send a mail from command line."
	app.Version = "0.9.0"
	app.Author = "Shintaro Kaneko"
	app.Email = "kaneshin0120@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "subject, s",
			Usage: "Message subject",
		},
		cli.StringFlag{
			Name:  "body, b",
			Usage: "Message body",
		},
		cli.StringFlag{
			Name:  "from, f",
			Usage: "RFC2047",
		},
		cli.StringFlag{
			Name:  "to, t",
			Usage: "RFC2047",
		},
	}
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) {
	fm, fmerr := mail.ParseAddress(c.String("from"))
	to, toerr := mail.ParseAddress(c.String("to"))
	if fmerr != nil {
		log.Fatal(fmerr)
		os.Exit(1)
	}
	if toerr != nil {
		log.Fatal(toerr)
		os.Exit(1)
	}
	data := SmtpData{
		*fm,
		*to,
		c.String("subject"),
		c.String("body"),
	}
	if err := data.sendMail(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (d *SmtpData) sendMail() error {
	auth := smtp.PlainAuth(
		"",
		SmtpUsername,
		SmtpPassword,
		SmtpHostName,
	)

	header := map[string]string{}
	header["From"] = d.From.String()
	header["To"] = d.To.String()
	header["Subject"] = encodeRFC2047(d.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(d.Body))
	err := smtp.SendMail(
		SmtpHostName+":"+SmtpHostProtocol,
		auth,
		d.From.Address,
		[]string{d.To.Address},
		[]byte(message),
	)
	return err
}

func encodeRFC2047(str string) string {
	addr := mail.Address{str, ""}
	return strings.Trim(strings.Trim(addr.String(), " <>"), "\"")
}
