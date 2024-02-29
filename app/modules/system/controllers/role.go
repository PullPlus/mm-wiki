package controllers

import (
	"fmt"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"strings"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Add() {
	this.viewLayout("role/form", "role")
}

func (this *RoleController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/role/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	if name == "" {
		this.jsonError("<LABEL_422>！")
	}

	ok, err := models.RoleModel.HasRoleName(name)
	if err != nil {
		this.ErrorLog("<LABEL_813>：" + err.Error())
		this.jsonError("<LABEL_813>！")
	}
	if ok {
		this.jsonError("<LABEL_587>！")
	}

	roleId, err := models.RoleModel.Insert(map[string]interface{}{
		"name": name,
	})

	if err != nil {
		this.ErrorLog("<LABEL_813>：" + err.Error())
		this.jsonError("<LABEL_813>")
	}
	this.InfoLog("<LABEL_1187> " + utils.Convert.IntToString(roleId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_814>", nil, "/system/role/list")
}

func (this *RoleController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var roles []map[string]string
	if keyword != "" {
		count, err = models.RoleModel.CountRolesByKeyword(keyword)
		roles, err = models.RoleModel.GetRolesByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.RoleModel.CountRoles()
		roles, err = models.RoleModel.GetRolesByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_423>: " + err.Error())
		this.ViewError("<LABEL_423>", "/system/main/index")
	}

	this.Data["roles"] = roles
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("role/list", "role")
}

func (this *RoleController) Edit() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("<LABEL_986>", "/system/role/list")
	}
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.ViewError("<LABEL_290>", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_815>：" + err.Error())
		this.ViewError("<LABEL_815>", "/system/role/list")
	}
	if len(role) == 0 {
		this.ViewError("<LABEL_986>", "/system/role/list")
	}

	this.Data["role"] = role
	this.viewLayout("role/form", "role")
}

func (this *RoleController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/role/list")
	}
	roleId := this.GetString("role_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))

	if roleId == "" {
		this.jsonError("<LABEL_986>！")
	}
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_291>！")
	}
	if name == "" {
		this.jsonError("<LABEL_422>！")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1188> " + roleId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_816>！")
	}
	if len(role) == 0 {
		this.jsonError("<LABEL_986>！")
	}
	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_141>！")
	}

	ok, _ := models.RoleModel.HasSameName(roleId, name)
	if ok {
		this.jsonError("<LABEL_587>！")
	}
	_, err = models.RoleModel.Update(roleId, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		this.ErrorLog("<LABEL_1188> " + roleId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_1188>" + roleId + "<LABEL_1618>")
	}
	this.InfoLog("<LABEL_1188> " + roleId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_817>", nil, "/system/role/list")
}

func (this *RoleController) User() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	roleId := strings.TrimSpace(this.GetString("role_id", ""))

	if roleId == "" {
		this.ViewError("<LABEL_792>！")
	}
	keywords["role_id"] = roleId

	number, _ := this.GetRangeInt("number", 15, 10, 100)
	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	count, err = models.UserModel.CountUsersByKeywords(keywords)
	if err != nil {
		this.ErrorLog("<LABEL_215>: " + err.Error())
		this.ViewError("<LABEL_215>！", "/system/role/list")
	}
	users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	if err != nil {
		this.ErrorLog("<LABEL_398>: " + err.Error())
		this.ViewError("<LABEL_398>！", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_398>: " + err.Error())
		this.ViewError("<LABEL_215>！", "/system/main/index")
	}
	for _, user := range users {
		user["role_name"] = role["name"]
	}

	this.Data["users"] = users
	this.Data["roleId"] = roleId
	this.SetPaginator(number, count)
	this.viewLayout("role/user", "role")
}

func (this *RoleController) Privilege() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("<LABEL_986>", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_424>：" + err.Error())
		this.ViewError("<LABEL_425>！", "/system/role/list")
	}
	if len(role) == 0 {
		this.ViewError("<LABEL_986>", "/system/role/list")
	}

	menus, controllers, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("<LABEL_424>！")
	}

	var rolePrivileges = []map[string]string{}
	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		rolePrivileges, err = models.RolePrivilegeModel.GetRootRolePrivileges()
	} else {
		rolePrivileges, err = models.RolePrivilegeModel.GetRolePrivilegesByRoleId(roleId)
	}
	if err != nil {
		this.ViewError("<LABEL_426>")
	}

	this.Data["role"] = role
	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.Data["rolePrivileges"] = rolePrivileges
	this.Data["disabledPrivilegeIds"] = models.Privilege_Default_Ids

	this.viewLayout("role/privilege", "role")
}

func (this *RoleController) GrantPrivilege() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/role/list")
	}
	privilegeIds := this.GetStrings("privilege_id", []string{})
	roleId := this.GetString("role_id", "")

	if roleId == "" {
		this.jsonError("<LABEL_588>")
	}
	//if len(privilegeIds) == 0 {
	//	this.jsonError("<LABEL_589>")
	//}
	// add default privileges
	privilegeIds = append(privilegeIds, models.Privilege_Default_Ids...)

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1634> " + roleId + " <LABEL_1189>：" + err.Error())
		this.jsonError("<LABEL_986>")
	}
	if len(role) == 0 {
		this.jsonError("<LABEL_986>")
	}

	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_216>！")
	}

	res, err := models.RolePrivilegeModel.GrantRolePrivileges(roleId, privilegeIds)
	if err != nil {
		this.ErrorLog("<LABEL_1634> " + roleId + " <LABEL_1189>：" + err.Error())
		this.jsonError("<LABEL_818>！")
	}
	if !res {
		this.jsonError("<LABEL_818>")
	}

	this.InfoLog("<LABEL_1634> " + roleId + " <LABEL_1190>")
	this.jsonSuccess("<LABEL_819>", nil, "/system/role/list")
}

func (this *RoleController) Delete() {
	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/role/list")
	}

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.jsonError("<LABEL_792>！")
	}
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("<LABEL_292>！")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1191> " + roleId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_820>")
	}
	if len(role) == 0 {
		this.jsonError("<LABEL_986>")
	}
	if role["type"] == fmt.Sprintf("%d", models.Role_Type_System) {
		this.jsonError("<LABEL_427>！")
	}

	// check role user
	users, err := models.UserModel.GetUsersByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1191> " + roleId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_820>")
	}
	if len(users) > 0 {
		this.jsonError("<LABEL_821>，<LABEL_142>")
	}

	// delete role privilege by role id
	err = models.RolePrivilegeModel.DeleteByRoleId(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1191> " + roleId + " <LABEL_1151>: " + err.Error())
		this.jsonError("<LABEL_820>")
	}

	// delete role by role id
	err = models.RoleModel.Delete(roleId)
	if err != nil {
		this.ErrorLog("<LABEL_1191> " + roleId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_820>")
	}

	this.InfoLog("<LABEL_1191> " + roleId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_822>", nil, "/system/role/list")
}

func (this *RoleController) ResetUser() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/role/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("<LABEL_962>")
	}

	if this.UserId == "1" {
		this.jsonError("root <LABEL_428>！")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("<LABEL_1192> " + userId + " <LABEL_1193>: " + err.Error())
		this.jsonError("<LABEL_429>")
	}
	if len(user) == 0 {
		this.jsonError("<LABEL_962>")
	}

	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"role_id": models.Role_Default_Id,
	})
	if err != nil {
		this.ErrorLog("<LABEL_1192> " + userId + " <LABEL_1193>: " + err.Error())
		this.jsonError("<LABEL_429>")
	}

	this.InfoLog("<LABEL_1192> " + userId + " <LABEL_1194>")
	this.jsonSuccess("<LABEL_430>", nil, "/system/role/list")
}
