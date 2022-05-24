// package logic

// import (
// 	"fmt"

// 	"gopkg.in/gomail.v2"
// )

// // 发送邮件
// func SendMail() {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "soakcode@qq.com")
// 	m.SetHeader("To", "soakcode@163.com")
// 	m.SetAddressHeader("Cc", "soakcode@qq.com", "蓝兔系统")
// 	m.SetHeader("Subject", "Hello!")
// 	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
// 	// m.Attach("/home/Alex/lolcat.jpg")

// 	d := gomail.NewDialer("smtp.exmail.qq.com", 25, "soakcode@qq.com", "ickffwmknsakjaii")
// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		fmt.Println("错误：", err)
// 		panic(err)
// 	}
// 	// 生成6位随机验证码
// 	// rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	// vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
// 	// // 存到redis中
// 	// redis.Put(core.GetConfiguration().Redis.Modules, email, vcode, 300)
// 	// now := time.Now()
// 	// t := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
// 	// html := fmt.Sprintf(`<div>
// 	// 	<div>
// 	// 		尊敬的%s，您好！
// 	// 	</div>
// 	// 	<div style="padding: 8px 40px 8px 50px;">
// 	// 		<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
// 	// 	</div>
// 	// 	<div>
// 	// 		<p>此邮箱为系统邮箱，请勿回复。</p>
// 	// 	</div>
// 	// </div>`, email, t, vcode)

// 	// m := gomail.NewMessage()
// 	// m.SetAddressHeader("From", bdvoice.Config["email"].(string), "XXX")
// 	// m.SetHeader("To", email)
// 	// m.SetHeader("Subject", "[我的验证码]邮箱验证") //设置邮件主题
// 	// m.SetBody("text/html", html)          //设置邮件正文
// 	// // 第一个参数是host 第三个参数是发送邮箱，第四个参数 是邮箱密码
// 	// d := gomail.NewDialer(smtp.exmail.qq.com, 25, "1234567@qq.com", 123456)
// 	// if err := d.DialAndSend(m); err != nil {
// 	// 	fmt.Println("错误：", err)
// 	// 	return err
// 	// }
// 	// return nil
// }

package logic

import (
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string, aliasName string, attachPath string) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": "soakcode@qq.com",  //发送人邮箱（邮箱以自己的为准）
		"pass": "ickffwmknsakjaii", //发送人邮箱的密码，现在可能会需要邮箱 开启授权密码后在pass填写授权码
		"host": "smtp.qq.com",      //邮箱服务器（此时用的是qq邮箱）
	}

	m := gomail.NewMessage(
		//发送文本时设置编码，防止乱码。 如果txt文本设置了之后还是乱码，那可以将原txt文本在保存时
		//就选择utf-8格式保存
		gomail.SetEncoding(gomail.Base64),
	)
	// 添加别名
	m.SetHeader("From", m.FormatAddress(mailConn["user"], aliasName))
	// 发送给用户(可以多个)
	m.SetHeader("To", mailTo...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	// 设置附件
	if attachPath != "" {
		m.Attach(attachPath)
	}

	//创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、密码，如果端口号为465的话，自动开启SSL，这个时候需要指定TLSConfig
	d := gomail.NewDialer(mailConn["host"], 465, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)

	return err
}
