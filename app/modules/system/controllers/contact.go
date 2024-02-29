package controllers

import (
	"strings"
	"time"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"

	"github.com/astaxie/beego/validation"
)

type ContactController struct {
	BaseController
}

func (this *ContactController) Add() {
	this.viewLayout("contact/form", "contact")
}

func (this *ContactController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/contact/list")
	}
	name := strings.Trim(this.GetString("name", ""), "")
	mobile := strings.Trim(this.GetString("mobile", ""), "")
	position := strings.Trim(this.GetString("position", ""), "")
	email := strings.Trim(this.GetString("email", ""), "")

	v := validation.Validation{}
	if name == "" {
		this.jsonError("<LABEL_270>！")
	}
	if email == "" {
		this.jsonError("<LABEL_774>！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("<LABEL_568>！")
	}

	contact := map[string]interface{}{
		"name":     name,
		"mobile":   mobile,
		"position": position,
		"email":    email,
	}

	contactId, err := models.ContactModel.Insert(contact)
	if err != nil {
		this.ErrorLog("<LABEL_569>：" + err.Error())
		this.jsonError("<LABEL_569>")
	}
	this.InfoLog("<LABEL_978> " + utils.Convert.IntToString(contactId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_570>", nil, "/system/contact/list")
}

func (this *ContactController) List() {

	var err error
	var contacts []map[string]string
	contacts, err = models.ContactModel.GetAllContact()

	if err != nil {
		this.ErrorLog("<LABEL_271>: " + err.Error())
		this.ViewError("<LABEL_271>", "/system/main/index")
	}

	this.Data["contacts"] = contacts
	this.viewLayout("contact/list", "contact")
}

func (this *ContactController) Edit() {

	contactId := this.GetString("contact_id", "")
	if contactId == "" {
		this.ViewError("<LABEL_775>", "/system/contact/list")
	}

	contact, err := models.ContactModel.GetContactByContactId(contactId)
	if err != nil {
		this.ViewError("<LABEL_775>", "/system/contact/list")
	}

	this.Data["contact"] = contact
	this.viewLayout("contact/form", "contact")
}

func (this *ContactController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/contact/list")
	}
	contactId := strings.Trim(this.GetString("contact_id", ""), "")
	name := strings.Trim(this.GetString("name", ""), "")
	mobile := strings.Trim(this.GetString("mobile", ""), "")
	position := strings.Trim(this.GetString("position", ""), "")
	email := strings.Trim(this.GetString("email", ""), "")

	v := validation.Validation{}
	if contactId == "" {
		this.jsonError("<LABEL_1144>！")
	}
	if name == "" {
		this.jsonError("<LABEL_270>！")
	}
	if position == "" {
		this.jsonError("<LABEL_776>！")
	}
	if mobile == "" {
		this.jsonError("<LABEL_396>！")
	}
	if !v.Phone(mobile, "mobile").Ok {
		this.jsonError("<LABEL_272>！")
	}
	if email == "" {
		this.jsonError("<LABEL_774>！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("<LABEL_568>！")
	}

	contact := map[string]interface{}{
		"name":        name,
		"mobile":      mobile,
		"position":    position,
		"email":       email,
		"update_time": time.Now().Unix(),
	}
	_, err := models.ContactModel.UpdateByContactId(contact, contactId)
	if err != nil {
		this.ErrorLog("<LABEL_979> " + contactId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_571>")
	}
	this.InfoLog("<LABEL_979> " + contactId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_572>", nil, "/system/contact/list")
}

func (this *ContactController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/contact/list")
	}
	contactId := this.GetString("contact_id", "")
	if contactId == "" {
		this.jsonError("<LABEL_573>！")
	}

	contact, err := models.ContactModel.GetContactByContactId(contactId)
	if err != nil {
		this.ErrorLog("<LABEL_980> " + contactId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_574>")
	}
	if len(contact) == 0 {
		this.jsonError("<LABEL_775>")
	}

	_, err = models.ContactModel.DeleteByContactId(contactId)
	if err != nil {
		this.ErrorLog("<LABEL_980> " + contactId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_574>")
	}

	this.InfoLog("<LABEL_980> " + contactId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_575>", nil, "/system/contact/list")
}

// <LABEL_397>
func (this *ContactController) Import() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	username := strings.TrimSpace(this.GetString("username", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	if username != "" {
		keywords["username"] = username
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

	this.Data["users"] = users
	this.Data["username"] = username
	this.SetPaginator(number, count)
	this.viewLayout("contact/import", "contact")
}
