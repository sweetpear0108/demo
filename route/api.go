package route

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var vcode string

func Verify(c *gin.Context) {
	email := c.PostForm("email")
	var mailConf model.MailboxConf
	mailConf.Title = "注册验证"
	mailConf.RecipientList = []string{email}
	mailConf.Sender = `2720102562@qq.com`
	mailConf.SPassword = "exkpksftmruedchh"
	mailConf.SMTPAddr = `smtp.qq.com`
	mailConf.SMTPPort = 25

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	m.SetHeader(`From`, mailConf.Sender, "mxplayer")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Send Email Failed",
		})
		return
	}
	log.Printf("Send Email Success")
	c.IndentedJSON(http.StatusOK, email)
}

func RegisterRequest(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	pwd := c.PostForm("pwd")
	if strings.Compare(code, vcode) != 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Verification code wrong",
		})
		return
	}
	name, id, err := service.Reg(email, pwd)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Server is busy",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func GetUser(c *gin.Context) {
	reqAccountId := c.Param("id")
	id, err := strconv.Atoi(reqAccountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "the id must be integer",
		})
		return
	}
	user, err := service.QueryById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "user doesn't exist",
		})
		return

	}
	c.IndentedJSON(http.StatusOK, user)

}
