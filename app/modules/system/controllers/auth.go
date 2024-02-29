package controllers

import (
	"github.com/chaiyd/mm-wiki/app/services"
	"net/url"
	"strings"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var auths []map[string]string
	if keyword != "" {
		count, err = models.AuthModel.CountAuthsByKeyword(keyword)
		auths, err = models.AuthModel.GetAuthsByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.AuthModel.CountAuths()
		auths, err = models.AuthModel.GetAuthsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_219>: " + err.Error())
		this.ViewError("<LABEL_219>", "/system/main/index")
	}

	this.Data["auths"] = auths
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("auth/list", "auth")
}

func (this *AuthController) Add() {
	this.viewLayout("auth/form", "auth")
}

func (this *AuthController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/auth/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	authUrl := strings.TrimSpace(this.GetString("url", ""))
	//usernamePrefix := strings.TrimSpace(this.GetString("username_prefix", ""))
	extData := strings.TrimSpace(this.GetString("ext_data", ""))

	//v := validation.Validation{}
	if name == "" {
		this.jsonError("<LABEL_220>！")
	}
	//if usernamePrefix == "" {
	//	this.jsonError("<LABEL_296>！")
	//}
	//if !v.AlphaNumeric(usernamePrefix, "username_prefix").Ok {
	//	this.jsonError("<LABEL_221>！")
	//}
	if authUrl == "" {
		this.jsonError("<LABEL_1635>URL<LABEL_1117>！")
	}
	u, err := url.Parse(authUrl)
	if err != nil || u == nil {
		this.jsonError("<LABEL_1635> URL <LABEL_1199>！")
		return
	}
	if !services.AuthLogin.UrlIsSupport(u.Scheme) {
		this.jsonError("<LABEL_1635> URL <LABEL_975>！")
	}

	ok, err := models.AuthModel.HasAuthName(name)
	if err != nil {
		this.ErrorLog("<LABEL_439>：" + err.Error())
		this.jsonError("<LABEL_439>！")
	}
	if ok {
		this.jsonError("<LABEL_222>！")
	}

	//ok, err = models.AuthModel.HasAuthUsernamePrefix(usernamePrefix)
	//if err != nil {
	//	this.ErrorLog("<LABEL_439>：" + err.Error())
	//	this.jsonError("<LABEL_439>！")
	//}
	//if ok {
	//	this.jsonError("<LABEL_297>！")
	//}

	authId, err := models.AuthModel.Insert(map[string]interface{}{
		"name": name,
		"url":  authUrl,
		//"username_prefix": usernamePrefix,
		"ext_data": extData,
	})

	if err != nil {
		this.ErrorLog("<LABEL_439>：" + err.Error())
		this.jsonError("<LABEL_439>")
	}
	this.InfoLog("<LABEL_830> " + utils.Convert.IntToString(authId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_440>", nil, "/system/auth/list")
}

func (this *AuthController) Edit() {

	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.ViewError("<LABEL_591>", "/system/auth/list")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ViewError("<LABEL_591>", "/system/auth/list")
	}

	this.Data["auth"] = auth
	this.viewLayout("auth/form", "auth")
}

func (this *AuthController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	authUrl := strings.TrimSpace(this.GetString("url", ""))
	//usernamePrefix := strings.TrimSpace(this.GetString("username_prefix", ""))
	extData := strings.TrimSpace(this.GetString("ext_data", ""))

	//v := validation.Validation{}
	if authId == "" {
		this.jsonError("<LABEL_591>！")
	}
	if name == "" {
		this.jsonError("<LABEL_220>！")
	}
	//if usernamePrefix == "" {
	//	this.jsonError("<LABEL_296>！")
	//}
	//if !v.AlphaNumeric(usernamePrefix, "username_prefix").Ok {
	//	this.jsonError("<LABEL_221>！")
	//}
	if authUrl == "" {
		this.jsonError("<LABEL_1635>URL<LABEL_1117>！")
	}
	u, err := url.Parse(authUrl)
	if err != nil || u == nil {
		this.jsonError("<LABEL_1635> URL <LABEL_1199>！")
		return
	}
	if !services.AuthLogin.UrlIsSupport(u.Scheme) {
		this.jsonError("<LABEL_1635> URL <LABEL_975>！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("<LABEL_831> " + authId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_441>！")
	}
	if len(auth) == 0 {
		this.jsonError("<LABEL_591>！")
	}

	//ok, _ := models.AuthModel.HasSameName(authId, name)
	//if ok {
	//	this.jsonError("<LABEL_222>！")
	//}
	//ok, _ = models.AuthModel.HasSameUsernamePrefix(authId, usernamePrefix)
	//if ok {
	//	this.jsonError("<LABEL_297>！")
	//}

	_, err = models.AuthModel.Update(authId, map[string]interface{}{
		"name": name,
		"url":  authUrl,
		//"username_prefix": usernamePrefix,
		"ext_data": extData,
	})

	if err != nil {
		this.ErrorLog("<LABEL_831> " + authId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_441>")
	}
	this.InfoLog("<LABEL_831> " + authId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_442>", nil, "/system/auth/list")
}

func (this *AuthController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.jsonError("<LABEL_443>！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("<LABEL_832> " + authId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_444>")
	}
	if len(auth) == 0 {
		this.jsonError("<LABEL_591>")
	}

	err = models.AuthModel.Delete(authId)
	if err != nil {
		this.ErrorLog("<LABEL_832> " + authId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_444>")
	}

	this.InfoLog("<LABEL_832> " + authId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_445>", nil, "/system/auth/list")
}

func (this *AuthController) Used() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.jsonError("<LABEL_443>！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("<LABEL_1200> " + authId + " <LABEL_1184>: " + err.Error())
		this.jsonError("<LABEL_446>")
	}
	if len(auth) == 0 {
		this.jsonError("<LABEL_591>")
	}
	_, err = models.AuthModel.SetAuthUsed(authId)
	if err != nil {
		this.ErrorLog("<LABEL_1200> " + authId + " <LABEL_1184>: " + err.Error())
		this.jsonError("<LABEL_446>")
	}

	this.InfoLog("<LABEL_833> " + authId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_447>", nil, "/system/auth/list")
}

func (this *AuthController) Doc() {
	this.viewLayout("auth/doc", "auth")
}
