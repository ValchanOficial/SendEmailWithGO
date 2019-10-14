package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	sendMail("...@gmail.com", []string{"bar@foo.com", "foo@bar.com"}, "Hello World")
	fmt.Println("Email successfully sent!")
}

func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func sendMail(from string, to []string, body string) {
	//configuração
	servername := "smtp.gmail.com:465"                  //servidor SMTP e PORTA
	host := "smtp.gmail.com"                            //host
	pass := "UmaCalopsitaEsteveAqui"                    //senha
	auth := smtp.PlainAuth("Valchan", from, pass, host) //autenticação
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	toHeader := strings.Join(to, ",")
	msg := "From: " + from + "\n" + "To: " + toHeader + "\n" + "Subject: Hello World\n\n" + body

	//conecta com o servidor SMTP
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	//retorna client SMTP
	c, err := smtp.NewClient(conn, host)
	checkErr(err)

	//autentica
	err = c.Auth(auth)
	checkErr(err)

	//adiciona remetente
	err = c.Mail(from)
	checkErr(err)

	//adiciona destinatários
	for _, addr := range to {
		err = c.Rcpt(addr)
		checkErr(err)
	}

	//prepara corpo do email
	w, err := c.Data()
	checkErr(err)

	//adiciona corpo do e-mail
	_, err = w.Write([]byte(msg))
	checkErr(err)

	//fecha corpo do e-mail
	err = w.Close()
	checkErr(err)

	//encerra conexão
	c.Quit()
}
