package admin

import (
	"encoding/json"
	"invest_dairy/bizerrors"
	"invest_dairy/common"
	ma "invest_dairy/model/admin"
	"invest_dairy/util"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserInfoVo struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func Login(bo ma.LoginBo) *common.ResponseData {
	result := new(UserInfoVo)
	sysUser, err := ma.Login(bo)
	if err == gorm.ErrRecordNotFound {
		return common.SetError("用户不存在")
	}
	if err != nil {
		common.Mlog.Errorf("login error : %s", err.Error())
		return common.SetError(err.Error())
	}
	result.Id = sysUser.Id
	result.NickName = sysUser.NickName
	result.Avatar = sysUser.Avatar
	result.Token = util.Md5Str(bo.UserName + strconv.Itoa(int(time.Now().Unix())))
	userJson, _ := json.Marshal(result)
	common.RedisCache.Set(result.Token, string(userJson), 0)
	return common.SetData(result)
}

func Logout(token string) *common.ResponseData {
	common.RedisCache.Del(token)
	return common.CommonSuccess()
}

func ValidateSysUser(token string) (*UserInfoVo, error) {
	result := new(UserInfoVo)
	userJson := common.RedisCache.Get(token).Val()
	if len(userJson) <= 0 {
		return nil, bizerrors.VerifyTokenError
	}
	json.Unmarshal([]byte(userJson), &result)
	return result, nil
}

func FindSysUser(filter ma.SysUserFilter) *common.ResponseData {
	users, err := ma.FindSysUser(filter)
	if err != nil {
		common.Mlog.Errorf("find sys user error: %s", err.Error())
		return common.CommonError()
	}
	total, err := ma.FindSysUserCount(filter)
	if err != nil {
		common.Mlog.Errorf("find sys user count error: %s", err.Error())
		return common.CommonError()
	}
	return common.SetResult(users, filter.PageNo, filter.PageSize, total)
}

type SysUserBo struct {
	Id              int    `json:"id"`
	NickName        string ` json:"nick_name"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Status          int    `json:"status"`
}

func AddOrUpdateSysUser(bo SysUserBo) *common.ResponseData {
	user := new(ma.SysUser)
	if bo.Id > 0 {
		user.Id = bo.Id
		err := user.Detail()
		if err != nil {
			common.Mlog.Errorf("load sys user error: %s", err.Error())
			return common.CommonError()
		}
	}
	user.NickName = bo.NickName
	user.UserName = bo.UserName
	user.Status = bo.Status
	if len(bo.ConfirmPassword) > 0 {
		if bo.ConfirmPassword != bo.Password {
			return common.SetError("密码不一致")
		}
		user.Password = bo.Password
	}
	user.ModifyTime = time.Now().Unix()
	var err error
	if bo.Id <= 0 {
		user.CreateTime = time.Now().Unix()
		err = user.Insert()
	} else {
		err = user.Update()
	}
	if err != nil {
		common.Mlog.Errorf("insert or update sys user error: %s", err.Error())
		return common.CommonError()
	}
	return common.CommonSuccess()
}

type SysUserDetail struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Status   int    `json:"status"`
}

func FindUserDetail(id int) *common.ResponseData {
	user, err := ma.GetSysUser(id)
	if err != nil {
		common.Mlog.Errorf("load sys user error: %s", err.Error())
		return common.CommonError()
	}
	sysUser := SysUserDetail{}
	sysUser.Id = user.Id
	sysUser.Status = user.Status
	sysUser.UserName = user.UserName
	sysUser.NickName = user.NickName
	return common.SetData(sysUser)
}

type ModifySysUserBo struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func UpdatePassword(bo ModifySysUserBo, user *UserInfoVo) *common.ResponseData {
	sysUser, err := ma.GetSysUser(user.Id)
	if err != nil {
		common.Mlog.Errorf("load sys user error: %s", err.Error())
		return common.CommonError()
	}
	if bo.NewPassword != bo.ConfirmPassword {
		return common.CommonError()
	}
	if bo.Password != sysUser.Password {
		return common.SetError("密码错误")
	}
	sysUser.Password = bo.ConfirmPassword
	err = sysUser.Update()
	if err != nil {
		common.Mlog.Errorf("update sys user error: %s", err.Error())
		return common.CommonError()
	}
	return common.CommonSuccess()
}
