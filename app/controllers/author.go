package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/chaiyd/mm-wiki/app/services"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
)

type AuthorController struct {
	BaseController
}

// login index
func (this *AuthorController) Index() {

	// is open auth login
	ssoOpen := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyAuthLogin, "0")
	this.Data["sso_open"] = ssoOpen
	this.viewLayout("author/login", "author")
}

// login
func (this *AuthorController) Login() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！")
	}
	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))

	if username == "" {
		this.jsonError("<LABEL_262>！")
	}
	if strings.Contains(username, "_") {
		this.jsonError("<LABEL_373>！")
	}
	if password == "" {
		this.jsonError("<LABEL_737>！")
	}

	user, err := models.UserModel.GetUserByUsername(username)
	if err != nil {
		this.jsonError("<LABEL_1140>")
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_263>")
	}
	if user["is_forbidden"] == fmt.Sprintf("%d", models.User_Forbidden_True) {
		this.jsonError("<LABEL_556>")
	}

	password = utils.Encrypt.Md5Encode(password)

	if user["password"] != password {
		this.jsonError("<LABEL_263>")
	}

	// update last_ip and last_login_time
	updateValue := map[string]interface{}{
		"last_time": time.Now().Unix(),
		"last_ip":   this.GetClientIp(),
	}
	_, err = models.UserModel.Update(user["user_id"], updateValue)
	if err != nil {
		this.jsonError("<LABEL_1140>")
	}

	// save session
	this.SetSession("author", user)
	// save cookie
	identify := utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.GetClientIp() + password)
	passportValue := utils.Encrypt.Base64Encode(username + "@" + identify)
	passport := beego.AppConfig.String("author::passport")
	cookieExpired, _ := beego.AppConfig.Int64("author::cookie_expired")
	this.Ctx.SetCookie(passport, passportValue, cookieExpired)

	this.Ctx.Request.PostForm.Del("password")

	this.InfoLog("<LABEL_1141>")
	this.jsonSuccess("<LABEL_1141>！", nil, "/main/index")
}

// auth login
func (this *AuthorController) AuthLogin() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！")
	}

	// is open auth login
	authLoginConf := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyAuthLogin, "0")
	if authLoginConf != "1" {
		this.jsonError("<LABEL_124>！")
	}
	// get auth login config
	authLogin, err := models.AuthModel.GetUsedAuth()
	if err != nil || len(authLogin) == 0 {
		this.jsonError("<LABEL_125>！")
	}

	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	if username == "" {
		this.jsonError("<LABEL_126>！")
	}
	if password == "" {
		this.jsonError("<LABEL_194>！")
	}
	authLoginRes, err := services.AuthLogin.AuthLogin(username, password)
	if err != nil {
		logs.Error("<LABEL_738>：", err.Error())
		this.jsonError("<LABEL_738>！")
		return
	}
	if authLoginRes == nil {
		this.jsonError("<LABEL_738>！")
		return
	}
	//realUsername := authLogin["username_prefix"] + "_" + username
	realUsername := username
	passwordEncode := models.UserModel.EncodePassword(password)
	userValue := map[string]interface{}{
		"username":   realUsername,
		"given_name": authLoginRes.GivenName,
		"password":   passwordEncode,
		"email":      authLoginRes.Email,
		"mobile":     authLoginRes.Mobile,
		"phone":      authLoginRes.Phone,
		"department": authLoginRes.Department,
		"position":   authLoginRes.Position,
		"location":   authLoginRes.Location,
		"im":         authLoginRes.Im,
		"last_time":  time.Now().Unix(),
		"last_ip":    this.GetClientIp(),
	}
	ok, err := models.UserModel.HasUsername(realUsername)
	if err != nil {
		this.jsonError("<LABEL_969>")
	}
	if ok {
		// update user info
		_, err = models.UserModel.UpdateUserByUsername(userValue)
	} else {
		// insert user info
		userValue["role_id"] = models.Role_Default_Id
		_, err = models.UserModel.Insert(userValue)
	}
	if err != nil {
		this.jsonError("<LABEL_1142>！" + err.Error())
	}

	// get user by username
	user, err := models.UserModel.GetUserByUsername(realUsername)
	if err != nil {
		this.jsonError("<LABEL_1142>：" + err.Error())
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_969>")
	}

	// save session
	this.SetSession("author", user)
	// save cookie
	identify := utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.GetClientIp() + passwordEncode)
	passportValue := utils.Encrypt.Base64Encode(user["username"] + "@" + identify)
	passport := beego.AppConfig.String("author::passport")
	cookieExpired, _ := beego.AppConfig.Int64("author::cookie_expired")
	this.Ctx.SetCookie(passport, passportValue, cookieExpired)

	this.Ctx.Request.PostForm.Del("password")

	this.InfoLog("<LABEL_1141>")
	this.jsonSuccess("<LABEL_1141>！", nil, "/main/index")
}

//logout
func (this *AuthorController) Logout() {
	this.InfoLog("<LABEL_1143>")
	passport := beego.AppConfig.String("author::passport")
	this.Ctx.SetCookie(passport, "")
	this.SetSession("author", nil)
	this.DelSession("author")

	this.Redirect("/", 302)
}
