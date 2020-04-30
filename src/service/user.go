package service

import (
	"errors"
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gasolzhang/jdbackend/src/dao"
	"github.com/gasolzhang/jdbackend/src/info"
	"github.com/gasolzhang/jdbackend/src/jwt"
	"log"
	"time"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService() *UserService {
	service := new(UserService)
	service.dao = dao.NewUserDao()
	return service
}

func (this *UserService) Login(phone, password string) (err error, code int) {
	err = this.dao.Login(phone, password)
	if err != nil {
		log.Printf("登录失败：错误信息：%v", err)
		return errors.New("账号或密码错误"), consts.ErrorCodeServer
	}
	return nil, 0
}

func (this *UserService) Regist(phone, password, verifyCode string) (err error, code int) {
	//TODO 先检查验证码是否正确
	err = this.dao.Regist(phone, password)
	if err != nil {
		log.Printf("注册失败：%v", err.Error())
		return errors.New("账号已存在，无法再注册"), consts.ErrorCodeServer
	}
	return nil, 0
}

func (this *UserService) GetUserInfo(phone string) (user info.User, err error) {
	user, err = this.dao.GetUserInfo(phone)
	if err != nil {
		log.Printf("获取用户信息失败，%v", err)
		return user, errors.New("获取用户信息失败，请确认手机号码是否正确")
	}
	return user, nil
}

func (this *UserService) GenToken(phone string) (token string) {
	claim := map[string]string{
		"iss ":  "gasolzhang",              //签发人
		"iat":   string(time.Now().Unix()), //签发时间
		"phone": phone,
	}
	return jwt.Generate(claim, 86400) //token时效24h
}

func (this *UserService) ResetPassword(phone, password, verifyCode string) (err error, code int) {
	//TODO 先检查验证码是否正确
	err = this.dao.ResetPassword(phone, password)
	if err != nil {
		log.Printf("重置密码失败:%v", err)
		return errors.New("重置密码失败"), consts.ErrorCodeServer
	}
	return nil, 0
}

func (this *UserService) ModifyPassword(phone, password string) (err error, code int) {
	err = this.dao.ResetPassword(phone, password)
	if err != nil {
		log.Printf("修改密码失败:%v", err)
		return errors.New("修改密码失败"), consts.ErrorCodeServer
	}
	return nil, 0
}
