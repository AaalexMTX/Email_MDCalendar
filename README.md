# Email_MDCalendar
定时（月度）邮件推送 当月markdown版，可导入日历表格

## config 配置参数设置
| 参数           | 含义      | 示例            |
| ------------ | ------- |---------------|
| From         | 送件人邮箱号  | xx@163.com    |
| To           | 收件人邮箱号  | yy@163.com    |
| Auth_code    | 授权码     | 生成方式见资料       |
| email_server | 邮件服务器   | 默认stmp.qq.com |
| server_port  | 邮件服务器端口 | 默认 465        |


## 资料

---

### 邮箱的授权码
[QQ邮箱授权码获取](https://service.mail.qq.com/detail/128/53?expand=14)
### 邮件服务器

| 邮箱类型        | SMTP 服务器地址         | SSL 协议端口 | 非 SSL 协议端口 |
| ----------- | ------------------ | -------- | ---------- |
| QQ 邮箱       | smtp.qq. com       | 465/587  | 25         |
| 163 邮箱      | smtp.163. com      | 465/994  | 25         |
| office365邮箱 | smtp.office365.com | 465/587  | 25         |
