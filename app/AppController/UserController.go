package AppController

import (
	"github.com/BeeCmf/models"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	AppBaseController
}

const SESSION_USER_KEY = "UserData"

var cpt *captcha.Captcha //验证码
func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 120
	cpt.StdHeight = 38
}

// @router /login [get]
func (c *UserController) Login() {
	c.TplName = "app/login.html"
}

// @router /login [post]
func (c *UserController) LoginPost() {
	//验证邮箱密码是否正确
	user, _ := models.UserLogin(
		c.GetString("username"),
		c.GetString("password"),
	)
	logs.Info("查询的用户数据", user)
	return
	if !cpt.VerifyReq(c.Ctx.Request) {
		c.Abort500("验证码不正确", "")
	}
	valid := validation.Validation{}
	valid.Required(c.GetString("username"), "login").Message("帐号必须填写")
	valid.Required(c.GetString("password"), "password").Message("密码必须填写")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Abort500(err.Message, "")
		}
	}
	//验证邮箱密码是否正确
	//user, err := models.UserLogin(
	//	c.GetString("username"),
	//	c.GetString("password"),
	//)
	//if err != nil {
	//	c.Abort500(err.Error(), "")
	//} else {
	//	//验证成功 进行session的记录
	//	c.SetSession(SESSION_USER_KEY, user)
	//	c.Abort200("登录成功", "/")
	//}
}

// @router /user [get]
func (c *UserController) User() {
	c.Data["Success"] = cpt.VerifyReq(c.Ctx.Request)
	c.TplName = "user.html"
}