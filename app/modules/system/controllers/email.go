package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"

	"github.com/astaxie/beego"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EmailController struct {
	BaseController
}

func (this *EmailController) List() {

	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	var err error
	var emails []map[string]string
	if keyword != "" {
		emails, err = models.EmailModel.GetEmailsByLikeName(keyword)
	} else {
		emails, err = models.EmailModel.GetEmails()
	}
	if err != nil {
		this.ErrorLog("<LABEL_136>: " + err.Error())
		this.ViewError("<LABEL_136>", "/system/main/index")
	}

	this.Data["emails"] = emails
	this.Data["keyword"] = keyword
	this.viewLayout("email/list", "email")
}

func (this *EmailController) Add() {
	this.viewLayout("email/form", "email")
}

func (this *EmailController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", "25"))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))

	if name == "" {
		this.jsonError("<LABEL_137>！")
	}
	if host == "" {
		this.jsonError("<LABEL_138>！")
	}
	if validation.Validate(host, is.Host) != nil {
		this.jsonError("<LABEL_90>！")
	}
	if port == "" {
		this.jsonError("<LABEL_139>！")
	}
	if validation.Validate(port, is.Port) != nil {
		this.jsonError("<LABEL_91>！")
	}
	if senderAddress == "" {
		this.jsonError("<LABEL_278>！")
	}
	if username == "" {
		this.jsonError("<LABEL_279>！")
	}
	if password == "" {
		this.jsonError("<LABEL_280>！")
	}

	ok, err := models.EmailModel.HasEmailName(name)
	if err != nil {
		this.ErrorLog("<LABEL_281>：" + err.Error())
		this.jsonError("<LABEL_281>！")
	}
	if ok {
		this.jsonError("<LABEL_140>！")
	}

	emailId, err := models.EmailModel.Insert(map[string]interface{}{
		"name":                name,
		"sender_address":      senderAddress,
		"sender_name":         senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host":                host,
		"port":                port,
		"username":            username,
		"password":            password,
		"is_ssl":              isSsl,
	})

	if err != nil {
		this.ErrorLog("<LABEL_281>：" + err.Error())
		this.jsonError("<LABEL_281>")
	}
	this.InfoLog("<LABEL_583> " + utils.Convert.IntToString(emailId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_282>", nil, "/system/email/list")
}

func (this *EmailController) Edit() {

	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.ViewError("<LABEL_418>", "/system/email/list")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ViewError("<LABEL_418>", "/system/email/list")
	}

	this.Data["email"] = email
	this.viewLayout("email/form", "email")
}

func (this *EmailController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", ""))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))

	if emailId == "" {
		this.jsonError("<LABEL_418>！")
	}
	if name == "" {
		this.jsonError("<LABEL_137>！")
	}
	if host == "" {
		this.jsonError("<LABEL_138>！")
	}
	if validation.Validate(host, is.Host) != nil {
		this.jsonError("<LABEL_90>！")
	}
	if port == "" {
		this.jsonError("<LABEL_139>！")
	}
	if validation.Validate(port, is.Port) != nil {
		this.jsonError("<LABEL_91>！")
	}
	if senderAddress == "" {
		this.jsonError("<LABEL_278>！")
	}
	if username == "" {
		this.jsonError("<LABEL_279>！")
	}
	if password == "" {
		this.jsonError("<LABEL_280>！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("<LABEL_584> " + emailId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_283>！")
	}
	if len(email) == 0 {
		this.jsonError("<LABEL_418>！")
	}

	ok, _ := models.EmailModel.HasSameName(emailId, name)
	if ok {
		this.jsonError("<LABEL_140>！")
	}
	_, err = models.EmailModel.Update(emailId, map[string]interface{}{
		"name":                name,
		"sender_address":      senderAddress,
		"sender_name":         senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host":                host,
		"port":                port,
		"username":            username,
		"password":            password,
		"is_ssl":              isSsl,
	})

	if err != nil {
		this.ErrorLog("<LABEL_584> " + emailId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_283>")
	}
	this.InfoLog("<LABEL_584> " + emailId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_284>", nil, "/system/email/list")
}

func (this *EmailController) Used() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.jsonError("<LABEL_285>！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("<LABEL_984> " + emailId + " <LABEL_1184>: " + err.Error())
		this.jsonError("<LABEL_286>")
	}
	if len(email) == 0 {
		this.jsonError("<LABEL_418>")
	}
	_, err = models.EmailModel.SetEmailUsed(emailId)
	if err != nil {
		this.ErrorLog("<LABEL_984> " + emailId + " <LABEL_1184>: " + err.Error())
		this.jsonError("<LABEL_286>")
	}

	this.InfoLog("<LABEL_585> " + emailId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_287>", nil, "/system/email/list")
}

func (this *EmailController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.jsonError("<LABEL_285>！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("<LABEL_586> " + emailId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_288>")
	}
	if len(email) == 0 {
		this.jsonError("<LABEL_418>")
	}
	err = models.EmailModel.Delete(emailId)
	if err != nil {
		this.ErrorLog("<LABEL_586> " + emailId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_288>")
	}

	this.InfoLog("<LABEL_586> " + emailId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_289>", nil, "/system/email/list")
}

func (this *EmailController) Test() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", "25"))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))
	emails := strings.TrimSpace(this.GetString("emails", ""))

	if name == "" {
		this.jsonError("<LABEL_137>！")
	}
	if host == "" {
		this.jsonError("<LABEL_138>！")
	}
	if validation.Validate(host, is.Host) != nil {
		this.jsonError("<LABEL_90>！")
	}
	if port == "" {
		this.jsonError("<LABEL_139>！")
	}
	if validation.Validate(port, is.Port) != nil {
		this.jsonError("<LABEL_91>！")
	}
	if senderAddress == "" {
		this.jsonError("<LABEL_278>！")
	}
	if username == "" {
		this.jsonError("<LABEL_279>！")
	}
	if password == "" {
		this.jsonError("<LABEL_280>！")
	}
	if emails == "" {
		this.jsonError("<LABEL_92>！")
	}

	emailConfig := map[string]string{
		"sender_address":      senderAddress,
		"port":                port,
		"password":            password,
		"host":                host,
		"sender_name":         senderName,
		"username":            username,
		"sender_title_prefix": senderTitlePrefix,
		"is_ssl":              isSsl,
	}

	to := strings.Split(emails, ";")
	documentValue := map[string]string{
		"name":         "MM-Wiki<LABEL_1185>",
		"username":     this.User["username"],
		"update_time":  fmt.Sprintf("%d", time.Now().Unix()),
		"comment":      "",
		"document_url": "",
		"content":      "<LABEL_1186> <a href='https://github.com/chaiyd/mm-wiki'>MM-Wiki</a>，<LABEL_419>，<LABEL_985>",
	}

	emailTemplate := beego.BConfig.WebConfig.ViewsPath + "/system/email/template_test.html"
	body, err := utils.Email.MakeDocumentHtmlBody(documentValue, emailTemplate)
	if err != nil {
		this.ErrorLog("<LABEL_420>：" + err.Error())
		this.jsonError("<LABEL_420>！")
	}
	// start send email
	err = utils.Email.Send(emailConfig, to, "<LABEL_1185>", body)
	if err != nil {
		this.ErrorLog("<LABEL_420>：" + err.Error())
		this.jsonError("<LABEL_420>！")
	}

	this.jsonSuccess("<LABEL_421>", nil)
}
