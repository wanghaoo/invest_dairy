package model

import (
	"invest_dairy/common"

	"gorm.io/gorm"
)

type User struct {
	Id               int
	Phone            string
	Photo            string
	Name             string
	Mail             string
	Gender           string
	CurrentState     string
	CurrentCity      string
	Birth            string
	Education        string
	SchoolName       string
	Major            string
	StartYear        string
	Company          string
	Industry         string
	StartTime        string
	FunctionArea     string
	Department       string
	JobContent       string
	Achievement      string
	JobType          string
	JobFunctionArea  string
	ExpectedIndustry string
	PreferredCity    string
	ExpectedSalary   string
	CreateTime       int64
	ModifyTime       int64
	PassYear         string
	EndTime          string
}

func ExistUserPhone(phone string) (*User, error) {
	return FindUserByPhone(phone)
}

func FindUserByPhone(phone string) (*User, error) {
	user := new(User)
	err := common.MySQL.Where("phone = ?", phone).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func (user *User) Detail(id int) error {
	err := common.MySQL.First(&user, id).Error
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (user *User) Insert() error {
	err := common.MySQL.Create(user).Error
	if err != nil {
		common.Mlog.Errorf("insert user error: %s", err.Error())
	}
	return err
}

func (user *User) Update() error {
	err := common.MySQL.Save(user).Error
	if err != nil {
		common.Mlog.Errorf("insert user error: %s", err.Error())
	}
	return err
}
