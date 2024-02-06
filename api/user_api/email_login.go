package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"server/plugins/log_stash"
	"server/utils"
	"server/utils/jwts"
	"server/utils/pwd"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) LoginView(c *gin.Context) {
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 创建日志对象
	log := log_stash.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("用户名不存在，用户名：%s", cr.UserName))

		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误, 用户名: %s, 密码: %s", cr.UserName, cr.Password))

		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		log.Error(fmt.Sprintf("token生成失败, %s", err.Error()))
		res.FailWithMessage("token生成失败", c)
		return
	}

	// 记录ip地址
	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info(fmt.Sprintf("用户登录成功, 用户名: %s, 密码: %s", cr.UserName, cr.Password))
	// 添加记录
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		NickName:  userModel.NickName,
		Token:     token,
		LoginType: ctype.SignEmail,
		IP:        ip,
		Device:    c.Request.UserAgent(),
		Addr:      addr,
	})
	res.OkWithData(token, c)

}
