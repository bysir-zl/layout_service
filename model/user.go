package model

import (
	"fmt"
	"time"
	"github.com/bysir-zl/bygo/log"
	"strconv"
	"github.com/bysir-zl/layout/config"
	"github.com/bysir-zl/bygo/util/auth"
	"github.com/bysir-zl/bygo/wx_open/mp"
)

// 用户
type User struct {
	Id         int64  `json:"id" xorm:"not null pk autoincr INT(10)"`
	IcCardId   string `json:"ic_card_id" xorm:"varchar(50)"`
	Nickname   string `json:"nickname" xorm:"varchar(50)"` // 我们自定义的昵称
	Sex        int    `json:"sex" xorm:"int(11)"`          // 0: 1:男 2:女
	Headimgurl string `json:"headimgurl" xorm:"varchar(200)"`
	Wx         string `json:"wx" xorm:"varchar(50)"`          // 微信号
	WxNickname string `json:"wx_nickname" xorm:"varchar(50)"` // 微信昵称
	WxOpenId   string `json:"wx_open_id" xorm:"varchar(50)"`
	WxUnionid  string `json:"wx_unionid" xorm:"varchar(50)"`
	IsCanLogin bool   `json:"is_can_login" xorm:"default b'1' bit(1)"` // 是否能登陆
	IsInBlack  bool   `json:"is_in_black" xorm:"default b'0' bit(1)"`  // 是否在黑名单
	Point      int    `json:"point" xorm:"default 0 int(11)"`          // 积分
	FixedPhone string `json:"fixed_phone" xorm:"varchar(50)"`          // 座机
	Phone      string `json:"phone" xorm:"varchar(50)"`                // 手机号
	// Address     string `json:"address" xorm:"varchar(200)"`             // 详细地址
	Age                  int    `json:"age" xorm:"-"`                // 年龄, 由生日算
	Garden               string `json:"garden" xorm:"varchar(100)"`  // 小区名称
	Building             string `json:"building" xorm:"varchar(20)"` // 栋
	Unit                 string `json:"unit" xorm:"varchar(20)"`     // 单元
	HomeNum              string `json:"home_num" xorm:"varchar(20)"` // 房号
	Birthday             string `json:"birthday" xorm:"varchar(20)"` // 生日
	Level                int    `json:"level" xorm:"tinyint(4)"`     // 等级 1: 普通 2: vip
	CreatedAt            int64  `json:"created_at" xorm:"int(11) created"`
	LastLoginAt          int64  `json:"last_login_at" xorm:"int(11)"`
	DeleteAt             int64  `json:"-" xorm:"int(11) deleted"`
	LastPushAt           int64  `json:"last_push_at" xorm:"int(11)"` // 最后一次投放时间，降级逻辑用
	AccessToken          string `json:"-" xorm:"varchar(128)"`       // 微信accesstoken
	RefreshToken         string `json:"-" xorm:"varchar(128)"`       // 微信refreshtoken
	AccessTokenExpiredAt int64  `json:"-" xorm:"int(11)"`            // 微信accesstoken过期时间

	Token string `json:"token,omitempty" xorm:"-"`
}

const (
	UserLevelNomal = 1
	UserLevelVip   = 2
)

//  信息是否已经完整
func (u User) IsCompete() bool {
	return u.Sex != 0 && u.Phone != "" && u.Birthday != "" && u.Garden != "" && u.Building != "" && u.Unit != "" && u.HomeNum != ""
}

//  算年龄
func (u *User) CalAge() {
	if u.Birthday == "" {
		return
	}
	t, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return
	}
	u.Age = int(time.Now().Sub(t).Hours() / 24 / 365)
}

// 获取用户
func GetUser(id int64) (*User, error) {
	user := User{}
	_, err := engine.ID(id).Get(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, ErrNotFind.Append(fmt.Sprintf("user by %d", id))
	}
	user.CalAge()
	return &user, nil
}

// 获取用户
func GetUserByIc(icCardId string) (*User, error) {
	if icCardId == "" {
		return nil, ErrBadParams.Append("icCardId is empty")
	}

	s := User{}
	has, err := engine.Where("ic_card_id=?", icCardId).Get(&s)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrNotFind.Append(fmt.Sprintf("user by iccard %s", icCardId))
	}

	s.CalAge()
	return &s, nil
}

// 获取列表
func GetUserList(key string, per int, page int) (total int64, rsp []User, err error) {
	as := []User{}
	session := engine.Desc("id")

	if key != "" {
		key = "%" + key + "%"
		session = session.Where("(phone like ? OR wx_nickname like ? OR ic_card_id like ?)", key, key, key)
	}
	total, err = session.Clone().Count(new(User))
	if err != nil {
		return
	}

	err = session.Limit(per, per*(page-1)).Find(&as)
	if err != nil {
		err = err
		return
	}

	rsp = as
	return
}

// 用户登陆
// 拿到WxOpenId之后调用这个换取token
func LoginUserByWx(s *User) (*User, error) {
	//s.Id = 0
	if s.WxOpenId == "" {
		return nil, ErrBadParams.Append("WxOpenId is empty")
	}

	has, err := engine.Where("wx_open_id=?", s.WxOpenId).Exist(new(User))
	if err != nil {
		return nil, err
	}
	if has {
		u, err := GetUserByWxId(s.WxOpenId)
		if err != nil {
			return nil, err
		}
		s.LastLoginAt = time.Now().Unix()
		s.Id = u.Id
		_, err = engine.Where("wx_open_id=?", s.WxOpenId).Update(s)
		if err != nil {
			return nil, err
		}
	} else {
		s.Nickname = s.WxNickname
		s.LastLoginAt = time.Now().Unix()
		s.Level = UserLevelNomal
		_, err = engine.Insert(s)
		if err != nil {
			return nil, err
		}
	}

	u, err := GetUser(s.Id)
	if err != nil {
		return nil, err
	}

	token := auth.JWTEncode(config.Key, "", time.Now().Unix()+15*24*3600, "user", strconv.FormatInt(u.Id, 10), "")
	u.Token = token
	return u, nil
}

// 删除用户
func DelUser(id int64) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Delete(new(User))
	if err != nil {
		return err
	}

	return nil
}

// 更新用户信息
func UpdateUserBySelf(s *User) (error) {
	if s.Id == 0 {
		return ErrBadParams.Append("id is empty")
	}

	_, err := engine.ID(s.Id).
		Cols("nickname", "sex", "birthday", "garden", "building", "unit", "home_num").
		Update(s)
	if err != nil {
		return err
	}

	return nil
}

// 更新用户是否在黑名单
func UpdateUserIsInBlack(id int64, isInBlack bool) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Cols("is_in_black").Update(&User{
		IsInBlack: isInBlack,
	})
	if err != nil {
		return err
	}

	return nil
}

// 更新用户手机
func UpdateUserPhone(id int64, phone string) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Cols("phone").Update(&User{
		Phone: phone,
	})
	if err != nil {
		return err
	}

	return nil
}

// 更新用户是否能登陆
func UpdateUserCanLogin(id int64, isCanLogin bool) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Cols("is_can_login").Update(&User{
		IsCanLogin: isCanLogin,
	})
	if err != nil {
		return err
	}

	return nil
}

// 添加
func CreateUserForIcCard(s *User) (*User, error) {
	if s.IcCardId == "" {
		return nil, ErrBadParams.Append("Ic卡不能为空")
	}
	s.Wx = ""
	s.Id = 0
	s.IsInBlack = false
	s.IsCanLogin = true
	s.Level = UserLevelNomal

	// ic卡不能重复
	_, err := GetUserByIc(s.IcCardId)
	if err != nil {
		if ErrNotFind.Is(err) {
		} else {
			return nil, err
		}
	} else {
		return nil, ErrDefault.Append("Ic卡已被其他用户使用")
	}

	_, err = engine.Insert(s)
	if err != nil {
		return nil, err
	}

	return GetUser(s.Id)
}

// 获取用户
func GetUserByWxId(wid string) (*User, error) {
	user := User{}
	_, err := engine.Where("wx_open_id=?", wid).Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

var RefreshError = ErrDefault.Append("refresh error")
// 刷新用户的accessToken
func RefreshUserAccessToken(u *User) (err error) {
	if u.RefreshToken == "" {
		return RefreshError
	}
	needRefresh := false
	if u.AccessToken == "" {
		needRefresh = true
	} else if time.Now().Unix() >= u.AccessTokenExpiredAt-10 {
		needRefresh = true
	}

	if needRefresh {
		r, err := mp.RefreshUserAccessToken(config.WxAppId, u.RefreshToken)
		if err != nil {
			log.Error("RefreshUserAccessToken err: %s", err)
			return RefreshError
		}

		u.RefreshToken = r.RefreshToken
		u.AccessTokenExpiredAt = r.ExpiresIn + time.Now().Unix()
		u.AccessToken = r.AccessToken

		_, err = engine.Where("id=?", u.Id).Cols("access_token", "access_token_expired_at", "refresh_token").Update(u)
		if err != nil {
			return err
		}
	}
	return
}
