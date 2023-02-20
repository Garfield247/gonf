package controller

import (
	"github.com/Garfield247/gonf.git/pkg/e"
	"github.com/Garfield247/gonf.git/pkg/response"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `valid:"Required;MaxSize(20)" json:"username"`
	Password string `valid:"Required;MaxSize(20)" json:"password"`
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	var l Login
	err := ctx.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithResponse(ctx, e.UnknowError, "参数错误")
	}
}

// 添加用户
func AddAdminUser(ctx *gin.Context) {}

// 修改密码
func ChangePassword(ctx *gin.Context) {}

// 检测用户名是否可用
func CheckUserNameExist(ctx *gin.Context) {}

// 获取用户列表
func GetUserList(ctx *gin.Context) {}

// 获取用户详情
func GetUserInfo(ctx *gin.Context) {}
