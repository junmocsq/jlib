package jtools

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

func SendSmtpMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "xxx@163.com")
	m.SetHeader("To", "xxx@qq.com")
	m.SetAddressHeader("Cc", "xxxx@126.net", "张君宝")
	m.SetHeader("Subject", "Hello!")
	rand.Seed(time.Now().UnixNano())
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!"+strconv.Itoa(rand.Int()))
	m.Attach("/Users/junmo/Downloads/6 3 copy.JPG")

	//d := gomail.NewDialer("smtp.163.com", 465, "xxx@163.com", "xxx")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 网易邮箱使用代理授权密码
	d := gomail.NewDialer("smtp.163.com", 25, "xxx@163.com", "xxx")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
