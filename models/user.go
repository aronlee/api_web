package models

import (
	"api_web/util"
	"fmt"
	"math/rand"
	"time"

	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(UserLogin))
}

// UserLogin 用户登录信息
type UserLogin struct {
	UID       int       `json:"u_i_d" orm:"pk"`
	Username  string    `json:"username"`
	Passcode  string    `json:"passcode"` //加密随机串
	Password  string    `json:"password"`
	LoginTime time.Time `json:"login_time"`
}

// TableName 返回表名
func (*UserLogin) TableName() string {
	return "user_login"
}

// User 用户信息
type User struct {
	UID        int       `json:"u_i_d" orm:"pk;auto"`
	Username   string    `json:"username" orm:"size(100)"`
	Email      string    `json:"email" orm:"size(100)"`
	Status     int       `json:"status"`
	IsThird    int       `json:"is_third"`
	IsRoot     int       `json:"is_root"`
	CreateTime time.Time `json:"create_time"`
}

// TableName 返回表名
func (*User) TableName() string {
	return "user_info"
}

// UserExists 获取用户是否存在
func UserExists(field string, val string) bool {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	exist := qs.Filter(field, val).Exist()
	return exist
}

// CreateUser 创建用户
func CreateUser(u User, pwd string) (*UserLogin, error) {
	var ug = &UserLogin{}
	o := orm.NewOrm()
	o.Begin()
	user := User{
		Username:   u.Username,
		Email:      "112",
		Status:     0,
		IsThird:    0,
		IsRoot:     0,
		CreateTime: time.Now(),
	}
	id, err := o.Insert(&user)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	ug.UID = int(id)
	ug.Username = u.Username
	ug.LoginTime = time.Now()
	if len(pwd) > 0 {
		ug.Password = pwd
	} else {
		err = o.Rollback()
		beego.Error(err)
		return nil, errors.New("密码不能为空！")
	}
	err = ug.GetMD5Password()
	if err != nil {
		err = o.Rollback()
		beego.Error(err)
		return nil, err
	}
	id, err = o.Insert(ug)
	if err != nil {
		err = o.Rollback()
		beego.Error(err)
		return nil, err
	}
	err = o.Commit()
	if err != nil {
		beego.Error(err)
	}
	return ug, err
}

var (
	// ErrorName 用户名不存在
	ErrorName = errors.New("用户名不存在！")
	// ErrorPassword 密码错误
	ErrorPassword = errors.New("密码错误！")
)

// Login 用户登录
func Login(username, password string) (*UserLogin, error) {
	var userLogin = &UserLogin{}
	o := orm.NewOrm()
	ug := new(UserLogin)
	qs := o.QueryTable(ug)
	err := qs.Filter("username", username).One(userLogin)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return nil, errors.New("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return nil, ErrorName
	}
	pwd := util.ToHexStr(password + userLogin.Passcode)
	if pwd != userLogin.Password {
		return nil, ErrorPassword
	}
	go func() {
		UpdateLoginTime(userLogin)
	}()
	return userLogin, nil
}

// UpdateLoginTime 更新用户登录时间
func UpdateLoginTime(userLogin *UserLogin) {
	o := orm.NewOrm()
	ug := new(UserLogin)
	_, err := o.QueryTable(ug).Filter("u_i_d", userLogin.UID).Update(orm.Params{
		"login_time": time.Now(),
	})
	if err != nil {
		errStr := fmt.Sprintf("更新user：%s,登陆时间失败", userLogin.Username)
		beego.Error(errStr)
	}
}

// GetMD5Password chuan
func (ug *UserLogin) GetMD5Password() error {
	if ug.Password == "" {
		return errors.New("password is empty")
	}
	ug.Passcode = fmt.Sprintf("%x", rand.Int31())
	// 密码经过md5(Password+passcode)加密保存
	ug.Password = util.ToHexStr(ug.Password + ug.Passcode)
	return nil
}

// GetUserByUID 根据uid查找user
func GetUserByUID(id int) (*User, error) {
	var user = User{
		UID: id,
	}
	o := orm.NewOrm()
	err := o.Read(&user, "UID")
	if err != nil {
		return nil, err
	}
	return &user, nil
}
