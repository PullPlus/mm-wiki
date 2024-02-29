package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Add() {

	roles := []map[string]string{}
	var err error

	if this.IsRoot() {
		roles, err = models.RoleModel.GetRoles()
	} else {
		roles, err = models.RoleModel.GetRolesNotContainRoot()
	}
	if err != nil {
		this.ErrorLog("<LABEL_409>：" + err.Error())
		this.ViewError("<LABEL_409>！")
	}
	this.Data["roles"] = roles
	this.viewLayout("user/form", "user")
}

func (this *UserController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/user/add")
	}
	username := strings.TrimSpace(this.GetString("username", ""))
	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))
	this.Ctx.Request.PostForm.Del("password")

	v := validation.Validation{}
	if username == "" {
		this.jsonError("<LABEL_578>！")
	}
	if !v.AlphaNumeric(username, "username").Ok {
		this.jsonError("<LABEL_356>！")
	}
	if givenName == "" {
		this.jsonError("<LABEL_787>！")
	}
	if password == "" {
		this.jsonError("<LABEL_737>！")
	}
	if email == "" {
		this.jsonError("<LABEL_774>！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("<LABEL_568>！")
	}
	if mobile == "" {
		this.jsonError("<LABEL_576>！")
	}
	//if !v.Mobile(mobile, "mobile").Ok {
	//	this.jsonError("<LABEL_404>！")
	//}
	if roleId == "" {
		this.jsonError("<LABEL_792>！")
	}
	//if phone != "" && !v.Phone(phone, "phone").Ok {
	//	this.jsonError("<LABEL_577>！")
	//}

	ok, err := models.UserModel.HasUsername(username)
	if err != nil {
		this.ErrorLog("<LABEL_793>：" + err.Error())
		this.jsonError("<LABEL_793>！")
	}
	if ok {
		this.jsonError("<LABEL_579>！")
	}

	if !this.IsRoot() && roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_135>！")
	}

	userId, err := models.UserModel.Insert(map[string]interface{}{
		"username":   username,
		"given_name": givenName,
		"password":   models.UserModel.EncodePassword(password),
		"email":      email,
		"mobile":     mobile,
		"phone":      phone,
		"department": department,
		"position":   position,
		"location":   location,
		"im":         im,
		"role_id":    roleId,
	})

	if err != nil {
		this.ErrorLog("<LABEL_793>：" + err.Error())
		this.jsonError("<LABEL_793>")
	}
	this.InfoLog("<LABEL_1175> " + utils.Convert.IntToString(userId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_794>", nil, "/system/user/list")
}

func (this *UserController) List() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	username := strings.TrimSpace(this.GetString("username", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)

	if username != "" {
		keywords["username"] = username
	}
	if roleId != "" {
		keywords["role_id"] = roleId
	}

	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	if len(keywords) != 0 {
		count, err = models.UserModel.CountUsersByKeywords(keywords)
		users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	} else {
		count, err = models.UserModel.CountUsers()
		users, err = models.UserModel.GetUsersByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_398>: " + err.Error())
		this.ViewError("<LABEL_398>", "/system/main/index")
	}

	var roleIds = []string{}
	if roleId != "" {
		roleIds = append(roleIds, roleId)
	} else {
		for _, user := range users {
			roleIds = append(roleIds, user["role_id"])
		}
	}
	roles, err := models.RoleModel.GetRoleByRoleIds(roleIds)
	if err != nil {
		this.ErrorLog("<LABEL_398>: " + err.Error())
		this.ViewError("<LABEL_274>", "/system/main/index")
	}
	var roleUsers = []map[string]string{}
	for _, user := range users {
		roleUser := user
		for _, role := range roles {
			if role["role_id"] == user["role_id"] {
				roleUser["role_name"] = role["name"]
				break
			}
		}
		roleUsers = append(roleUsers, roleUser)
	}

	allRoles, err := models.RoleModel.GetRoles()
	if err != nil {
		this.ErrorLog("<LABEL_398>: " + err.Error())
		this.ViewError("<LABEL_398>！", "/system/main/index")
	}
	this.Data["users"] = roleUsers
	this.Data["username"] = username
	this.Data["roleId"] = roleId
	this.Data["roles"] = allRoles
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "user")
}

func (this *UserController) Edit() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("<LABEL_962>！", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_711>：" + err.Error())
		this.ViewError("<LABEL_711>！", "/system/user/list")
	}
	if len(user) == 0 {
		this.ViewError("<LABEL_962>！", "/system/user/list")
	}
	// <LABEL_1505> root <LABEL_795> root <LABEL_1176>
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) && !this.IsRoot() {
		this.ViewError("<LABEL_796>！", "/system/user/list")
	}

	roles := []map[string]string{}
	if this.IsRoot() {
		roles, err = models.RoleModel.GetRoles()
	} else {
		roles, err = models.RoleModel.GetRolesNotContainRoot()
	}
	if err != nil {
		this.ErrorLog("<LABEL_409>：" + err.Error())
		this.ViewError("<LABEL_409>！")
	}

	this.Data["user"] = user
	this.Data["roles"] = roles
	this.viewLayout("user/edit", "user")
}

func (this *UserController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/user/list")
	}
	userId := strings.TrimSpace(this.GetString("user_id", ""))
	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	this.Ctx.Request.PostForm.Del("password")

	v := validation.Validation{}
	if givenName == "" {
		this.jsonError("<LABEL_787>！")
	}
	if email == "" {
		this.jsonError("<LABEL_774>！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("<LABEL_568>！")
	}
	if mobile == "" {
		this.jsonError("<LABEL_576>！")
	}
	//if !v.Mobile(mobile, "mobile").Ok {
	//	this.jsonError("<LABEL_404>！")
	//}
	//if roleId == "" {
	//	this.jsonError("<LABEL_792>！")
	//}
	//if phone != "" && !v.Phone(phone, "phone").Ok {
	//	this.jsonError("<LABEL_577>！")
	//}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_1177> " + userId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_797>！")
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_962>！")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		roleId = fmt.Sprintf("%d", models.Role_Root_Id)
	}
	// <LABEL_1505> root <LABEL_795> root <LABEL_1176>
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) && !this.IsRoot() {
		this.jsonError("<LABEL_796>！")
	}

	updateUser := map[string]interface{}{
		"given_name": givenName,
		"email":      email,
		"mobile":     mobile,
		"phone":      phone,
		"department": department,
		"position":   position,
		"location":   location,
		"im":         im,
	}
	// <LABEL_18>
	if password != "" && this.IsRoot() {
		updateUser["password"] = models.UserModel.EncodePassword(password)
	}
	if roleId != "" {
		updateUser["role_id"] = roleId
	}
	_, err = models.UserModel.Update(userId, updateUser)
	if err != nil {
		this.ErrorLog("<LABEL_1177> " + userId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_798>")
	}
	this.InfoLog("<LABEL_1177> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_799>", nil, "/system/user/list")
}

func (this *UserController) Forbidden() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/user/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("<LABEL_962>")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_1178> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_800>")
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_962>")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_275>")
	}
	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"is_forbidden": models.User_Forbidden_True,
	})
	if err != nil {
		this.ErrorLog("<LABEL_1178> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_800>")
	}

	this.InfoLog("<LABEL_1178> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_801>", nil, "/system/user/list")
}

func (this *UserController) Recover() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/user/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("<LABEL_962>")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_1179> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_802>")
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_962>")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_275>")
	}
	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"is_forbidden": models.User_Is_Forbidden_False,
	})
	if err != nil {
		this.ErrorLog("<LABEL_1179> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_802>")
	}

	this.InfoLog("<LABEL_1179> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_803>", nil, "/system/user/list")
}

func (this *UserController) Info() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("<LABEL_962>！", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_711>：" + err.Error())
		this.ViewError("<LABEL_711>！", "/system/user/list")
	}
	if len(user) == 0 {
		this.ViewError("<LABEL_962>！", "/system/user/list")
	}
	role, err := models.RoleModel.GetRoleByRoleId(user["role_id"])
	if err != nil {
		this.ErrorLog("<LABEL_410>：" + err.Error())
		this.ViewError("<LABEL_711>！", "/system/user/list")
	}
	this.Data["user"] = user
	this.Data["role"] = role
	this.viewLayout("user/info", "user")
}
