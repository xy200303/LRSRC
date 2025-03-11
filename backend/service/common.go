package service

import (
	"fmt"
	"log"
	"time"
	"xiaoyun/backend/models"
	"xiaoyun/backend/service/user_service"
	"xiaoyun/backend/types"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/utils"
)

// Login 普通注册用户
func Login(req *types.LoginReq) (*user_types.User, string, error) {
	jwtSecret, err := models.RedisGetJwtSecret()
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	user, err := models.FindUser(req.Username)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	if utils.IsHashEqual(user.Password, req.Password) {
		expiresIn := time.Now().Add(time.Second * time.Duration(3600)).Unix()
		token, err := utils.GenJWTToken(user.Username, user.RoleID, expiresIn, jwtSecret)
		if err != nil {
			return user, "", err
		}
		//设置token
		err = models.SetRedisValue("token_"+user.Username, token, 3600)
		if err != nil {
			return user, "", err
		}
		return user, token, err
	}
	return user, "", fmt.Errorf("用户名或密码验证错误")
}

// Register 注册逻辑
func Register(req *types.RegisterReq) (*user_types.User, error) {
	//验证手机号是否存在
	if user_service.IsUserExist(req.Username) {
		return nil, fmt.Errorf("用户名已经存在")
	}
	if user_service.IsUserExist(req.Email) {
		return nil, fmt.Errorf("邮箱已经存在")
	}
	if user_service.IsUserExist(req.Phone) {
		return nil, fmt.Errorf("手机号已经存在")
	}
	user := user_types.User{
		Username: req.Username,
		Password: utils.GenPasswordHash(req.Password),
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
	}
	err := models.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ForgetPassword 忘记密码
func ForgetPassword(req *types.ForgetPwdReq) error {
	user, err := models.FindUser(req.Email)
	if err != nil {
		return err
	}
	err = models.UpdatePassword(&user_types.User{
		Username: user.Username,
		Password: utils.GenPasswordHash(req.Password),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetUserProfile 查看个人信息
func GetUserProfile(username string) (*user_types.User, error) {
	user, err := models.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckUserExist 判断用户是否存在
func CheckUserExist(param string) bool {
	_, err := models.FindUser(param)
	if err != nil {
		return false
	}
	return true
}

// 修改密码
func ChangePassword(auth *types.Auth, req *types.ChangePasswordReq) error {
	if !utils.IsHashEqual(auth.User.Password, req.OldPassword) {
		return fmt.Errorf("旧密码错误，验证失败")
	}
	if auth.User.Password == utils.GenPasswordHash(req.NewPassword) {
		return fmt.Errorf("旧密码不能和新密码一致")
	}
	err := models.UpdatePassword(&user_types.User{
		Username: auth.User.Username,
		Password: utils.GenPasswordHash(req.NewPassword),
	})
	if err != nil {
		return err
	}
	err = models.RedisDelUserToken(auth.User.Username)
	return err
}

// 更新用户信息
func UpdateProfile(req *user_types.User) error {
	err := models.UpdateProfile(req)
	return err
}
