package admin

import (
	"fmt"
	"invest_dairy/common"
	"invest_dairy/util"
)

type SysUserLoginVo struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type LoginBo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Login(bo LoginBo) (*SysUserLoginVo, error) {
	user := new(SysUserLoginVo)
	err := common.MySQL.Select("id, user_name,nick_name, avatar").Table("b_sys_user").
		Where("user_name = ?", bo.UserName).
		Where("password = ?", bo.Password).
		Where("status = 0").First(&user).Error
	if err != nil {
		common.Mlog.Errorf("query login error: %s", err.Error())
	}
	return user, err
}

type SysUser struct {
	Id         int
	UserName   string
	NickName   string
	Password   string
	CreateTime int64
	ModifyTime int64
	Status     int
}

type SysUserFilter struct {
	NickName string `json:"nick_name"`
	util.Pagination
}

type SysUserVo struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Status   int    `json:"status"`
}

func FindSysUser(filter SysUserFilter) ([]SysUserVo, error) {
	result := make([]SysUserVo, 0)
	querySql := common.MySQL.Select("id, nick_name, user_name, status").Table("b_sys_user")
	if len(filter.NickName) > 0 {
		querySql = querySql.Where(fmt.Sprintf("nick_name like '%%%s%%'", filter.NickName))
	}
	err := filter.PageLimit(querySql).Find(&result).Error
	if err != nil {
		common.Mlog.Errorf("query sys user error: %s", err.Error())
	}
	return result, err
}

func FindSysUserCount(filter SysUserFilter) (int, error) {
	querySql := common.MySQL.Model(&SysUser{})
	if len(filter.NickName) > 0 {
		querySql = querySql.Where(fmt.Sprintf("nick_name like '%%%s%%'", filter.NickName))
	}
	var count int64
	err := querySql.Count(&count).Error
	if err != nil {
		common.Mlog.Errorf("scan count sys user error: %s", err.Error())
	}
	return int(count), nil
}

func GetSysUser(id int) (*SysUser, error) {
	sysUser := new(SysUser)
	err := common.MySQL.First(&sysUser, id).Error
	if err != nil {
		common.Mlog.Errorf("find sys user error: %s", err.Error())
	}
	return sysUser, err
}

func (user *SysUser) Update() error {
	err := common.MySQL.Save(user).Error
	if err != nil {
		common.Mlog.Errorf("update sys user error: %s", err.Error())
	}
	return err
}

func (user *SysUser) Insert() error {
	err := common.MySQL.Create(user).Error
	if err != nil {
		common.Mlog.Errorf("insert sys user error: %s", err.Error())
	}
	return err
}

func (user *SysUser) Detail() error {
	err := common.MySQL.First(&user, user.Id).Error
	if err != nil {
		common.Mlog.Errorf("load sys user error: %s", err.Error())
	}
	return err
}
