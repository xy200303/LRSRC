package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/models"
	"xiaoyun/backend/service"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/service/user_service"
	"xiaoyun/backend/types"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/utils"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/utils/template"
	"xiaoyun/backend/validate"
)

// Login 登录接口
func Login(c *gin.Context) {
	var data types.LoginReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(data, err), nil)
		return
	}
	//普通成员登录
	if data.LoginType == "user" || data.LoginType == "" {
		user, token, err := service.Login(&data)
		if err != nil {
			resp.Err(c, err)
			return
		}
		resp.Ok(c, gin.H{
			"token":    token,
			"is_admin": user.IsAdmin,
			"username": user.Username,
			"role_id":  user.RoleID,
		})
	}
	//管理员登录
	if data.LoginType == "admin" {
		user, token, err := service.Login(&data)
		if err != nil {
			resp.Err(c, err)
			return
		}
		if user.RoleID != 100 {
			resp.Resp(c, 401, "非管理员禁止登录管理系统", nil)
			return
		}
		resp.Ok(c, gin.H{
			"token":    token,
			"username": user.Username,
			"role_id":  user.RoleID,
		})
	}
	//厂商登录
	if data.LoginType == "muna" {

	}

}

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	var duration = 10 * 60
	var remainTime = 30
	capType := c.Query("type")
	data := c.Query("data")
	serviceType := c.Query("service_type")
	if serviceType != "register" && serviceType != "forget" {
		resp.Resp(c, 400, "Invalid serviceType, must be 'register' or 'forget'", nil)
		return
	}
	if capType == "email" {
		//邮箱验证码
		if !utils.IsEmailValid(data) {
			resp.Resp(c, 400, "邮箱格式不正确", nil)
			return
		}
		code := utils.GenRandomNumber(6)
		var err error
		var ttl int64
		var html string
		if serviceType == "register" {
			if user_service.IsUserExist(data) {
				resp.Err(c, fmt.Errorf("邮箱已经存在"))
				return
			}
			ttl, err = models.RedisGetRegisterEmailCodeTTL(data)
			if duration-int(ttl) < remainTime {
				err = fmt.Errorf("请%ds之后请求验证码", remainTime-(duration-int(ttl)))
			}
			if err != nil {
				resp.Err(c, err)
				return
			}
			err = models.RedisSetRegisterEmailCode(data, code, duration)
			//获取HTML模板
			if err != nil {
				resp.Err(c, err)
				return
			}
			html, err = template.RenderRegisterHtml(code, duration, system_service.SysConfigMap.SysSmtpSender)
			if err != nil {
				resp.Err(c, err)
				return
			}
		}
		if serviceType == "forget" {
			if !user_service.IsUserExist(data) {
				resp.Err(c, fmt.Errorf("不存在的邮箱"))
				return
			}
			ttl, err = models.RedisGetForgetEmailCodeTTL(data)
			if duration-int(ttl) < remainTime {
				err = fmt.Errorf("请%ds之后请求验证码", remainTime-(duration-int(ttl)))
			}
			if err != nil {
				resp.Err(c, err)
				return
			}
			err = models.RedisSetForgetEmailCode(data, code, duration)
			if err != nil {
				resp.Err(c, err)
				return
			}
			html, err = template.RenderForgetHtml(code, duration, system_service.SysConfigMap.SysSmtpSender)
			if err != nil {
				resp.Err(c, err)
				return
			}
		}
		err = utils.SendEmail(
			data,
			system_service.SysConfigMap.SysSmtpCaptchaTitle,
			html,
			system_service.SysConfigMap.SysSmtpHost,
			system_service.SysConfigMap.SysSmtpUsername,
			system_service.SysConfigMap.SysSmtpPassword,
			system_service.SysConfigMap.SysSmtpPort,
			system_service.SysConfigMap.SysSmtpSender,
		)
		if err != nil {
			log.Println(err)
			resp.Err(c, fmt.Errorf("验证码发送错误"))
			return
		}
		resp.Ok(c, nil)
		return
	}
	resp.Resp(c, 400, "不支持的验证码类型", nil)
}

// Register 注册
func Register(c *gin.Context) {
	var data types.RegisterReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	//验证验证码
	code, err := models.RedisGetRegisterEmailCode(data.Email)
	if err != nil || code != data.Code {
		resp.Err(c, fmt.Errorf("验证码错误"))
		return
	}
	//调用Service层进行用户注册
	user, err := service.Register(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	//删除验证码
	_ = models.RedisDelRegisterEmailCode(data.Email)
	resp.Ok(c, gin.H{
		"username": user.Username,
		"role_id":  user.RoleID,
	})
}

// CheckToken 测试鉴权逻辑,返回解析数据
func CheckToken(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	resp.Ok(c, auth)
}

// ForgetPassword 忘记密码
func ForgetPassword(c *gin.Context) {
	var data types.ForgetPwdReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	//查询是否存在用户
	if !service.CheckUserExist(data.Email) {
		resp.Err(c, fmt.Errorf("不存在的用户邮箱，发送验证码失败"))
		return
	}
	//验证验证码
	code, err := models.RedisGetForgetEmailCode(data.Email)
	if err != nil || code != data.Code {
		resp.Err(c, fmt.Errorf("验证码错误"))
		return
	}
	err = service.ForgetPassword(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	//删除验证码
	_ = models.RedisDelForgetEmailCode(data.Email)
	resp.Ok(c, nil)
}

// GetMyProfile 查看个人信息
func GetMyProfile(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
	}
	//过滤信息
	resp.Ok(c, gin.H{
		"username":    auth.User.Username,
		"role_id":     auth.User.RoleID,
		"email":       auth.User.Email,
		"phone":       auth.User.Phone,
		"address":     auth.User.Address,
		"gender":      auth.User.Gender,
		"birthdate":   auth.User.Birthdate,
		"nickname":    auth.User.Nickname,
		"create_time": auth.User.CreatedAt,
		"avatar":      auth.User.Avatar,
		"integral":    auth.User.Integral,
		"level":       auth.User.Level,
		"is_admin":    auth.IsAdmin,
	})
}

// Logout 退出登录
func Logout(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
	}
	err = models.RedisDelUserToken(auth.User.Username)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var data types.ChangePasswordReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err = service.ChangePassword(auth, &data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateProfile 更新用户信息
func UpdateProfile(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var data user_types.User
	//确保json请求格式
	if err = c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	//防止任意用户越权
	data.Username = auth.User.Username
	err = service.UpdateProfile(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}
