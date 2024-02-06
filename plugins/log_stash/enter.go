package log_stash

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/utils"
	"server/utils/jwts"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	addr := utils.GetAddr(ip)
	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   addr,
		userId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DEBUG, content)
}
func (l Log) Info(content string) {
	l.send(INFO, content)
}
func (l Log) Warn(content string) {
	l.send(WARN, content)
}
func (l Log) Error(content string) {
	l.send(ERROR, content)
}

func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}
