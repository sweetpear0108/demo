package model

type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码
	SPassword string
	// SMTP 服务器地址
	SMTPAddr string
	// SMTP端口
	SMTPPort int
}
