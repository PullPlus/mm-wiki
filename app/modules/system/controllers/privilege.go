package controllers

import (
	"github.com/astaxie/beego"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"strings"
)

type PrivilegeController struct {
	BaseController
}

func (this *PrivilegeController) Add() {

	menus, _, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("<LABEL_399>！")
	}

	this.Data["menus"] = menus
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/form", "privilege")
}

func (this *PrivilegeController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_777>", "/system/privilege/add")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("<LABEL_45>")
	}

	name := strings.TrimSpace(this.GetString("name", ""))
	privilegeType := strings.TrimSpace(this.GetString("type", ""))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	controller := strings.TrimSpace(this.GetString("controller", ""))
	action := strings.TrimSpace(this.GetString("action", ""))
	target := strings.TrimSpace(this.GetString("target", ""))
	icon := strings.TrimSpace(this.GetString("icon", ""))
	isDisplay := strings.TrimSpace(this.GetString("is_display", "0"))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if name == "" {
		this.jsonError("<LABEL_400>！")
	}
	if privilegeType == "" {
		this.jsonError("<LABEL_401>！")
	}
	if privilegeType == "controller" {
		if parentId == "" {
			this.jsonError("<LABEL_133>！")
		}
		if controller == "" {
			this.jsonError("<LABEL_273>！")
		}
		if action == "" {
			this.jsonError("<LABEL_402>！")
		}
	}

	privilegeId, err := models.PrivilegeModel.Insert(map[string]interface{}{
		"name":       name,
		"type":       privilegeType,
		"parent_id":  parentId,
		"controller": controller,
		"action":     action,
		"target":     target,
		"icon":       icon,
		"is_display": isDisplay,
		"sequence":   sequence,
	})

	if err != nil {
		this.ErrorLog("<LABEL_778>：" + err.Error())
		this.jsonError("<LABEL_778>")
	}
	this.InfoLog("<LABEL_1172> " + utils.Convert.IntToString(privilegeId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_779>", nil, "/system/privilege/list")
}

func (this *PrivilegeController) List() {

	menus, controllers, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("<LABEL_780>！")
	}

	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/list", "privilege")
}

func (this *PrivilegeController) Edit() {

	privilegeId := this.GetString("privilege_id", "")
	if privilegeId == "" {
		this.ViewError("<LABEL_766>！", "/system/privilege/list")
	}

	privilege, err := models.PrivilegeModel.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("<LABEL_780>：" + err.Error())
		this.ViewError("<LABEL_780>！", "/system/privilege/list")
	}
	if len(privilege) == 0 {
		this.ViewError("<LABEL_981>！", "/system/privilege/list")
	}

	menus, _, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("<LABEL_780>！", "/system/privilege/list")
	}

	this.Data["menus"] = menus
	this.Data["privilege"] = privilege
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/form", "privilege")
}

func (this *PrivilegeController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_777>", "/system/privilege/list")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("<LABEL_46>")
	}
	privilegeId := strings.TrimSpace(this.GetString("privilege_id", ""))
	name := strings.TrimSpace(this.GetString("name", ""))
	privilegeType := strings.TrimSpace(this.GetString("type", ""))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	controller := strings.TrimSpace(this.GetString("controller", ""))
	action := strings.TrimSpace(this.GetString("action", ""))
	target := strings.TrimSpace(this.GetString("target", ""))
	icon := strings.TrimSpace(this.GetString("icon", "glyphicon-list"))
	isDisplay := strings.TrimSpace(this.GetString("is_display", "0"))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if name == "" {
		this.jsonError("<LABEL_400>！")
	}
	if privilegeType == "" {
		this.jsonError("<LABEL_401>！")
	}
	if privilegeType == "controller" {
		if parentId == "" {
			this.jsonError("<LABEL_133>！")
		}
		if controller == "" {
			this.jsonError("<LABEL_273>！")
		}
		if action == "" {
			this.jsonError("<LABEL_402>！")
		}
	}

	_, err := models.PrivilegeModel.Update(privilegeId, map[string]interface{}{
		"name":       name,
		"type":       privilegeType,
		"parent_id":  parentId,
		"controller": controller,
		"action":     action,
		"target":     target,
		"icon":       icon,
		"is_display": isDisplay,
		"sequence":   sequence,
	})

	if err != nil {
		this.ErrorLog("<LABEL_1173> " + privilegeId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_781>！")
	}
	this.InfoLog("<LABEL_1173> " + privilegeId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_782>", nil, "/system/privilege/list")
}

func (this *PrivilegeController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/privilege/list")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("<LABEL_47>")
	}
	privilegeId := this.GetString("privilege_id", "")
	if privilegeId == "" {
		this.jsonError("<LABEL_766>！")
	}

	privilege, err := models.PrivilegeModel.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("<LABEL_1174> " + privilegeId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_783>")
	}
	if len(privilege) == 0 {
		this.jsonError("<LABEL_981>")
	}

	// delete role_privilege by privilegeId
	err = models.RolePrivilegeModel.DeleteByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("<LABEL_784> " + privilegeId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_783>")
	}

	// delete privilege
	err = models.PrivilegeModel.Delete(privilegeId)
	if err != nil {
		this.ErrorLog("<LABEL_1174> " + privilegeId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_783>")
	}

	this.InfoLog("<LABEL_1174> " + privilegeId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_785>", nil, "/system/privilege/list")
}
