package service

import (
	"encoding/json"
	"invest_dairy/bizerrors"
	"invest_dairy/common"
	"invest_dairy/model"
	"invest_dairy/util"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type UserRegisterBo struct {
	Phone   string `json:"phone"`
	SmsCode string `json:"sms_code"`
	Sex     int    `json:"sex"`
	Age     int    `json:"age"`
}

type UserInfoVo struct {
	Id               int    `json:"id"`
	Token            string `json:"token"`
	Phone            string `json:"phone"`
	Photo            string `json:"photo"`
	Name             string `json:"name"`
	Mail             string `json:"mail"`
	Gender           string `json:"gender"`
	CurrentState     string `json:"current_state"`
	CurrentCity      string `json:"current_city"`
	Birth            string `json:"birth"`
	Education        string `json:"education"`
	SchoolName       string `json:"school_name"`
	Major            string `json:"major"`
	StartYear        string `json:"start_year"`
	Company          string `json:"company"`
	Industry         string `json:"industry"`
	StartTime        string `json:"start_time"`
	FunctionArea     string `json:"function_area"`
	Department       string `json:"department"`
	JobContent       string `json:"job_content"`
	Achievement      string `json:"achievement"`
	JobType          string `json:"job_type"`
	JobFunctionArea  string `json:"job_function_area"`
	ExpectedIndustry string `json:"expected_industry"`
	PreferredCity    string `json:"preferred_city"`
	ExpectedSalary   string `json:"expected_salary"`
	IsNew            bool   `json:"is_new"`
	IsVip            bool   `json:"is_vip"`
}

func DoToVo(d *model.User) *UserInfoVo {
	u := new(UserInfoVo)
	u.Id = d.Id
	u.Phone = d.Phone
	u.Photo = d.Photo
	u.Name = d.Name
	u.Mail = d.Mail
	u.Gender = d.Gender
	u.CurrentState = d.CurrentState
	u.CurrentCity = d.CurrentCity
	u.Birth = d.Birth
	u.Education = d.Education
	u.SchoolName = d.SchoolName
	u.Major = d.Major
	u.StartYear = d.StartYear
	u.Company = d.Company
	u.Industry = d.Industry
	u.StartTime = d.StartTime
	u.FunctionArea = d.FunctionArea
	u.Department = d.Department
	u.JobContent = d.JobContent
	u.Achievement = d.Achievement
	u.JobType = d.JobType
	u.JobFunctionArea = d.JobFunctionArea
	u.ExpectedIndustry = d.ExpectedIndustry
	u.PreferredCity = d.PreferredCity
	u.ExpectedSalary = d.ExpectedSalary
	return u
}

type UserLoginBo struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func Register(bo UserLoginBo) (*UserInfoVo, error) {
	userInfo := new(model.User)
	userInfo.Phone = bo.Phone
	userInfo.CreateTime = time.Now().Unix()
	userInfo.ModifyTime = time.Now().Unix()
	err := userInfo.Insert()
	if err != nil {
		return nil, errors.New("Registration failed. Please try again later")
	}
	userInfoVo := new(UserInfoVo)
	userInfoVo.Id = userInfo.Id
	userInfoVo.Token = util.Md5Str(bo.Phone + bo.Code)
	userInfoVo.Phone = bo.Phone
	userInfoJson, _ := json.Marshal(userInfoVo)
	common.RedisCache.Set(common.USER_INFO+userInfoVo.Token, userInfoJson, time.Hour*24*7)
	return userInfoVo, nil
}

func Login(info UserLoginBo) *common.ResponseData {
	user, err := model.ExistUserPhone(info.Phone)
	if err != nil {
		common.Mlog.Error(err)
		return common.CommonError()
	}
	userInfo := new(UserInfoVo)
	if user.Id <= 0 {
		userInfo, err = Register(info)
		if err != nil {
			return common.SetError(err.Error())
		}
		userInfo.IsNew = true
	} else {
		userInfo.IsNew = len(user.JobContent) <= 0
	}
	userInfo.Token = util.Md5Str(info.Phone + strconv.Itoa(int(time.Now().Unix())))
	userInfoJson, _ := json.Marshal(userInfo)
	common.RedisCache.Set(common.USER_INFO+userInfo.Token, userInfoJson, time.Hour*24*7)
	return common.SetData(userInfo)
}

func LoginOut(user *UserInfoVo) *common.ResponseData {
	common.RedisCache.Del(common.USER_INFO + user.Token)
	return common.CommonSuccess()
}

func VerifyUserToken(ctx echo.Context) (*UserInfoVo, error) {
	token := ctx.Request().Header.Get("ACCESS_TOKEN")
	if token == "" {
		token = ctx.Request().Header.Get("access_token")
	}
	var userInfoJson = common.RedisCache.Get(common.USER_INFO + token).Val()
	if len(userInfoJson) <= 0 {
		return nil, bizerrors.VerifyTokenError
	}
	userInfo := new(UserInfoVo)
	json.Unmarshal([]byte(userInfoJson), &userInfo)
	return userInfo, nil
}
