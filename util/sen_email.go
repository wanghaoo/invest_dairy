package util

import (
	"encoding/base64"
	"fmt"
	"invest_dairy/common"
	"net/mail"
	"net/smtp"
	"strconv"
)

func SendEmail(email, name, orderId, amount string) error {
	amountInt, _ := strconv.ParseInt(amount, 10, 64)
	amount = strconv.FormatInt(amountInt/100, 10)
	//发送
	b64 := base64.NewEncoding(common.Base64NewEncord)
	from := mail.Address{"发送人", common.CONTA_PADRAO}
	to := mail.Address{"接收人", email}
	host := common.SERVIDOR_SMTP
	password := common.SENHA_CONTA_PADRAO
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("邮件标题2")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	body := "<p>Hi there,</p>"
	body += "<p>Your payment [" + orderId + "]has been confirmed. Please see the order summary below：<p>"
	body += "<p><B>Order Summary</B></p>"
	body += "<p>Your Name: " + name + " </p>"
	body += "<p>Order id: " + orderId + "</p>"
	body += "<p>Amount NGN: " + amount + "</p>"
	body += "<p><B>Service Fee Description</B></p>"
	body += "<p>Our platform is a bigdata and AI based loan recommend App. Based on your personal information and credit information, we use machine learning technology to recommend loan product. We have cooperated with most of loan services in Nigeria, which will enhance your loan success probability.</p>"
	body += "<p>In order to get loan recommendation, we require you to pay a mount of money for service. The service fee is not relevant to the result of loan for you on our cooperated platform. The service fee is for our recommendation service and membership service, which we can recommend more loan for you in the future.</p>"
	body += "<p>The service fee is a one-time consumption fee and is non-refundable. Hope to get your understanding. If you have any questions, please go to the APP Personal Center and contact us through E-mail. We will answer your questions timely. Looking forward to your use of our products and services again.</p>"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := smtp.SendMail(
		host+":25",
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	return err
}
